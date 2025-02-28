# DevOps Challenge: Fix &amp; Deploy Go App with Redis

##  Overview:
- Your task is to troubleshoot, fix, and deploy a Go web application that uses Redis for caching. There are some issues in the Dockerfile or the Go code. After fixing these issues, you'll deploy the app to a Kubernetes cluster with Redis.


## Part 1: Fix the Dockerfile and Go Application

- You’re given a Dockerfile & Go source code & dependencies.
- The Go app uses Redis to cache the number of visitors.
- Review and fix the issues in the Dockerfile or Go code to be able to build and run the app successfully.
- **Important:** For the Go app to run, you must have a running Redis container.
- **Bonus:** optimize the Go app image size. 


## Part 2: Deploy to Kubernetes
- Create Kubernetes YAML files to deploy Go (as **stateless** workload) and Redis (as **stateful** workload).
- Use **separate namespaces: "app" & "db"**.
- Redis should have persistent storage & network ID.
- Manage variables using Kubernetes native capabilities.
- One pod per each workload is fine, no need for over provisioning.
- Expose the Go app using nodeport or loadbalancer service.
- Redis should be exposed with the appropriate service type that is suitable for internal communications.


# Deliverables:
- A Github repository contains:
    - Fixed Dockerfile & Go App code.
    - Working Kubernetes YAML files.
- Evidence of the working Docker container & image build (status & logs & screenshots).
- Evidence of the working Kubernetes YAML files (status & logs & screenshots).
- **Bonus:** Evidence of optimized image size. 


# This challenge will help you:

- Troubleshoot Dockerfiles and Go applications.
- Build and run Docker containers locally.
- Deploy stateful & stateless workload to Kubernetes.
- Manage configurations and variables in Kubernetes environments.
- Use storage persistence in Kubernetes environments.
- Manage traffic across namespaces & persistent network IDs.
- Optimize Docker images size. 


