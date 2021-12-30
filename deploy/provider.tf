terraform {
  required_providers {
    linode = {
      source  = "linode/linode"
      version = "1.25.0"
    }
    acme = {
      source = "vancluever/acme"
      version = "2.7.1"
    }
  }
  backend "s3" {
    bucket                      = "fealty"
    region                      = "us-east-1"
    endpoint                    = "us-east-1.linodeobjects.com"
    key                         = "tfstate"
    skip_credentials_validation = true
  }
}

provider "linode" {
  token = var.LINODE_TOKEN
}

provider "acme" {
  server_url = "https://acme-v02.api.letsencrypt.org/directory"
}
