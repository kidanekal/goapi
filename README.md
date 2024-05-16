## Boilerplate for GO API Design



## Setup for local development

1. Install `Go 1.22.1`.
2. Install `make`.
3. Clone this repository into your Go workspace.


## Run the app (non docker)

1. `make run`

Testing with health endpoint

```
curl -iX GET http://localhost:4200/api/v1/health
HTTP/1.1 200 OK
goapi-Version: 1.0
Date: Tue, 12 Apr 2016 17:33:34 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```


## Run the app in docker


1. `make docker_run`

```
curl -iX GET localhost:4200/api/v1/health
HTTP/1.1 200 OK
Goapi-Version: 1.0
Date: Wed, 13 Apr 2016 14:31:51 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

## Local deployment

1. Set a default value for testing
   `export CI_COMMIT_SHA=$(git rev-parse --short HEAD)`
2. build the docker image
   `docker build -t goapi:$CI_COMMIT_SHA . `
3. deploy using helm
   `helm upgrade --install goapi-dev deployment -f ./deployment/values.yaml --set image.repository=goapi --set image.tag=CI_COMMIT_SHA`
4. Use port-forward to test the api
   `kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT`