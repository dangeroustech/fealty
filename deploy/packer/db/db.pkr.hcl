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

        inline = [
            "wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -",
            "echo \"deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse\" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list",
            "apt update && apt upgrade -y && apt install -y mongodb-org",
            "systemctl start mongod && systemctl enable mongod",
            "mongosh /tmp/mongo-init.js",
            "sed -i.bak -e '/security/d' -e 's/127.0.0.1/\"*\"/g' /etc/mongod.conf",
            "echo \"\\nsecurity:\\n  authorization: enabled\" >> /etc/mongod.conf",
            "systemctl restart mongod",
            "mongosh mongodb://fealty:$MONGODB_FEALTY_PASS@localhost:27017/fealty?authSource=fealty --eval \"db.getCollectionNames()\"",
        ]
    }

    post-processor "manifest" {
      output = "../db-manifest.json"
      strip_time = true
    }
}
