ecs-cli configure --cluster web-app --default-launch-type FARGATE --region us-east-1

ecs-cli up --cluster-config web-app --ecs-profile your-profile

ecs-cli compose --file docker-compose.yml service up