# Pivy - PivNets little helper

A small cli which uses HTTP range Headers and partial unzip to fetch tile metadata directly from Pivotal Network.

## Usage

```
usage: pivy --pivnet-api-token=PIVNET-API-TOKEN [<flags>] <command> [<args> ...]

PivNets little helper

Flags:
      --help                  Show context-sensitive help (also try --help-long
                              and --help-man).
      --accept-eula           Automatically accept EULA if necessary (Available
                              to select users only)
  -t, --pivnet-api-token=PIVNET-API-TOKEN
                              API token for network.pivotal.io (see:
                              https://network.pivotal.io/users/dashboard/edit-profile)
      --include-placeholders  replace obscured credentials with interpolatable
                              placeholders

Commands:
  help [<command>...]
    Show help.


  generate-config-template --product-slug=PRODUCT-SLUG --release-version=RELEASE-VERSION [<flags>]
    Generate om cli compatible config template

    -p, --product-slug=PRODUCT-SLUG
                            Product slug e.g. p-mysql
    -r, --release-version=RELEASE-VERSION
                            Release version e.g. 0.1.2-rc1
    -g, --glob="*.pivotal"  Glob to match product name e.g. *aws* should include
                            *.pivotal

  download-product-template --product-slug=PRODUCT-SLUG --release-version=RELEASE-VERSION [<flags>]
    Download raw tile metadata

    -p, --product-slug=PRODUCT-SLUG
                            Product slug e.g. p-mysql
    -r, --release-version=RELEASE-VERSION
                            Release version e.g. 0.1.2-rc1
    -g, --glob="*.pivotal"  Glob to match product name e.g. *aws* should include
                            *.pivotal



```
