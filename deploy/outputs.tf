output "Domain_A_Record" {
    value = linode_domain_record.rewards.target
    description = "Domain A Reord"
    depends_on = [
        linode_instance.app,
    ]
}

output "App_Public_IP" {
    value = linode_instance.app.ip_address
    description = "App Instance's Public IP"
}

output "App_Private_IP" {
    value = linode_instance.app.config[*].interface[1].ipam_address
    description = "App Instance's Private IP"
}

output "DB_Private_IP" {
    value = linode_instance.db.config[*].interface[0].ipam_address
    description = "DB Instance's Private IP"
}