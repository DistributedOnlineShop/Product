DBName = Product
User = root
Password = secret

createDB:
	docker exec -it psql createdb --username=root --owner=root $(DBName)

dropDB:
	docker exec -it psql dropdb $(DBName)

# Migrate 初期設定
# createMigrate:
# 	migrate create -ext sql -dir db/migration -seq init_table

#up
migrateup:
	migrate -path db/migration -database postgresql://${User}:${Password}@localhost:5432/${DBName}?sslmode=disable -verbose up

#down
migratedown:
	migrate -path db/migration -database postgresql://${User}:${Password}@localhost:5432/${DBName}?sslmode=disable -verbose down

#Sqlc
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .


.PHONY: postgres dropPsql createDB dropDB migrateup migratedown sqlc redis dropRedis test server