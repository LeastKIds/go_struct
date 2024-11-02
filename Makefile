.PHONY: db
.DEFAULT_GOAL := help

db: ## 데이터베이스 실행
	docker compose -f ./docker/compose.yml up


help: ## 옵션 목록
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'