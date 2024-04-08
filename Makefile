buildLoms:
	cd loms && docker build -t loms .

buildCart:
	cd cart && docker build -t cart .

run-all: buildCart buildLoms
	docker compose up --build

cartServiceMockGen:
	minimock -i ./cart/internal/handler.Service -o ./cart/internal/service && \
	cd ./cart/internal/service && minimock

lomsServiceMockGen:
	minimock -i ./loms/internal/app.LOMSService -o ./loms/internal/service && \
	minimock -i ./loms/internal/service.OrderRepository -o ./loms/internal/service && \
	minimock -i ./loms/internal/service.StockRepository -o ./loms/internal/service

cartServiceCoverage:
	@go test ./cart/internal/service -timeout 5s -cover -count=1

cartRepoCoverage:
	@go test ./cart/internal/repo -cover -count=1

cartCoverage: cartServiceCoverage cartRepoCoverage

cartBenchmark:
	@go test -run=^# -bench=BenchmarkCreateCart route256.ozon.ru/project/cart/internal/repo -count=8 -cpu=16 -v

cyclo:
	@echo "Cyclomatic Complexity:"
	@gocyclo -avg -ignore "_test" ./cart
	@echo ""
	
cognit:
	@echo "Cognitive Complexity:"
	@gocognit -avg -ignore "_test" ./cart

complexityLinters: cyclo cognit

lomsServiceCoverage:
	@go test ./loms/internal/service -timeout 5s -cover -count=1

lomsCoverage: lomsServiceCoverage