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
##### `docker-compose`

Please, check for environment variables in `web` and `server` directories, and then follow example below

```bash
baseURL=\`danik.test\` baseAPIRoute=\`/api/server\` VUE_APP_BASE_URL=danik.test docker-compose up
```

Notes:
- it shoulbe "\`" because `traefik` waits for such string
- it should be `VUE_APP_BASE_URL` because `vuejs` expects the env var to have prefix `VUE_APP_`