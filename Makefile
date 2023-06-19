.PHONY: binary proto build-builder build-server run-server

build-builder:
	cd build/package && docker-compose build builder

binary:
	cd build/pgo ackage && docker-compose run --rm builder

build-server:
	cd build/package && docker-compose build server

run-server:
	 cd build/package && docker-compose up server pgsql

#example: make proto name=user
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/$(name).proto

#example: make migrate-create name=users
migrate-create:
	migrate create -ext sql -dir internal/server/migrations -seq $(name)

#example: make migrate type=up
migrate:
	migrate -path internal/server/migrations -database "postgres://postgres:root@localhost:5433/postgres?sslmode=disable" -verbose $(type)