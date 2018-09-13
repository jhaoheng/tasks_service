
# phalcon  
- docker image : `https://hub.docker.com/r/mileschou/phalcon/tags/`
- baenstalk doc
    - https://docs.phalconphp.com/en/3.2/api/Phalcon_Queue_Beanstalk
    - https://docs.phalconphp.com/en/3.2/api/Phalcon_Queue_Beanstalk_Exception
    - https://docs.phalconphp.com/en/3.2/api/Phalcon_Queue_Beanstalk_Job

# how to use client

> please refer `phalcon-app/app/tasls/BeanstalkTask.php`
> use `./phalcon-app/run baenstalk help` to see detail.

1. docker-compose up -d
2. docker exec -it phalcon-app /bin/bash
3. `sh /home/install_yaml.sh`
    1. `apt-get install libyaml-dev`
    2. `pecl install yaml`
    3. `echo "extension=yaml.so" > /usr/local/etc/php/conf.d/docker-php-ext-yaml.ini`
    4. `php -m | grep yaml`
4. check `./phalcon-app/run baenstalk help`