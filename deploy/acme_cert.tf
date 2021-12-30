resource "tls_private_key" "private_key" {
  algorithm = "RSA"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.private_key.private_key_pem
  email_address   = data.linode_profile.me.email
}

resource "acme_certificate" "certificate" {
  depends_on = [
    linode_domain.domain
  ]
  account_key_pem               = acme_registration.reg.account_key_pem
  common_name                   = "${var.subdomain}.${var.DOMAIN}"
  key_type                      = "4096"
  pre_check_delay               = 60
  revoke_certificate_on_destroy = false

  dns_challenge {
    provider = "linode"
    config = {
      LINODE_TOKEN = var.LINODE_TOKEN
      LINODE_TTL   = 300
    }
  }
}
