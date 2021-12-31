resource "linode_firewall" "fealty-fw" {
  label = "fealty-fw"
  tags  = ["fealty"]

  inbound_policy = "DROP"
  outbound_policy = "ACCEPT"

  linodes = [linode_instance.app.id]
}
