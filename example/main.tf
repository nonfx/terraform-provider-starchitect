# example/main.tf
terraform {
  required_providers {
    starchitect = {
      source = "registry.terraform.io/nonfx/starchitect"
      version = "1.0.0"
    }
  }
}

provider "starchitect" {}

resource "starchitect_iac_pac" "demo_example" {
    iac_path = var.iac_path
    # pac_path = var.pac_path
    # pac_version = var.pac_version
}

variable "iac_path" {
  default = "../testdata/valid_iac"
}

variable "pac_path" {
  default = "../testdata/valid_pac"
}

variable "pac_version" {
  // starchitect-cloudguard github branch reference
  default = "main"
}

output "scan_result" {
    value = starchitect_iac_pac.demo_example.scan_result
}

output "score" {
    value = starchitect_iac_pac.demo_example.score
}
