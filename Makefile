swag:
	swag init -q --parseDependency --parseDepth 2 -d cmd/app/,api/handler/ -o docs/swagger-ui/ --ot "json"
docker_build:
	sudo docker build -t getir_case .
docker_run:
	docker run -p 127.0.0.1:8080:8080 getir_case