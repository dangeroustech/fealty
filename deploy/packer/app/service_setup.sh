#!/bin/bash

echo "###---Make APT Up To Date---###"
apt-get -qq update
apt-get -qq upgrade &> /dev/null

echo "###---Enable Fealty Service---###"
chmod a+x /usr/bin/fealty
echo "MONGO_URI=$MONGODB_FEALTY_URI" >> /etc/fealty/VARS
echo "MONGO_USER=fealty" >> /etc/fealty/VARS
echo "MONGO_PASS=$MONGODB_FEALTY_PASS" >> /etc/fealty/VARS
echo "FEALTY_USER=$FEALTY_USER" >> /etc/fealty/VARS
echo "FEALTY_PASS=$FEALTY_PASS" >> /etc/fealty/VARS
echo "FEALTY_CONFIG=/etc/fealty" >> /etc/fealty/VARS
echo "DOMAIN=$DOMAIN" >> /etc/fealty/VARS
systemctl enable fealty
cat /etc/fealty/VARS

echo "###---Reboot---###"
reboot
