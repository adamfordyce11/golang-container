# Containerise a Golang application

## Setup

1. Install Golang version >= 1.20.
2. Install Docker or a similar container runtime.

## Building the container

Review the contents of the Dockerfile, this has the essential steps of taking the source code and building the application in a multi stage Dockerfile.

```bash
cd application
docker build . -t mytestimage
```

## Running Tests

```bash
cd application/src
go test -v
```

The expected output is as follows

```sh
=== RUN   TestGetOutput
--- PASS: TestGetOutput (0.00s)
PASS
ok  	printval	0.221s
```

## Verify the container using docker

```bash
# Start the container in the background
container_id=$(docker run -it --rm --detach mytestimage:latest)

# Execute a command in the container
docker exec $container_id /app/printval

# Review container logs
docker logs $container_id

# Review processes running inside the container
docler top $container_id
```
