# DevOps Challenge: Fix &amp; Deploy Go App with Redis

##  Overview:
- Your task is to troubleshoot, fix, and deploy a Go web application that uses Redis for caching. There are some issues in the Dockerfile or the Go code. After fixing these issues, you'll deploy the app to a Kubernetes cluster with Redis.

# 🚀 Go App Challenge - DEPI Program  

## 📌 Overview  
This project is a **Go web application** that connects to **Redis** for caching visitor counts. The challenge was to **fix, optimize, and deploy** the app using **Docker and Kubernetes** while ensuring persistent storage for Redis.  

This repository contains:
- **A fixed Dockerfile** with optimizations.
- **Kubernetes YAML files** for deployment.
- **Evidence of successful deployment** in Docker and Kubernetes.

---

## 🛠️ Challenge Breakdown

### **🔧 Part 1: Fixing & Optimizing Docker Deployment**  
1. **Troubleshooting & Fixes:**
   - The original **Dockerfile** contained incorrect configurations.
   - The Go app **failed to connect to Redis** due to missing environment variables.
   - The initial image size was **1GB**, which was too large.

2. **Improvements:**
   - Implemented **multi-stage builds**, reducing the image size from **1GB → 24.5MB**.
   - Ensured **correct paths and dependencies** were copied in the Dockerfile.
   - Configured **Docker networking** to allow the Go app to connect to Redis.
   - Successfully ran the app using **Docker Compose**.

3. **Pushing to Docker Hub:**  
   The optimized image was pushed to **Docker Hub** for easy deployment.  
   ```sh
   docker tag my-go-app:1.0 youssefashraf265/go-app-image:latest
   docker push youssefashraf265/go-app-image:latest
   
✅ **Image available at**: [`docker.io/youssefashraf265/go-app-image`](https://hub.docker.com/repository/docker/youssefashraf265/go-app-image)

---

## ☸️ Part 2: Deploying to Kubernetes

### **Setting Up Kubernetes Namespaces**
- `app` → For the Go application (**stateless**).
- `db` → For Redis (**stateful with persistent storage**).

### **Configuring Redis in Kubernetes**
- Deployed **Redis as a StatefulSet** to maintain a **stable network ID**.
- Ensured **persistent storage** for Redis.

### **Managing Configuration with ConfigMaps**
- Used **Kubernetes ConfigMaps** to manage environment variables.
- Exposed the Go app via a **NodePort service** for external access.

### **Deployment Issues & Fixes**
- Encountered **CrashLoopBackOff** due to incorrect image references.
- Fixed by **rebuilding and pushing the Docker image** before deployment.

---

## 📦 DevOps_Challenge_Depi  
 ┣ 📂 app  
 ┃ ┣ 📂 app  
 ┃ ┃ ┣ 📜 go-configmap.yaml  
 ┃ ┃ ┣ 📜 go-deployment.yaml  
 ┃ ┃ ┗ 📜 go-service.yaml  
 ┃ ┣ 📂 db  
 ┃ ┃ ┣ 📜 redis-deployment.yaml  
 ┃ ┃ ┗ 📜 redis-service.yaml  
 ┃ ┣ 📜 Dockerfile  
 ┃ ┣ 📜 docker-compose.yaml  
 ┃ ┣ 📜 go.mod  
 ┃ ┗ 📜 main.go  
 ┣ 📜 README.md  
 

# 🚀 Running the Project

## 1️⃣ Running with Docker Compose

*   Run the following command to start the application and Redis:

```sh
docker-compose up -d
```
*   Access the Go app at: http://localhost:8080

*   Test the Redis integration:

```sh
curl http://localhost:8080
```

Expected Output:

```sql
Hello, World! You are visitor number 1.
```

## 2️⃣ Running in Kubernetes
*   Apply the necessary Kubernetes configurations:

```sh
kubectl apply -f db/redis-service.yaml
kubectl apply -f db/redis-deployment.yaml
kubectl apply -f app/go-configmap.yaml
kubectl apply -f app/go-deployment.yaml
kubectl apply -f app/go-service.yaml
```
*   Get Minikube IP:
```sh
minikube ip
```
*   Access the app using:
```
curl http://<MINIKUBE_IP>:30080
```
##  📌 Key Takeaways

1.  Debugging & Fixing Dockerfiles✅ 
2.  Optimizing Image Size (1GB → 24.5MB)✅
3.  Deploying Stateful & Stateless Workloads in Kubernetes✅
4.  Managing Configuration via ConfigMaps✅
5.  Persistent Storage & Networking in Kubernetes✅
6.  Pushing Optimized Images to Docker Hub✅

## 🏆 Special Thanks
A huge thanks to DEPI Program for this hands-on challenge! 🚀
