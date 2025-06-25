create-db:
	docker run --name mysql-db \
        -e MYSQL_ROOT_PASSWORD=secret \
        -e MYSQL_DATABASE=acosoDB \
        -p 3309:3306 \
        -d mysql:8.0

create-db-mongo:
	docker run --name mongo-db \
		-p 27017:27017 \
		-d mongo