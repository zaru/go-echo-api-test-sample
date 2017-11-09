# Golang WebAPI server testing sample

- WAF: echo v3
- DB: PostgreSQL
  - sqlx
- Test tool
  - stretchr/testify/assert
  - DATA-DOG/go-sqlmock

## Init

### Docker

```
docker-compose up -d
```

### DB schema

```
create table users (
  id SERIAL,
  name varchar
) ;
```
### Start echo server

```
realize start
```

### Test

```
go test ./...
```
