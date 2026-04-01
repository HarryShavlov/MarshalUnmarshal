
.PHONY: run_client run_serv deps

run_client:
	go run client/main.go

run_serv:
	go run cmd/main.go

deps:
	go mod tidy

help:
	@echo "Доступные команды:"
	@echo "  make run    - Запустить сервер в режиме разработки"
	@echo "  make deps   - Обновить зависимости"


