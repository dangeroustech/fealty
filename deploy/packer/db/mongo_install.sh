#!/bin/bash

echo "###---Add MongoDB Repo Info---###"
wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
echo "###---Make APT Up To Date---###"
apt-get -qq update
apt-get -qq upgrade &> /dev/null
echo "###---Install MongoDB---###"
apt-get -qq install -y mongodb-org &> /dev/null
systemctl start mongod && systemctl enable mongod

echo "###---Configure MongoDB---###"
mongosh /tmp/mongo-init.js
sed -i.bak -e '/security/d' -e 's/127.0.0.1/"*"/g' /etc/mongod.conf
echo -e "security:\n  authorization: enabled" >> /etc/mongod.conf

echo "###---Reboot---###"
reboot
