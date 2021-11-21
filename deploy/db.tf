resource "linode_instance" "db" {
  label            = "fealty_db"
  tags             = ["fealty"]
  image            = local.db_manifest.builds[length(local.db_manifest.builds) -1].artifact_id
  region           = var.region
  type             = var.instance_type
  authorized_users = [data.linode_profile.me.username]
  // root_pass = "terr4form-test"
  backups_enabled  = true
  watchdog_enabled = true
  interface {
    purpose      = "vlan"
    label        = "fealty-vlan"
    ipam_address = "10.10.10.3/24"
  }
}
