# Starchitect Terraform Provider

The **Starchitect Terraform Provider** allows you to integrate Infrastructure as Code (IaC) and Policy as Code (PaC) workflows into your Terraform configuration. This provider scans your IaC and PaC files and populates a Terraform output variable named `scan_result` with the results.

---

## Features

- Accepts **IaC** (Infrastructure as Code) and **PaC** (Policy as Code) file paths as inputs.
- Scans the provided files for compliance and best practices.
- Outputs a detailed scan result to the Terraform output variable `scan_result`.

---

## Requirements

- Terraform `>= 1.0.0`

---

## Installation

Add the provider to your Terraform configuration. For example:

```hcl
terraform {
  required_providers {
    starchitect = {
      source  = "hashicorp.com/edu/starchitect"
      version = "1.0.0"
    }
  }
}
```
[Example Terraform](./example/main.tf)


# run locally

- update `dev_overrides` into `~/.terraformrc` file  to ensure that new provider's definition/functionality is considered from local

    ```
    provider_installation {
        dev_overrides {
            "hashicorp.com/edu/starchitect" = "<GOBIN PATH>"
        }
        direct {}
    }
    ```

- set go env `GOBIN` as `$GOPATH/bin`


- generate binary / update binary with new changes

    ```
    go install
    ```

- new provider is ready to be used locally. refer [example](./example/main.tf)