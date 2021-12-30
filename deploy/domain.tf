resource "linode_domain" "domain" {
  type      = "master"
  domain    = var.DOMAIN
  soa_email = data.linode_profile.me.email
  ttl_sec = 30
  tags      = ["fealty"]
}

resource "linode_domain_record" "rewards" {
  depends_on = [
    linode_instance.app,
    linode_nodebalancer.app_nb
  ]
  domain_id   = linode_domain.domain.id
  name        = var.subdomain
  record_type = "A"
  target      = linode_nodebalancer.app_nb.ipv4
  ttl_sec     = 30
}