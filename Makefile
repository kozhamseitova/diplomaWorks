.PHONY:
migrate-up:
	migrate -source file://./schema -database postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable up

.PHONY:
create-migration:
	migrate create -ext sql -dir schema -seq ${name}