include ../../../.env

upMigrationsOrdersDB:
	goose ${GOOSE_DRIVER} ${GOOSE_ORDERS_DBSTRING} up
downMigrationOrdersDB:
	goose ${GOOSE_DRIVER} ${GOOSE_ORDERS_DBSTRING} down
statusOrdersDB:
	goose ${GOOSE_DRIVER} ${GOOSE_ORDERS_DBSTRING} status