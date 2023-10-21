## Start Server Need two step
#### 1. Start Docker Environment
> docker compose -f docker-compose.db.yml up

#### 2. Select Environment File To Start Golang
> go run main.go ./env/dev/.env.(anything)
> go run main.go ./env/dev/.env.player