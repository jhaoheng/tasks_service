> https://github.com/henszey/etcd-browser

docker run --rm --name etcd-browser -p 0.0.0.0:8000:8000 --env ETCD_HOST=10.10.0.1 --env AUTH_PASS=doe -t -i etcd-browser


