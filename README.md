## Start Server Need two step
#### 1. Start Docker Environment
> docker compose -f docker-compose.db.yml up

#### 2. Select Environment File To Start Golang
> go run main.go ./env/dev/.env.(anything)
> go run main.go ./env/dev/.env.player

### Command TO Genrate Proto
```
Auth
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/auth/authPb/authPb.proto
Player
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/player/playerPb/playerPb.proto
item
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/item/itemPb/itemPb.proto
inventory
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/inventory/inventoryPb/inventoryPb.proto
```