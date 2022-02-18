# apex-api-sync

apex-api-sync sources data from https://apexlegendsapi.com/, saving it 
to a mongo instance

Currently, this extracts:
- Players Data (current level, name, selected legend, rank, others)
- Games Data (kills, damage, played legend, ranked points won, others)

This is distributed in the way of multiple Docker images, being that each
one is responsible for updating 1 type of data.

### sync-players

This app fetches the players registered in the mongo instance and refreshes their
data

### sync-games

This app fetches the players registered in the mongo instance and saves matches played
by these users in the mongo instance

## Environment variables

|Name|Description|Default Value|
|:-|:-|:-|
|API_KEY|API Key for apexlegendsapi||
|MONGO_CONNECTION_STRING|Mongo Connection String||
|MONGO_PLAYERS_DATABASE|Mongo Players Database||
|MONGO_PLAYERS_COLLECTION|Mongo Players Collection||
|MONGO_GAMES_DATABASE|Mongo Games Database||
|MONGO_GAMES_COLLECTION|Mongo Games Collection||

## Running this app locally

1. Create a .env file with the following content:
```
API_KEY=<your-api-key>
```

2. Use compose to build and run the app you want to test (for example, `sync-players`):
```shell
docker-compose build sync-players && docker-compose run sync-players
```

3. (optional) Use `mongo-express` to see which data was written to mongo. Boot up the container
and access localhost:8081 in your browser
```shell
docker compose up -d mongo-express
```

## Deploying this app

To do this, make sure the following images are accessible to the k8s cluster (in a image repository where it
can access them):
- apex-sync-players:1.4.0
- apex-sync-games:1.4.0

If using Docker's local k8s cluster, build the images locally and they will be available:
```shell
docker build -t apex-sync-players:1.4.0 -f sync-players\Dockerfile .
docker build -t apex-sync-games:1.4.0 -f sync-games\Dockerfile .
```

Create a k8s secret named `apex-api-sync-secret` where API_KEY and MONGO_CONNECTION_STRING are defined

Apply the cron-job specs to your k8s cluster:
```shell
kubectl apply -f k8s/sync-players-cron.yaml
kubectl apply -f k8s/sync-games-cron.yaml
```
