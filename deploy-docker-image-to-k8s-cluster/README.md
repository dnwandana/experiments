# Deploy Docker Image to Kubernetes Cluster

1. Run code locally

   ```bash
   # download dependency modules
   go mod download

   # run code
   go run main.go
   ```

2. Build docker image

   ```bash
   docker build -t dnwandana/k8s-example:1.0 .
   ```

   Optional, you can push to dockerhub or any container registry.

   ```bash
   docker push dnwandana/k8s-example:1.0
   ```

3. Create deployment and service

   ```bash
   # create deployment
   kubectl apply -f manifests/deployment.yaml

   # create service
   kubectl apply -f manifests/service.yaml

   # or create deployment and service
   kubectl apply -f manifests/deployment-with-service.yaml
   ```

4. List all resource

   ```bash
   kubectl get all
   ```

5. Access the application

   Try visit `http://localhost:31000` and/or refresh the browser.

6. Try Deleting one pod

   ```bash
   # list all pods
   kubectl get pods

   # delete one of those pod
   kubectl delete pod k8s-example-deployment-xxxx-xxxx

   # list all pods
   kubectl get pods
   ```

   Then, access the application again.

7. Clean-up

   ```bash
   # delete deployment
   kubectl delete deployment k8s-example-deployment

   # delete service
   kubectl delete service k8s-example-service
   ```
