resource "linode_firewall" "fealty_fw" {
  count = var.node_count
  label = "fealty-fw"
  tags  = ["fealty"]

  inbound_policy = "DROP"
  outbound_policy = "ACCEPT"

  linodes = [linode_instance.app[count.index].id]
}
