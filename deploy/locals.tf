locals {
  app_manifest = jsondecode(file("${path.module}/packer/app-manifest.json"))
  db_manifest      = jsondecode(file("${path.module}/packer/db-manifest.json"))
}
