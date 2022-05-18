# container-status

A simple API that returns the status of all running Docker containers.

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
      - /var/run/docker.sock:/var/run/docker.sock
```

# Endpoints

| Method | Endpoint | Description                                 |
| ------ | -------- | ------------------------------------------- |
| GET    | /        | Returns a list of all running containers.   |
| GET    | /{name}  | Returns the status of a specific container. |
