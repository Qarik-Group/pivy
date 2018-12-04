package pivnet

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/gobwas/glob"
	"github.com/jeffallen/seekinghttp"
	"github.com/mholt/archiver"
	gopivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/go-pivnet/logshim"
	"github.com/pivotal-cf/pivnet-cli/filter"
)

const (
	retryAttempts = 5 // How many times to retry downloading a tile from PivNet
	retryDelay    = 5 // How long wait in between download retries
)

// Sdk interacts with the Pivotal Network API.
type Sdk struct {
	logger     *log.Logger
	client     gopivnet.Client
	filter     *filter.Filter
	acceptEULA bool
}

// NewSdk creates a new Pivotal Network Sdk.
// It validates that the given apiToken is valid.
func NewSdk(apiToken string, acceptEULA bool, logger *log.Logger) (*Sdk, error) {
	sdk := &Sdk{logger: logger, acceptEULA: acceptEULA}

	cfg := gopivnet.ClientConfig{
		Host:  gopivnet.DefaultHost,
		Token: apiToken,
	}
	sdk.client = gopivnet.NewClient(cfg, logshim.NewLogShim(logger, logger, false))

	sdk.filter = filter.NewFilter(logshim.NewLogShim(logger, logger, false))

	return sdk, sdk.checkCredentials()
}

func (s *Sdk) checkCredentials() error {
	ok, err := s.client.Auth.Check()

	if !ok {
		return fmt.Errorf("authorizing pivnet credentials: %v", err)
	}

	return nil
}

// TileFile is an PivNet tile
type TileFile struct {
	ProductSlug    string
	ReleaseVersion string
	Glob           string
}

// TileDownloadLink retrieves expiring download link for given Tile from PivNet.
func (s *Sdk) TileDownloadLink(tile TileFile) (downloadLink string, err error) {
	releases, err := s.client.Releases.List(tile.ProductSlug)
	if err != nil {
		return "", err
	}

	releases, err = s.filter.ReleasesByVersion(
		releases,
		tile.ReleaseVersion,
	)
	if err != nil {
		return "", err
	}

	if len(releases) != 1 {
		return "", fmt.Errorf(
			"Unable to find version %s for tile %s",
			tile.ReleaseVersion,
			tile.ProductSlug,
		)
	}
	release := releases[0]

	productFiles, err := s.client.ProductFiles.ListForRelease(
		tile.ProductSlug,
		release.ID,
	)
	if err != nil {
		return "", err
	}

	productFiles, err = s.filter.ProductFileKeysByGlobs(
		productFiles,
		[]string{tile.Glob},
	)
	if err != nil {
		return "", err
	}

	if len(productFiles) == 0 {
		return "", fmt.Errorf(
			"Unable find file with glob %s for tile %s",
			tile.Glob,
			tile.ProductSlug,
		)
	}

	if len(productFiles) > 1 {
		return "", fmt.Errorf(
			"More then one file matched glob %s for tile %s",
			tile.Glob,
			tile.ProductSlug,
		)
	}
	productFile := productFiles[0]

	if s.acceptEULA {
		err = s.client.EULA.Accept(tile.ProductSlug, release.ID)
		if err != nil {
			return "", err
		}
	}

	link, err := productFile.DownloadLink()
	if err != nil {
		return "", err
	}

	fetcher := gopivnet.NewProductFileLinkFetcher(link, s.client)
	downloadLink, err = fetcher.NewDownloadLink()
	if err != nil {
		return "", err
	}
	return downloadLink, nil
}

func (s *Sdk) DownloadTileMetaDataFromLink(link string) ([]byte, error) {
	seek := seekinghttp.New(link)
	size, err := seek.Size()
	if err != nil {
		return []byte{}, err
	}

	var DefaultZip = archiver.NewZip()
	err = DefaultZip.Open(seek, size)
	if err != nil {
		return []byte{}, err
	}
	defer DefaultZip.Close()

	metaData := new(bytes.Buffer)

	for {
		f, err := DefaultZip.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return []byte{}, err
		}

		zfh, ok := f.Header.(zip.FileHeader)
		if !ok {
			return []byte{},
				fmt.Errorf("expected header to be zip.FileHeader but was %T", f.Header)
		}

		if glob.MustCompile("metadata/*.yml").Match(zfh.Name) {
			metaData.ReadFrom(f)
			break
		}

		err = f.Close()
		if err != nil {
			return []byte{}, err
		}

	}
	return metaData.Bytes(), nil
}
