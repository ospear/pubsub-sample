# pubsub sample

## Getting started

- Make topic
- Make subscription

```
cp .env.sample .env
gcloud auth application-default login

# Run subscriber
go run sub.go
# Run publisher with another terminal
go run pub.go
```
