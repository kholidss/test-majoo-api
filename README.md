# Test Majoo API

### Prerequisite

1. Install go-migrate `https://github.com/golang-migrate/migrate` for running migration
2. Setup firebase dan download file secret, rename dengan nama `serviceAccountKey.json` dan simpan di root folder

### Migration

Run below command to run migration

```
migrate -path migration -database "mysql://user:password@tcp(host:port)/dbname?query" up
```

To create a new migration file

```
migrate create -ext sql -dir migration -seq name
```

or execure command in `Makefile`

### Setup

First install, you need to run setup.sh

```
./setup.sh
```

### Running

Run below command to run app server

```
go run cmd/api-server/main.go
```
