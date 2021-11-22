terraform {
  required_providers {
    linode = {
      source  = "linode/linode"
      version = "1.24.0"
    }
  }
  backend "s3" {
    bucket     = "fealty"
    region     = "us-east-1"
    endpoint   = "us-east-1.linodeobjects.com"
    key        = "tfstate"
    skip_credentials_validation = true
  }
}

provider "linode" {
  token = var.LINODE_TOKEN
}
