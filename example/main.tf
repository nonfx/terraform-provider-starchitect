# example/main.tf
terraform {
  required_providers {
    starchitect = {
      source = "hashicorp.com/edu/starchitect"
      version = "1.0.0"
    }
  }
}

provider "starchitect" {}

resource "starchitect_iac_pac" "demo_example" {
    iac_path = var.iac_path
    # pac_path = var.pac_path
}

variable "iac_path" {
  default = "../testdata/valid_iac"
}

variable "pac_path" {
  default = "../testdata/valid_pac"
}

output "scan_result" {
    value = starchitect_iac_pac.demo_example.scan_result
}

output "score" {
    value = starchitect_iac_pac.demo_example.score
}
