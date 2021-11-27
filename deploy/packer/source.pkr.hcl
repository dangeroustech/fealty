source "linode" "fealty" {
  image             = "linode/ubuntu20.04"
  image_description = var.image_description
  image_label       = var.image_label
  instance_label    = var.image_label
  instance_type     = var.instance_type
  linode_token      = var.LINODE_TOKEN
  region            = var.region
  ssh_username      = "root"
}
