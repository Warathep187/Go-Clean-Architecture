start-server:
	go run main.go

start-postgres:
	docker run -d \
		-p 5432:5432 \
		--name gocleanarch-postgres \
		-e POSTGRES_USER=homestead \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=clean_arch_blogs \
		postgres:alpine

stop-postgres:
	docker stop gocleanarch-postgres

start-mongo:
	docker run -d \
		-p 27017:27017 \
		--name gocleanarch-mongo \
		-e MONGO_INITDB_ROOT_USERNAME=homestead \
		-e MONGO_INITDB_ROOT_PASSWORD=secret \
		mongo:8.0.3

stop-mongo:
	docker stop gocleanarch-mongo
