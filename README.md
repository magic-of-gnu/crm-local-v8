#### Generate ed25519 keys

openssl genpkey -algorithm ed25519 -out keys/private.pem
openssl pkey -in keys/private.pem -pubout -out keys/public.pem


#### Set env vars
Set env vars everywhere and when running docker-compose
Possible paths for env vars:
- server/.env
- web/.env


eg:
```bash
baseURL=malcorp.test
baseAPIRoute=/api/server
VUE_APP_BASE_URL=malcorp.test/api/server
```


#### How to run
##### Run locally through `docker-compose`

Please, check for environment variables in `web` and `server` directories, and then follow example below

```bash
baseURL=\`danik.test\` baseAPIRoute=\`/api/server\` VUE_APP_BASE_URL=danik.test docker-compose up
```

Notes:
- it shoulbe "\`" because `traefik` waits for such string
- it should be `VUE_APP_BASE_URL` because `vuejs` expects the env var to have prefix `VUE_APP_`


#### Run remotely on digitalocean (do-server) server

Steps:
- configure environment variables
- copy docker images for backend and postgres or build them on the server
- copy front-end to `do-server`
- configure `nginx`
- - you have to have `location /` to front end
- - you have to have `location /api/server` to backend
- create docker network `docker network create -d bridge --attachable`
- run docker containers for backend and postres with proper container names and network names
- check for backend, front-end and db connection
- populate DB with tables and initial values
- you are ready to go :)


```bash
# creating docker network
docker network create -d bridge --attachable my-crm-network

# running docker container for backend
docker run --rm --name my-crm-server -it --network my-crm-network -p 8081:8081 --env-file .env server

# running docker container for postgres
docker run --rm --name postgres-server -it --network my-crm-network --expose 5432 --env-file .env postgres
```
