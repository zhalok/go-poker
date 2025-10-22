# Future Plan: REST API + Kubernetes Deployment

## Overview
Transform the CLI-based poker game into a scalable REST API service deployed on Kubernetes for improved availability and scalability.

---

## Phase 1: REST API Implementation

### API Endpoints

#### 1. Compare Poker Hands
```
POST /api/v1/poker/compare
Content-Type: application/json

Request Body:
{
  "hand1": "AAAKK",
  "hand2": "22233"
}

Response:
{
  "winner": "AAAKK",
  "hand1_type": "full_house",
  "hand2_type": "full_house",
  "result": "hand1",
  "timestamp": "2025-10-23T10:30:00Z"
}
```

#### 2. Health Check
```
GET /health

Response:
{
  "status": "healthy"
}
```

#### 3. Readiness Probe
```
GET /ready

Response:
{
  "status": "ready"
}
```

### Implementation Details
- **Framework**: Use Go standard library `net/http` or lightweight framework like `gin`
- **Port**: Expose on port `8080`
- **Logging**: Add structured logging for API requests
- **Error Handling**: Return proper HTTP status codes (400 for bad input, 500 for server errors)

### Code Changes Required
1. Create new `app/api/` directory
   - `server.go` - HTTP server setup
   - `handlers.go` - Request handlers
   - `middleware.go` - Logging, CORS, etc.
2. Update `app/main.go` to start HTTP server instead of CLI game
3. Update `dockerfile` to expose port 8080
4. Add health check endpoints for Kubernetes probes

---

## Phase 2: Kubernetes Deployment

### Architecture
- **Cluster**: Single-node Kubernetes cluster (Minikube, Kind, or k3s for local testing)
- **Pods**: 2 replica pods for high availability and load distribution
- **Service**: ClusterIP or LoadBalancer to expose the API
- **Scaling**: Horizontal Pod Autoscaler (HPA) based on CPU/memory usage

### Kubernetes Resources

#### 1. Deployment (`k8s/deployment.yaml`)
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: poker-api
  labels:
    app: poker-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: poker-api
  template:
    metadata:
      labels:
        app: poker-api
    spec:
      containers:
      - name: poker-api
        image: poker-go:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "128Mi"
            cpu: "200m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

#### 2. Service (`k8s/service.yaml`)
```yaml
apiVersion: v1
kind: Service
metadata:
  name: poker-api-service
spec:
  type: LoadBalancer  # Use NodePort for Minikube
  selector:
    app: poker-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

#### 3. ConfigMap (`k8s/configmap.yaml`)
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: poker-api-config
data:
  LOG_LEVEL: "info"
  PORT: "8080"
```

#### 4. Horizontal Pod Autoscaler (`k8s/hpa.yaml`)
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: poker-api-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: poker-api
  minReplicas: 2
  maxReplicas: 5
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

---

## Deployment Steps

### 1. Build Docker Image
```bash
docker build -t poker-go:latest .
```

### 2. Setup Local Kubernetes Cluster
```bash
# Using Minikube
minikube start

# OR using Kind
kind create cluster --name poker-cluster

# OR using k3s
curl -sfL https://get.k3s.io | sh -
```

### 3. Load Image to Cluster
```bash
# For Minikube
minikube image load poker-go:latest

# For Kind
kind load docker-image poker-go:latest --name poker-cluster
```

### 4. Deploy to Kubernetes
```bash
# Apply all manifests
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/hpa.yaml

# Verify deployment
kubectl get deployments
kubectl get pods
kubectl get services
```

### 5. Access the API
```bash
# For Minikube
minikube service poker-api-service --url

# For Kind/k3s with NodePort
kubectl port-forward service/poker-api-service 8080:80
```

### 6. Test the API
```bash
curl -X POST http://localhost:8080/api/v1/poker/compare \
  -H "Content-Type: application/json" \
  -d '{"hand1": "AAAKK", "hand2": "22233"}'
```

---

## Benefits of This Architecture

### Scalability
- **2 Pods**: Distribute load across multiple instances
- **HPA**: Automatically scale up to 5 pods based on traffic
- **Load Balancing**: Service distributes requests across healthy pods

### High Availability
- **Multiple Replicas**: If one pod fails, traffic routes to healthy pods
- **Self-Healing**: Kubernetes automatically restarts failed pods
- **Rolling Updates**: Zero-downtime deployments

### Resource Efficiency
- **Resource Limits**: Prevent pods from consuming excessive resources
- **Single Node**: Efficient for development/testing environments
- **Lightweight**: Alpine-based image keeps footprint small

### Monitoring & Observability
- **Health Probes**: Kubernetes monitors pod health automatically
- **Logs**: Centralized logging via `kubectl logs`
- **Metrics**: Can integrate Prometheus for detailed metrics

---

## Future Enhancements

1. **Persistent Storage**: Add Redis/PostgreSQL for game history
2. **Ingress Controller**: Use NGINX Ingress for advanced routing
3. **TLS/SSL**: Add HTTPS support with cert-manager
4. **Multi-Node Cluster**: Deploy across multiple nodes for true HA
5. **CI/CD Pipeline**: Automate builds and deployments with GitHub Actions
6. **Monitoring Stack**: Add Prometheus + Grafana for observability
7. **API Gateway**: Add rate limiting, authentication, API versioning
8. **Namespace Isolation**: Use different namespaces for dev/staging/prod

---

## Tools Required

- Docker Desktop or Docker Engine
- kubectl CLI
- Minikube / Kind / k3s (for local cluster)
- Optional: Helm (for package management)
- Optional: k9s (for cluster visualization)

---

## Estimated Timeline

- REST API Implementation: 2-3 hours
- Kubernetes Manifests: 1-2 hours
- Testing & Validation: 1-2 hours
- Documentation: 1 hour

**Total**: ~5-8 hours for complete implementation
