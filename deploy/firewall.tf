resource "linode_firewall" "fealty_fw" {
  count = var.node_count
  label = "fealty-fw"
  tags  = ["fealty"]
  disabled = true

  inbound_policy = "DROP"

  inbound {
    label    = "drop-inbound-tcp"
    action   = "DROP"
    protocol = "TCP"
    ipv4     = ["0.0.0.0/0"]
    ipv6     = ["::/0"]
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
