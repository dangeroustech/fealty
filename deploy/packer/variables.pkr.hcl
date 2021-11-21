#######################
## LINODE VARIABLES  ##
#######################

variable "linode_token" {
  type = string
  default = "{{env `PACKER_LINODE_TOKEN`}}"
}

variable "instance_type" {
  type = string
  default = "g6-nanode-1"
}

variable "image_label" {
  type = string
  default = "fealty"
}

variable "image_description" {
  type = string
  default = "Fealty Image"
}

variable "region"{
  type = string
  default = "eu-west"
}

#######################
## MONGODB VARIABLES ##
#######################

variable "MONGODB_ROOT_USER" {
  type    = string
  default = "admin"
}

variable "MONGODB_ROOT_PASS" {
  type    = string
  default = "admin"
}

variable "MONGODB_FEALTY_PASS" {
  type    = string
  default = "fealty"
}
