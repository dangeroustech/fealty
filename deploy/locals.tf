locals {
  backend_manifest = jsondecode(file("${path.module}/packer/backend-manifest.json"))
  db_manifest = jsondecode(file("${path.module}/packer/db-manifest.json"))
}
