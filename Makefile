proto:
	@echo "--> Generating gRPC clients"
	@buf generate

run-servers:
	@echo "--> Starting servers"
	@docker compose up

run-frontend:
	@echo "--> Starting frontend"
	cd ./frontend-ts
	pnpm run dev