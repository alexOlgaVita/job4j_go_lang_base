# Переменные
GO := go
PKG := ./...

.PHONY: test lint coverage all

test:
	$(GO) test $(PKG)

coverage:
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html

# Вывод покрытия в терминал (опционально)
.PHONY: cover-report
cover-report:
	$(GO) test -cover $(PKG)

lint:
	golangci-lint run

# Запуск всех проверок
.PHONY: all
all: lint test coverage