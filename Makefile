.PHONY: up db
.DEFAULT_GOAL := help

up: ## 서버 실행
	go run ./cmd/app/main.go

db: ## 데이터베이스 실행
	make docker-down
	docker compose -f ./docker/compose.yml up

db-drop-local: ## 데이터베이스 삭제(로컬에 go가 설치되어있으면)
	go run ./cmd/db/drop/main.go

db-migrate-local: ## 데이터베이스 마이그레이션(로컬에 go가 설치되어있으면)
	go run ./cmd/db/migrate/main.go

db-seed-local: ## 데이터베이스 시드(로컬에 go가 설치되어있으면)
	go run ./cmd/db/seed/main.go

db-reset-local: ## 데이터베이스 리셋(로컬에 go가 설치되어있으면)
	make db-drop-local
	make db-migrate-local
	make db-seed-local

env-local: ## 환경변수 설정(로컬에 go가 설치되어있으면)
	go run ./cmd/env/main.go

docker-down: ## 도커 컨테이너 종료
	docker compose -f ./docker/compose.yml down

help: ## 항목 목록
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'