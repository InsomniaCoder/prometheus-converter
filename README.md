# Instruction

This project is developped using Golang.

It is intended to be deployed as a sidecar container which will respond to a Prometheus scraping requests.

It will call to an actual service endpoint that expose healthcheck as a custom JSON format and convert the data into

prometheus readable format.

## Local

- run `make local` to run API internally
- run `make run` to rebuilt the project and start the service along with database using docker-compose
- run `make docker` to build a Docker image of this service

## Kubernetes

Use `helm upgrade prom-convert ./helm --install` to install the release

Please change `target.url` parameter in `values.yaml` to change to remote target
## API usage

### Endpoints

#### GET /api/v1/prometheus

##### using curl to try

```
curl --request GET http://localhost:8082/api/v1/prometheus
```

the response should look like this

```
gateway_up 1
face_comparison_up 1
thai_id_up 1
antispoofing_up 1
```