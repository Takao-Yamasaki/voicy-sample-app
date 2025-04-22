.DEFAULT_GOAL := help
up: ## 全コンテナを起動する
	docker compose up -d
build: ## appコンテナをビルドする
	docker compose up app -d --build
app: ## appコンテナに接続する
	docker compose exec -it app sh
down: ## 全コンテナを停止する
	docker compose down
ps: ## コンテナの稼働状況を確認する
	docker compose ps
test: ## テストを実行する
	go test -v ./handlers
help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
