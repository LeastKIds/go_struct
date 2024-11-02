.PHONY: up db
.DEFAULT_GOAL := help

up: ## 서버 실행
	go run ./cmd/app/main.go

db: ## 데이터베이스 실행
	make docker-down
	docker compose -f ./docker/compose.yml up

docker-down: ## 도커 컨테이너 종료
	docker compose -f ./docker/compose.yml down

help: ## 항목 목록
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'