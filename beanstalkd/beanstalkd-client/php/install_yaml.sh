apt-get update
apt-get install libyaml-dev -y
printf "\n" | pecl install yaml
echo "extension=yaml.so" > /usr/local/etc/php/conf.d/docker-php-ext-yaml.ini
php -m | grep yaml