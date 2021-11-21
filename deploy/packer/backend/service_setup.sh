#!/bin/bash

echo "###---Making APT Up To Date---###"
apt-get -qq update
apt-get -qq upgrade

echo "###---Enabling Fealty Service---###"
systemctl enable fealty

echo "###---Rebooting---###"
reboot