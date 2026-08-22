export PROJECT_ROOT=${shell pwd}

.DEFAULT_GOAL := help

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



run-parser: ## Go: Запустить парсер
	@go run -C ./JOB_FINDER ./cmd/parser/main.go