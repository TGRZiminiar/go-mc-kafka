## Start Server Need two step
#### 1. Start Docker Environment

```
docker compose -f docker-compose.db.yml up -d
docker compose -f docker-compose.db.yml down
docker compose -f docker-compose.kafka.yml up -d
```

<br/>
<hr/>
<br/>

#### 2. Select Environment File To Start Golang
go run main.go ./env/dev/.env.(anything)
go run main.go ./env/dev/.env.player
go run main.go ./env/dev/.env.item
go run main.go ./env/dev/.env.auth
go run main.go ./env/dev/.env.payment
go run main.go ./env/dev/.env.inventory

<br/>
<hr/>
<br/>

### Command TO Genrate Proto
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

<h2>🍰 Generate a Proto File Command</h2>
<p>player</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/player/playerProto/playerProto.proto
```

<p>auth</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/auth/authProto/authProto.proto
```

<p>item</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/item/itemProto/itemProto.proto
```

<p>inventory</p>

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/inventory/inventoryProto/inventoryProto.proto
```
<br/>
<hr/>
<br/>

### Command To Migrate
```
go run ./pkg/database/script/migration.go ./env/dev/.env.auth && \
go run ./pkg/database/script/migration.go ./env/dev/.env.player && \
go run ./pkg/database/script/migration.go ./env/dev/.env.item && \
go run ./pkg/database/script/migration.go ./env/dev/.env.payment && \
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


<h2>🦋 Kafka Command</h2>

<p>Start Docker Compose for Kafka</p>

```bash
docker compose -f docker-compose.kafka.yml up -d
```

<p>Enter into the Kafka container</p>

```bash
docker exec -it kafka-1 bash
```

<p>Create a topic</p>

```bash
./opt/bitnami/kafka/bin/kafka-topics.sh --create --topic inventory --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-topics.sh --create --topic payment --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-topics.sh --create --topic player --replication-factor 1 --partitions 1 --bootstrap-server localhost:9092
```

<p>Add topic retention</p>

```bash
./opt/bitnami/kafka/bin/kafka-configs.sh --bootstrap-server localhost:9092 --entity-type topics --entity-name inventory --alter --add-config retention.ms=180000
./opt/bitnami/kafka/bin/kafka-configs.sh --bootstrap-server localhost:9092 --entity-type topics --entity-name payment --alter --add-config retention.ms=180000
./opt/bitnami/kafka/bin/kafka-configs.sh --bootstrap-server localhost:9092 --entity-type topics --entity-name player --alter --add-config retention.ms=180000
```

<p>See all topics list</p>

```bash
./opt/bitnami/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

<p>Describe topic</p>

```bash
./opt/bitnami/kafka/bin/kafka-topics.sh --describe --topic inventory --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-topics.sh --describe --topic payment --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-topics.sh --describe --topic player --bootstrap-server localhost:9092
```

<p>Write a message into the topic</p>

```bash
./opt/bitnami/kafka/bin/kafka-console-producer.sh --topic inventory --bootstrap-server localhost:9092 
./opt/bitnami/kafka/bin/kafka-console-producer.sh --topic payment --bootstrap-server localhost:9092 
./opt/bitnami/kafka/bin/kafka-console-producer.sh --topic player --bootstrap-server localhost:9092
```

<p>Write a message with key into the topic</p>

```bash
--property "key.separator=:" --property "parse.key=true"
```

<p>Read a message on that topic</p>

```bash
./opt/bitnami/kafka/bin/kafka-console-consumer.sh --topic inventory --from-beginning --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-console-consumer.sh --topic payment --from-beginning --bootstrap-server localhost:9092
./opt/bitnami/kafka/bin/kafka-console-consumer.sh --topic player --from-beginning --bootstrap-server localhost:9092
```

<p>Delete topic</p>

```bash
./opt/bitnami/kafka/bin/kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic inventory
./opt/bitnami/kafka/bin/kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic payment
./opt/bitnami/kafka/bin/kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic player
```