db-start:
	docker run --name postgres-container -p 5432:5432  -e POSTGRES_DB=test \
	-e POSTGRES_USER=test -e POSTGRES_PASSWORD=qwerty -d postgres:13.3

db-stop:
	docker stop postgres-container

db-clean:
	docker rm postgres-container