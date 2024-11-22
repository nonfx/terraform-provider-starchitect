
# run locally

- update `dev_overrides` into `~/.terraformrc` file  to ensure that new provider's definition/functionality is considered from local

    ```
    provider_installation {

    dev_overrides {
        "hashicorp.com/edu/hashicups" = "/Users/  chandrashekhar/source-code/bin"
        "hashicorp.com/edu/starchitect" = "/Users/    chandrashekhar/source-code/bin"
    }

    # For all other providers, install them   directly from their origin provider
    # registries as normal. If you omit this,     Terraform will _only_ use
    # the dev_overrides block, and so no other    providers will be available.
    direct {}
    }

    ```

- set go env `GOBIN` as `$GOPATH/bin`


- generate binary / update binary with new changes

    ```
    go install
    ```

- new provider is ready to be used locally. refer example/main.tf