run:
	go run cmd/main.go

make_migration:
	migrate create -ext sql -dir ./schema -seq init

migrate:
	migrate -path ./schema -database 'root:$(PASSWORD)@tcp(127.0.0.1:3306)/mysqlpool'

rollback:
	migrate -path ./schema -database 'root:$(PASSWORD)@tcp(127.0.0.1:3306)/mysqlpool'

path_reset:
	export PATH=$PATH:$(go env GOPATH)/bin

mysql_connect:
	mysql -u 127.0.0.1 -u root -p avito