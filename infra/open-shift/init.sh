# Log in (using Red Hat OpenShift Dev Sandbox or your cluster)
oc login https://api.your-cluster.com:6443

# Create project
oc new-project web-app

# Deploy Nginx
oc new-app nginx

# Expose service
oc expose svc/nginx

# Get route to access the app
oc get route