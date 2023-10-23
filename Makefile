swag:
	swag init -q --parseDependency --parseDepth 2 -d cmd/app/,api/handler/ -o docs/swagger-ui/ --ot "json"
docker_build:
	docker compose up -d
	docker compose up --build
docker_run:
	docker compose up