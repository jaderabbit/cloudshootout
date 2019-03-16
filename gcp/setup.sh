# Create locally


# KUBECTL
gcloud components install kubectl

gcloud config set project cloud-shootout
gcloud config set compute/zone us-central1-b

git clone https://github.com/jaderabbit/cloudshootout.git
cd cloudshootout/helloapp

export PROJECT_ID="$(gcloud config get-value project -q)"
export VERSION=v2

docker build -t gcr.io/${PROJECT_ID}/hello-app:${VERSION} .

docker images

# Configure docker
gcloud auth configure-docker

# Create the image on container registry
docker push gcr.io/${PROJECT_ID}/hello-app:${VERSION}

# Run container locally
docker run --rm -p 8080:8080 gcr.io/${PROJECT_ID}/hello-app:${VERSION}

# Create locally
curl http://localhost:8080

# Create cluster with 3 nodes (#enough?!)
gcloud container clusters create cloudshootout --num-nodes=3


# List the instances
gcloud compute instances list

# Version
#kubectl run hello-web --image=gcr.io/${PROJECT_ID}/hello-app:${VERSION} --port 8080
# Update to use correct image.
kubectl apply -f ../gcp/helloweb-deployment.yaml

# Get Pods
kubectl get pods

kubectl expose deployment hello-web --type=LoadBalancer --port 80 --target-port 8080

# IP Assigned to service resource, not deployment
kubectl get service

# To Scale up 
kubectl scale deployment hello-web --replicas=3


# See new replicas
kubectl get deployment hello-web


# To deploy a new version
# docker build -t gcr.io/${PROJECT_ID}/hello-app:v2 .
# docker push gcr.io/${PROJECT_ID}/hello-app:v2
# kubectl set image deployment/hello-web hello-web=gcr.io/${PROJECT_ID}/hello-app:v2

# Clean up
# kubectl delete service hello-web
# gcloud container clusters delete hello-cluster