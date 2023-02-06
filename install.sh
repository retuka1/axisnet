#!/bin/bash
#installer Tunnaapi 

#Install Tunnapi
wget -O /usr/bin/axis https://raw.githubusercontent.com/adipatiarya/apis/master/axis


#izin permision
chmod +x /usr/bin/axis

#System tunnapi
wget -O /etc/systemd/system/axis.service https://raw.githubusercontent.com/adipatiarya/apis/master/axis.service && chmod +x /etc/systemd/system/axis.service

#restart service
systemctl daemon-reload

#Enable & Start & Restart 
systemctl enable axis.service
systemctl start axis.service
systemctl restart axis.service

rm -rf install.sh