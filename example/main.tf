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

resource "starchitect_greeting" "example" {
  name = "starchitect"
}

resource "starchitect_iac_pac" "demo_example" {
    iac_path = var.iac_path
    pac_path = var.pac_path
}

variable "iac_path" {
  default = "/Users/chandrashekhar/source-code/src/github.com/nonfx/tf-regula-test/module-test/iac"
}

variable "pac_path" {
  default = "/Users/chandrashekhar/source-code/src/github.com/nonfx/stance/foundation/data/rules/aws"
}

output "scan_result" {
    value = starchitect_iac_pac.demo_example.scan_result
}

output "greeting" {
  value = starchitect_greeting.example.greeting
}