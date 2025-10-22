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
docker compose down
docker compose up -d --build
