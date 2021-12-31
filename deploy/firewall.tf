resource "linode_firewall" "fealty_fw" {
  depends_on = [
    linode_nodebalancer.app_nb,
  ]
  count    = var.node_count
  label    = "fealty-fw"
  tags     = ["fealty"]
  disabled = false

  inbound_policy = "DROP"

  inbound {
    label    = "allow-nodebalancer"
    action   = "ACCEPT"
    protocol = "TCP"
    ipv4     = ["${linode_nodebalancer.app_nb.ipv4}"]
    ipv6     = ["${linode_nodebalancer.app_nb.ipv6}"]
  }
  inbound {
    label    = "drop-inbound-udp"
    action   = "DROP"
    protocol = "UDP"
    ipv4     = ["0.0.0.0/0"]
    ipv6     = ["::/0"]
  }

  inbound {
    label    = "allow-inbound-icmp"
    action   = "ACCEPT"
    protocol = "ICMP"
    ipv4     = ["0.0.0.0/0"]
    ipv6     = ["::/0"]
  }

  outbound_policy = "ACCEPT"

  outbound {
    label    = "allow-https"
    action   = "ACCEPT"
    protocol = "TCP"
    ports    = "443"
    ipv4     = ["0.0.0.0/0"]
    ipv6     = ["::/0"]
  }

  linodes = [linode_instance.app[count.index].id]
}
