# BUILD
buildLoms:
	cd loms && docker build -t loms .

buildCart:
	cd cart && docker build -t cart .

buildNotifier:
	cd notifier && docker build -t notifier .

run-all: buildCart buildLoms buildNotifier
	docker compose up --build

# TESTS
# mocks gen
cartServiceMockGen:
	minimock -i ./cart/internal/handler.Service -o ./cart/internal/service && \
	cd ./cart/internal/service && minimock

lomsServiceMockGen:
	minimock -i ./loms/internal/app.LOMSService -o ./loms/internal/service && \
	minimock -i ./loms/internal/service.OrderRepository -o ./loms/internal/service && \
	minimock -i ./loms/internal/service.StockRepository -o ./loms/internal/service

# cart coverage
cartServiceCoverage:
	@go test ./cart/internal/service -timeout 5s -cover -count=1

cartRepoCoverage:
	@go test ./cart/internal/repo -cover -count=1

cartCoverage: cartServiceCoverage cartRepoCoverage

# benchmark
cartBenchmark:
	@go test -run=^# -bench=BenchmarkCreateCart route256.ozon.ru/project/cart/internal/repo -count=8 -cpu=16 -v

# race tests
cartGetTestRace:
	@go test -timeout 30s -run TestGetCart route256.ozon.ru/project/cart/internal/service -count=1 -v -race

cartAddTestRace:
	@go test -timeout 30s -run TestAddToCart route256.ozon.ru/project/cart/internal/service -count=1 -v -race

# cognit utils
cyclo:
	@echo "Cyclomatic Complexity:"
	@gocyclo -avg -ignore "_test" ./cart
	@echo ""
	
cognit:
	@echo "Cognitive Complexity:"
	@gocognit -avg -ignore "_test" ./cart

complexityLinters: cyclo cognit

# loms coverage
lomsServiceCoverage:
	@go test ./loms/internal/service -timeout 5s -cover -count=1

lomsCoverage: lomsServiceCoverage

# GUI for grpc requets
lomsGRPCUI:
	grpcui -port 5012 -plaintext localhost:8084