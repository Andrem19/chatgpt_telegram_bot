migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/GptRemBot?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/GptRemBot?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: migratedown migrateup sqlc