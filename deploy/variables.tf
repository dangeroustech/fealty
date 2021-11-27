variable "LINODE_TOKEN" {
  type        = string
  description = "Linode APIv4 Token"
}

variable "region" {
  type    = string
  default = "us-southeast"
}

variable "instance_type" {
  type    = string
  default = "g6-nanode-1"
}

variable "DOMAIN" {
  type        = string
  default     = ""
  description = "Domain to create"
}