#!/bin/bash

echo "###---Making APT Up To Date---###"
apt-get -qq update
apt-get -qq upgrade

echo "###---Enabling Fealty Service---###"
echo "MONGO_URI=$MONGODB_FEALTY_URI" >> /etc/fealty/VARS
echo "MONGO_USER=fealty" >> /etc/fealty/VARS
echo "MONGO_PASS=$MONGODB_FEALTY_PASS" >> /etc/fealty/VARS
echo "FEALTY_CONFIG=/etc/fealty" >> /etc/fealty/VARS
systemctl enable fealty

journalctl | grep fealty

# echo "###---Rebooting---###"
# reboot