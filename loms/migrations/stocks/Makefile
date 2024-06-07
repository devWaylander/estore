include ../../../.env

upMigrationsStocksDB:
	goose ${GOOSE_DRIVER} ${GOOSE_STOCKS_DBSTRING} up
downMigrationStocksDB:
	goose ${GOOSE_DRIVER} ${GOOSE_STOCKS_DBSTRING} down
statusStocksDB:
	goose ${GOOSE_DRIVER} ${GOOSE_STOCKS_DBSTRING} status
.reset:
	goose ${GOOSE_DRIVER} ${GOOSE_STOCKS_DBSTRING} reset
resetStocksDB: .reset upMigrationsStocksDB