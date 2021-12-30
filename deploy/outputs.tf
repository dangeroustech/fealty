output "Domain_IP" {
  value       = linode_domain_record.rewards.target
  description = "Domain A Reord"
  depends_on = [
    linode_instance.app,
    linode_nodebalancer.app_nb
  ]
}

output "NodeBalancer_Hostname" {
  value       = linode_nodebalancer.app_nb.hostname
  description = "NodeBalancers Hostname"
}

output "NodeBalancer_IPv4" {
  value       = linode_nodebalancer.app_nb.ipv4
  description = "NodeBalancers IPv4"
}

output "NodeBalancer_IPv6" {
  value       = linode_nodebalancer.app_nb.ipv6
  description = "NodeBalancers IPv6"
}

# output "App_Public_IP" {
#     value = linode_instance.app[count.index].ip_address
#     description = "App Instance's Public IP"
# }

output "DB_Private_IP" {
  value       = linode_instance.db.config[*].interface[0].ipam_address
  description = "DB Instance's Private IP"
}