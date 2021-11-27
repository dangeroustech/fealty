packer {
  required_version = "~> 1.7.0"
  required_plugins {
    linode = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/linode"
    }
  }
}

build {
  name    = "backend"
  sources = [
    "source.linode.fealty",
  ]

  provisioner "file" {
    source      = "deploy/packer/backend/fealty"
    destination = "/usr/bin/fealty"
  }

  provisioner "file" {
    source      = "deploy/packer/backend/fealty.service"
    destination = "/etc/systemd/system/fealty.service"
  }

  provisioner "shell" { 
    inline = [
      "mkdir /etc/fealty"
    ]
  }

  provisioner "file" {
    source      = "../../../backend/static"
    destination = "/etc/fealty/static"
  }

  provisioner "shell" {
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive",
      "MONGODB_FEALTY_URI=${var.MONGODB_FEALTY_URI}",
      "MONGODB_FEALTY_PASS=${var.MONGODB_FEALTY_PASS}",
    ]
    script = "service_setup.sh"
    expect_disconnect = true
  }

  provisioner "shell" {
    inline = [
      "systemctl status fealty",
    ]
  }

    post-processor "manifest" {
      output = "../backend-manifest.json"
    }
}
