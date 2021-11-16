variable "linode_token" {
    type = string
    description = "Linode APIv4 Token"
}

variable "region" {
    type = string
    default = "us-southeast"
}

variable "instance_type" {
    type = string
    default = "g6-nanode-1"
}

variable "vlan_cidr" {
    type = string
    description = "IP Range to give to the Fealty VLAN"
    default = "10.10.10.0/24"
}
