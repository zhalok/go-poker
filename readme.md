# Go Poker CLI (Dockerized)

This is a simple **Go Poker CLI application** that can evaluate poker hands. It is packaged as a **Docker container** and managed using Docker Compose.

---

## Prerequisites

- Docker  
- Docker Compose  
- Go installed (optional, for local development)

---

## Build and Run

### 1. Build and start the container (detached mode)

```bash
docker compose up -d --build
```

### 2. Run the game in interactive mode
```bash
docker exec -it go-poker ./main
``` 

### 3. Stop and remove the containers to keep things clean
```bash
docker compose down
```