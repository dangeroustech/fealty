resource "linode_nodebalancer" "app_nb" {
  label                = "fealty-nodebalancer"
  region               = var.region
  client_conn_throttle = 20
  tags                 = ["fealty"]
}

resource "linode_nodebalancer_config" "app_nb_config" {
  nodebalancer_id = linode_nodebalancer.app_nb.id
  port            = 443
  protocol        = "https"
  check           = "http"
  check_path      = "/healthz"
  check_interval  = 30
  check_attempts  = 3
  check_timeout   = 5
  stickiness      = "http_cookie"
  algorithm       = "leastconn"
  cipher_suite    = "recommended"
  ssl_cert        = acme_certificate.certificate.certificate_pem
  ssl_key         = acme_certificate.certificate.private_key_pem
}

resource "linode_nodebalancer_node" "app_nb_node" {
  lifecycle {
    ignore_changes = all
  }
  count           = var.node_count
  nodebalancer_id = linode_nodebalancer.app_nb.id
  config_id       = linode_nodebalancer_config.app_nb_config.id
  label           = "app-node-${count.index + 1}"
  address         = "${linode_instance.app[count.index].private_ip_address}:3000"
  mode            = "accept"
}
