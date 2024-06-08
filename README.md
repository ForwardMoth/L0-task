## L0-task 

Work with nats-streaming, database Postgres and realize http handler

### Start project 

1. Start nats streaming
```bash
  nats-server --js
```
2. Start listener nats and http handler (go app)
```bash
  make
```
3. Publish in nats json
```bash
  make pub
```

Migration of table in migration directory 

YAML config struct in src/config

```yaml
env: "local"
http_server:
  address: "localhost:8080"
  timeout: time
  idle_timeout: time
db:
  username: "username"
  password: "password"
  DBname: "db_name"
  host: "host"
  port: "port"
  dialect: "dialect"
```
