up:
	docker compose up -d
app:
	docker compose up app -d --build
exec:
	docker compose exec -it app sh
down:
	docker compose down
ps:
	docker compose ps
