# container-status

A simple API that returns the status of all running Docker containers.

# Example docker-compose

```yaml
version: "3"

services:
  container-status:
    image: "ghcr.io/tetra-fox/container-status:latest"
    ports:
      - "3621:80"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

# Endpoints

| Method | Endpoint | Description                                 |
| ------ | -------- | ------------------------------------------- |
| GET    | /        | Returns a list of all running containers.   |
| GET    | /{name}  | Returns the status of a specific container. |
