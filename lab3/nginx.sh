
apt update
apt install nginx -y
cp -r /vagrant/html/. /var/www/html/
rm /var/www/html/index.nginx-debian.html


systemctl start nginx
systemctl enable nginx
systemctl restart nginx