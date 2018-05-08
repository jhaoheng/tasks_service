FROM        ubuntu:16.04

RUN         apt-get update && apt-get install nodejs npm -y 
RUN         ln -s $(which nodejs) /usr/bin/node && npm install -g redis-commander
