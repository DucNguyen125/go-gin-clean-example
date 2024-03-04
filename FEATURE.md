# Feature

## 1. Apply clean architecture

## 2. Dependencies Injection
2.1. Auto generate DI using [wire](https://github.com/google/wire)

2.2. Auto generate mock for write unit tests using [gomock](https://github.com/uber-go/mock)

## 3. Auto generate swagger docs
Fully auto generate swagger docs from input, output golang structs using [huma](https://github.com/danielgtaylor/huma)

## 4. Apply very strict coding convention
Using githook([husky](https://github.com/automation-co/husky)) and [golangci-lint](https://github.com/golangci/golangci-lint)

## 5. Manage version migration
Using [atlas](https://github.com/ariga/atlas) and [go-migrate](https://github.com/golang-migrate/migrate), auto gen up,down file sql when modify model

## 6. Gen module boilerplate

## 7. Manage environment variable
Using [go-env](https://github.com/Netflix/go-env)

# Some small features

Pre-set adapter for MongoDB, Redis, Elasticsearch, AWS S3

Monitoring logs through ELK stack (filebeat, logstash, elasticsearch, kibana)
