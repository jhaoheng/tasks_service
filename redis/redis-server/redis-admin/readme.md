# redis-commander install

> https://github.com/joeferner/redis-commander

```
apt-get update
apt-get install nodejs npm -y
npm install -g redis-commander
```

# redis-commander help

```
Options:
  --redis-port                    The port to find redis on.              [string]
  --sentinel-port                 The port to find sentinel on.           [string]
  --redis-host                    The host to find redis on.              [string]
  --sentinel-host                 The host to find sentinel on.           [string]
  --redis-socket                  The unix-socket to find redis on.       [string]
  --redis-password                The redis password.                     [string]
  --redis-db                      The redis database.                     [string]
  --http-auth-username, --http-u  The http authorisation username.        [string]
  --http-auth-password, --http-p  The http authorisation password.        [string]
  --address, -a                   The address to run the server on.       [string]  [default: "0.0.0.0"]
  --port, -p                      The port to run the server on.          [string]  [default: 8081]
  --nosave, --ns                  Do not save new connections to config.  [boolean]  [default: true]
  --save, -s                      Save new connections to config.         [boolean]
  --noload, --nl                  Do not load connections from config.    [boolean]
  --clear-config, --cc            clear configuration file              
  --root-pattern, --rp            default root pattern for redis keys     [default: "*"]
```