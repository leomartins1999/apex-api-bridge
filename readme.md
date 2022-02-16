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

## Deploying this app

Create a k8s secret named `apex-api-sync-secret` where API_KEY and MONGO_CONNECTION_STRING are defined

Apply the cronjob spec to your k8s cluster:
`kubectl apply -f k8s\cron-job.yaml`
