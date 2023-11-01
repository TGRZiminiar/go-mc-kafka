## Start Server Need two step
#### 1. Start Docker Environment
```
docker compose -f docker-compose.db.yml up
```

#### 2. Select Environment File To Start Golang
> go run main.go ./env/dev/.env.(anything)
> go run main.go ./env/dev/.env.player

### Command TO Genrate Proto
```
export PATH="$PATH:$(go env GOPATH)/bin"
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

### Command To Migrate
```
go run ./pkg/database/script/migration.go ./env/dev/.env.auth
go run ./pkg/database/script/migration.go ./env/dev/.env.player
go run ./pkg/database/script/migration.go ./env/dev/.env.item
go run ./pkg/database/script/migration.go ./env/dev/.env.payment
go run ./pkg/database/script/migration.go ./env/dev/.env.inventory

```