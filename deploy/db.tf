data "linode_images" "fealty_image" {
  filter {
    name = "label"
    values = ["fealty-db"]
  }

  filter {
    name = "is_public"
    values = ["false"]
  }

  latest = true
}

data "linode_profile" "me" {}

data "linode_user" "me" {
    username = data.linode_profile.me.username
}

resource "linode_instance" "db" {
    label = "fealty_db"
    tags = [ "fealty" ]
    image = data.linode_images.fealty_image.images.0.id
    region = var.region
    type = var.instance_type
    authorized_users = [ data.linode_profile.me.username ]
    // root_pass = "terr4form-test"
    backups_enabled = true
    watchdog_enabled = true
    interface {
        purpose = "vlan"
        label = "fealty-vlan"
        ipam_address = var.vlan_cidr
    }
}
