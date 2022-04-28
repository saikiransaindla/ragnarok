migrateup:
	migrate -path migration -database "mysql://root:password@tcp(localhost:3306)/ragnarok" -verbose up

migratedown:
	migrate -path migration -database "mysql://root:password@tcp(localhost:3306)/ragnarok" -verbose down