#######################
## LINODE VARIABLES
#######################

../.variables.pkr.hcl

#######################
## MONGODB VARIABLES
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
