resource "linode_instance" "app" {
  count            = var.node_count
  label            = "fealty-app-${count.index + 1}"
  tags             = ["fealty"]
  image            = local.app_manifest.builds[length(local.app_manifest.builds) - 1].artifact_id
  region           = var.region
  type             = var.instance_type
  authorized_users = [data.linode_profile.me.username]
  backups_enabled  = false
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
