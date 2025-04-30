dockerup:
	docker compose --env-file=./configs/.env --env-file=./configs/.docker_env up -d

terminalup:
	go run ./cmd/gojwt/main.go