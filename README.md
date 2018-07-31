## Boilerplate for GO API Design



## Setup for local development

1. install godep from here https://github.com/tools/godep
2. install `make`


## Run the app (non docker)

1. `make run`

Testing with health endpoint

```
curl -iX GET http://localhost:4200/health
HTTP/1.1 200 OK
goapi-Version: 1.0
Date: Tue, 12 Apr 2016 17:33:34 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```


## Run the app in docker


1. `make docker_run`

```
curl -iX GET localhost:4200/health
HTTP/1.1 200 OK
Goapi-Version: 1.0
Date: Wed, 13 Apr 2016 14:31:51 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```
