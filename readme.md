Добавить товар в корзину: POST /user/<user_id>/cart/<sku_id>
Удалить товар из корзины: DELETE /user/<user_id>/cart/<sku_id>
Зачистить корзину: DELETE /user/<user_id>/cart
Получить корзину: GET /user/<user_id>/cart
Создать заказ из корзины: POST /user/<user_id>/cart/checkout

# GRPC
    - protoc
    - protoc-dependencies, protoc-generate, lomsGRPCUI - UI for GRPC requests
# Testing
    - testify, minimock, gocyclo, gocognit
    - cartServiceMockGen, lomsServiceMockGen
# SQL
    - sqlc, goose go
    - upMigrationsOrdersDB, downMigrationOrdersDB, statusOrdersDB
    - upMigrationsStocksDB, downMigrationStocksDB, statusStocksDB, resetStocksDB
# Kafka
    - sarama
    - everything is configurated in docker-compose

# ENV file
```bash
# postgres
POSTGRES_USER
POSTGRES_PASSWORD
DATABASE_HOST

# goose
GOOSE_DRIVER
GOOSE_STOCKS_DBSTRING
GOOSE_ORDERS_DBSTRING

# kafka-ui
KAFKA_CLUSTERS_0_NAME
KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS

# kafka
KAFKA_NODE_ID
KAFKA_LISTENER_SECURITY_PROTOCOL_MAP
KAFKA_ADVERTISED_LISTENERS
KAFKA_LISTENERS
KAFKA_CONTROLLER_LISTENER_NAMES
KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR
KAFKA_TRANSACTION_STATE_LOG_MIN_ISR
KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR
KAFKA_CONTROLLER_QUORUM_VOTERS
KAFKA_PROCESS_ROLES
KAFKA_LOG_DIRS
CLUSTER_ID
```