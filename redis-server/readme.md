# use

> image : https://hub.docker.com/_/redis/

1. 填入 docker-compose.yml 中的 ip 位置
    - 如果是要接收本地端 docker-redis，請用 `links: redis` 進行設定，如目前 docker-compose.yml 設定
    - 如果要接收外部的 redis，請設定 ip。注意 : 
        1. redis-server，conf 裏的 `protected-mode no`
        2. docker-compose.yml 中的 'redis' 改為 redis-server ip
2. `docker-compose up -d`
3. browser : `localhost:8081`