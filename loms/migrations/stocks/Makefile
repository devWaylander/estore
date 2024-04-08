GOOSE_DRIVER := "postgres"
GOOSE_DBSTRING := "postgres://postgres:12341234@127.0.0.1:1234/stocks?sslmode=disable"

upMigrationsStocksDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
downMigrationStocksDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
statusStocksDB:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) status
.reset:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset
resetStocksDB: .reset upMigrationsStocksDB