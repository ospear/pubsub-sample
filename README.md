# Cloud Pub/Sub Sample

## Getting started

- gcloud auth

```shell script
gcloud auth application-default login
```

- set GCP project

```
gcloud config set project your_project
```

- Make topic

```shell script
gcloud pubsub topics create test
```

- Make subscription

```shell script
gcloud pubsub topics create test-subscription
```

- Edit .env

```
cp .env.sample .env
# Replace GCP_PROJECT
```

## Usage

```shell script
# Run subscriber
go run sub.go

# Run publisher with another terminal
go run pub.go
```

## Links

- https://cloud.google.com/pubsub/docs/publisher?hl=ja
- https://cloud.google.com/pubsub/docs/pull?hl=ja
