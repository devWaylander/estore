GOOSE_DRIVER := "postgres"
GOOSE_DBSTRING := "postgres://postgres:12341234@127.0.0.1:1234/orders?sslmode=disable"

upMigrationsOrdersDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
downMigrationOrdersDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
statusOrdersDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) status