# apex-api-sync

This app updates a mongo instance with the players most recent data.
Player data is sourced from https://apexlegendsapi.com/

## Environment variables

|Name|Description|Default Value|
|:-|:-|:-|
|API_KEY|API Key for apexlegendsapi||
|MONGO_CONNECTION_STRING|Mongo Connection String||
|MONGO_DATABASE|Mongo Database||
|MONGO_COLLECTION|Mongo Collection||

## Running this app locally

1. Create a .env file with the following content:
```
API_KEY=<your-api-key>
```

2. Build the app's image and run int
```shell
docker compose build && docker compose run app
```

3. (Optional) to validate the data written to mongo, start the mongo-express container and access localhost:8081 in your browser
```shell
docker compose up -d mongo-express
```

## Deploying this app

Create a k8s secret named `apex-api-sync-secret` where API_KEY and MONGO_CONNECTION_STRING are defined

Apply the cronjob spec to your k8s cluster:
`kubectl apply -f k8s\cron-job.yaml`
