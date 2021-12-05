resource "linode_domain" "domain" {
  type      = "master"
  domain    = var.DOMAIN
  soa_email = data.linode_profile.me.email
  tags      = ["fealty"]
}

resource "linode_domain_record" "rewards" {
  depends_on = [
    linode_instance.app
  ]
  domain_id   = linode_domain.domain.id
  name        = "rewards"
  record_type = "A"
  target      = linode_instance.app.ip_address
  ttl_sec     = 30
}