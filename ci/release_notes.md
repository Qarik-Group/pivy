# Initial Release

```
usage: pivy --pivnet-api-token=PIVNET-API-TOKEN [<flags>] <command> [<args> ...]

PivNets little helper

Flags:
      --help                  Show context-sensitive help (also try --help-long and --help-man).
      --version               Show application version.
      --accept-eula           Automatically accept EULA if necessary (Available to select users only)
  -t, --pivnet-api-token=PIVNET-API-TOKEN
                              API token for network.pivotal.io (see: https://network.pivotal.io/users/dashboard/edit-profile)
      --include-placeholders  replace obscured credentials with interpolatable placeholders

Commands:
  help [<command>...]
    Show help.

  generate-config-template --product-slug=PRODUCT-SLUG --release-version=RELEASE-VERSION [<flags>]
    Generate om cli compatible config template

  download-product-template --product-slug=PRODUCT-SLUG --release-version=RELEASE-VERSION [<flags>]
    Download raw tile metadata
```
