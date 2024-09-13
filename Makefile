create-db:
	docker run --name mysql-db \
        -e MYSQL_ROOT_PASSWORD=secret \
        -e MYSQL_DATABASE=acosoDB \
        -p 3309:3306 \
        -d mysql:8.0

create-db-neo4j:
	docker volume create neo4j_data
	docker run --name neo4j \
                 -p 7474:7474 -p 7687:7687 \
                 -v neo4j_data:/data \
                 -e NEO4J_AUTH=neo4j/testpassword \
                 -d neo4j

create-db-mongo:
	docker run --name mongo-db \
		-p 27017:27017 \
		-d mongo