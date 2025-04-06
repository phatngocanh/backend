swagger:
	swag init --parseDependency --parseInternal
wire:
	wire ./internal
build:
	docker rmi vukhoa23/pna-invoice-be:latest || true
	docker build -t vukhoa23/pna-invoice-be:latest .
	docker push vukhoa23/pna-invoice-be:latest