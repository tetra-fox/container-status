# container-status

A tiny API written in Go that returns metadata of all Docker containers.

I created this because I wanted a simple way to display the status of the various services that run my home network without displaying sensitive information (such as my network configuration, volume bindings or entrypoint arguments).

You can see a live instance of this [here](https://home.tetra.cool/status)!

# docker-compose example

```yaml
version: "3.7"

services:
  container-status:
    container_name: status
    image: "ghcr.io/tetra-fox/container-status:latest"
    ports:
      - "3621:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # So we can get information from Docker!
```

# Endpoints

| Method | Endpoint       | Description                                                           |
| ------ | -------------- | --------------------------------------------------------------------- |
| GET    | /              | Returns the metadata of all containers.                               |
| GET    | /{name},{name} | Returns the metadata of the specified container(s). (comma-separated) |
