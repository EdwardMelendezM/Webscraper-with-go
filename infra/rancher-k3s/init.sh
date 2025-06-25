# Install K3s (on Linux)
curl -sfL https://get.k3s.io | sh -

# Confirm it's running
kubectl get nodes

# Deploy a simple app
kubectl create deployment web --image=nginx
kubectl expose deployment web --port=80 --type=NodePort