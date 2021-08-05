# Instruction

This project is developped using Golang.

It is intended to be deployed as a sidecar container which will respond to a Prometheus scraping requests.

It will call to an actual service endpoint that expose healthcheck as a custom JSON format and convert the data into

prometheus readable format.

## Local

- run `make local` to run API internally
- run `make run` to rebuilt the project and start the service along with database using docker-compose
- run `make docker` to build a Docker image of this service

## API usage

### Endpoints

#### GET /api/v1/prometheus

##### using curl to try

```
curl --request GET http://localhost:8082/api/v1/prometheus
```