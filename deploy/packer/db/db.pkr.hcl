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
  name    = "db"
  sources = [
    "source.linode.fealty",
  ]

  provisioner "file" {
      source      = "mongo-init.js"
      destination = "/tmp/mongo-init.js"
  }

  provisioner "shell" {
        environment_vars = [
            "MONGODB_ROOT_USER=${var.MONGODB_ROOT_USER}",
            "MONGODB_ROOT_PASS=${var.MONGODB_ROOT_PASS}",
            "MONGODB_FEALTY_PASS=${var.MONGODB_FEALTY_PASS}",
            "DEBIAN_FRONTEND=noninteractive",
        ]
        scripts = [
          "mongo_install.sh",
          "mongo_test.sh",
        ]
        expect_disconnect = true
    }

    post-processor "manifest" {
      output = "../db-manifest.json"
    }
}
