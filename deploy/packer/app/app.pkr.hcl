build {
  name    = "app"
  sources = [
    "source.linode.fealty",
  ]

  provisioner "file" {
    source      = "deploy/packer/app/fealty"
    destination = "/usr/bin/fealty"
  }

  provisioner "file" {
    source      = "deploy/packer/app/fealty.service"
    destination = "/etc/systemd/system/fealty.service"
  }

  provisioner "shell" { 
    inline = [
      "mkdir /etc/fealty"
    ]
  }

  provisioner "file" {
    source      = "app/static"
    destination = "/etc/fealty/static"
  }

  provisioner "shell" {
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive",
      "MONGODB_FEALTY_URI=${var.MONGODB_FEALTY_URI}",
      "MONGODB_FEALTY_PASS=${var.MONGODB_FEALTY_PASS}",
      "FEALTY_USER=${var.FEALTY_USER}",
      "FEALTY_PASS=${var.FEALTY_PASS}",
    ]
    script = "deploy/packer/app/service_setup.sh"
    expect_disconnect = true
  }

  provisioner "shell" {
    pause_before = "5s"
    inline = [
      "systemctl status fealty",
    ]
  }

    post-processor "manifest" {
      output = "deploy/packer/app-manifest.json"
    }
}
