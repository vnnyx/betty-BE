wire:
	wire github.com/vnnyx/betty-BE/internal/delivery/http
migrate:
	migrate -path ./migrations -database "cockroachdb://betty:BNvNycBF-NW-4vtAXthG6Q@aws-cockroachdb-5562.6xw.cockroachlabs.cloud:26257/betty?sslmode=verify-full" up