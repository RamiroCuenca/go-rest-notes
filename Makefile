# Create the container with docker using the postgres image
postgres:
	docker run --name go-rest-notes -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

# This rule creates the db
createdb:
	docker exec -it go-rest-notes createdb --username=root --owner=root notes_db

# This rule deletes the db
dropdb:
	docker exec -it go-rest-notes dropdb bank_app

# This rule runs the migrations up
run-migrations-up:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/notes_db?sslmode=disable" --verbose up

run-migrations-down:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/notes_db?sslmode=disable" --verbose down

# .PHONY tell explicitly to MAKE that those rules are not associated with files
.PHONY: postgres createdb dropb run-migrations-up run-migrations-down