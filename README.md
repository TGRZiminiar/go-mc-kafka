## Start Server Need two step
#### 1. Start Docker Environment
```
docker compose -f docker-compose.db.yml up
docker compose -f docker-compose.kafka.yml up -d
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


### Command To Executute Command In Kafka
```
docker exec -it kafka-1 bash
cd ./opt/bitnami/kafka/bin
ls | grep kafka-topics.sh

# Use to create topic
kafka-topics.sh --create --topic inventory --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092
kafka-topics.sh --create --topic player --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092
kafka-topics.sh --create --topic payment --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092

# Use to list all topic
kafka-topics.sh --list --bootstrap-server localhost:9092

# Use to describe topic
kafka-topics.sh --describe --topic inventory --bootstrap-server localhost:9092

```