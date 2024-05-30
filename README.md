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

### Prerequisites
* Ensure you have minikube installed
* Ensure you have Docker installed on your machine
1. start minikube
   ```sh 
   minikube start
2. set-up minikube's Docker Daemon
   * macOS and Linux
     * configure your shell to use Minikube's Docker daemon by running the following command:
     ```sh 
     eval $(minikube -p minikube docker-env) 
   * Windows
     * In PowerShell, configure the environment by running:
     ```sh
     & minikube -p minikube docker-env | Invoke-Expression
3. verify the docker environment
   ```sh
   docker info
4. Use `make all` command to start the deployment 
   ```sh
   make all
5. Forward the port using `kubectl`
   ```sh
   kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
   
## Setting Up Monitoring for goapi

### Prerequisites

- Helm installed and initialized
- Kubernetes cluster running
- `goapi` application deployed

### Step 1: Install Prometheus using Helm

Run the following command to deploy Prometheus and set up monitoring in the `monitoring` namespace:

```sh
helm upgrade --install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace
```
### Step 2: Apply the ServiceMonitor

To make Prometheus start scraping metrics from your goapi application, you need to apply servicemonitor-goapi.yaml configuration.

```sh
kubectl apply -f deployment/servicemonitor-goapi.yaml

