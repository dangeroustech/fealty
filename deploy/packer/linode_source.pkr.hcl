source "linode" "fealty" {
  image             = "linode/ubuntu20.04"
  image_description = "Fealty ${{var.image_short_description}} Image"
  image_label       = var.image_label
  instance_label    = "fealty-${{var.instance_short_label}}-packer"
  instance_type     = var.instance_type
  linode_token      = var.linode_token
  region            = var.region
  ssh_username      = "root"
}
