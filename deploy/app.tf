resource "linode_instance" "app" {
  count            = var.node_count
  label            = "fealty_app_${count.index + 1}"
  tags             = ["fealty"]
  image            = local.app_manifest.builds[length(local.app_manifest.builds) - 1].artifact_id
  region           = var.region
  type             = var.instance_type
  authorized_users = [data.linode_profile.me.username]
  # root_pass = "terr4form-test"
  backups_enabled  = true
  watchdog_enabled = true
  interface {
    purpose = "public"
  }
  interface {
    purpose      = "vlan"
    label        = "fealty-vlan"
    ipam_address = "10.10.10.2/24"
  }
  private_ip = true
}

// Extra Private IP for NodeBalancer to use
# resource "linode_instance_ip" "app_private" {
#     linode_id = linode_instance.app[count.index].id
#     public = false
#     apply_immediately = true
# }