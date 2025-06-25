# Init Swarm (only once)
docker swarm init

# Deploy an Nginx service
docker service create \
  --name web \
  --publish 8080:80 \
  --replicas 3 \
  nginx

# To see running services
docker service ls
docker service ps web