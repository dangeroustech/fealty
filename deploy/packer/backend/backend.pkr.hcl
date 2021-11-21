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
    source      = "fealty"
    destination = "/usr/bin/fealty"
  }

  provisioner "file" {
    source      = "fealty.service"
    destination = "/etc/systemd/system/fealty.service"
  }

  provisioner "shell" {
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive",
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
