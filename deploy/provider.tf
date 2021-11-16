terraform {
  required_providers {
    linode = {
      source  = "linode/linode"
      version = "1.24.0"
    }
  }
}

provider "linode" {
  token = var.linode_token
}
