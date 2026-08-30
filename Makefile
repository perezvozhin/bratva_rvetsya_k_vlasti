## Запуск каждого таргета осуществляется с корня рабочей директории!
include .env
export

export PROJECT_ROOT=${shell pwd}

.DEFAULT_GOAL := help

db-cleanup-sqlite: ## SQLite: Удалить файл базы данных
	@read -p "Удалить базу данных SQLite? Все собранные вакансии будут потеряны. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -f ./out/sqlite/vacancies.db; \
		echo "Файл vacancies.db удален. Для пересоздания запусти make goose-up-sqlite"; \
	else \
		echo "Очистка отменена"; \
	fi

goose-create-sqlite: ## SQLite: Создать миграции
	@mkdir -p ${PROJECT_ROOT}/out/sqlite ${PROJECT_ROOT}/migrations/sqlite; \
	if [ -z "$(name)" ]; then \
		echo "Отсутствует необходимый параметр name. Пример: make goose-create-sqlite name=init"; \
		exit 1; \
	fi; \
	docker compose run --rm --user "$$(id -u):$$(id -g)" jobfinder-sqlite-goose create -s "$(name)" sql

goose-up-sqlite: ## SQLite: Накатить миграции
	@$(MAKE) goose-action-sqlite action=up

goose-down-sqlite: ## SQLite: Откатить миграции
	@$(MAKE) goose-action-sqlite action=down

goose-action-sqlite: ## SQLite: Применить команду миграции
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make goose-action-sqlite action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm --user "$$(id -u):$$(id -g)" jobfinder-sqlite-goose "$(action)"



run-app: ## Go: Запустить приложение
	@go run -C ./JOB_FINDER ./cmd/main.go

run-front: ## Front: Запустить Vite dev-сервер
	@npm --prefix ./web/app run dev

run-all: ## Запустить фронт и бэк вместе
	@trap 'kill 0' INT TERM EXIT; \
	$(MAKE) run-app & \
	$(MAKE) run-front & \
	wait


help: ## Show help for commands
	@echo "=== Help ==="
	@echo ""
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
