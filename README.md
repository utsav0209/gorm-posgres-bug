## About
- A sample repo to demonstrate a bug in [gorm](https://github.com/go-gorm/gorm) library

## Bug Info
- When using Save method to update an object with existing Primary Key it resets the `createdAt` timestamp to `0000-00-00 00:00:00`
- Expected: `CreatedAt` timestamp should not be changed
- Actual: `CreatedAt` timestamp is reset to `0000-00-00 00:00:00`


## How to reproduce:
1. setup dependencies: `go mod vendor`
2. start a postgres database using following command
```cmd
    docker run -d --name postgres_bug_check --rm  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres_bug_check -p 5432:5432 -it postgres:14.1-alpine
```
3. start the server: `go run main.go`

## Output Log:
```
❯ go run main.go

2022/07/16 16:57:05 /Users/utsav.varia/go/pkg/mod/gorm.io/driver/postgres@v1.3.8/migrator.go:167
[12.689ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'users' AND table_type = 'BASE TABLE'

2022/07/16 16:57:05 /Users/utsav.varia/go/pkg/mod/gorm.io/driver/postgres@v1.3.8/migrator.go:140
[10.659ms] [rows:0] CREATE TABLE "users" ("created_at" timestamptz,"updated_at" timestamptz,"user_id" bigserial,"name" text,PRIMARY KEY ("user_id"))

2022/07/16 16:57:05 /Users/utsav.varia/open source/gorm-postgres-bug/main.go:32
[1.786ms] [rows:0] UPDATE "users" SET "created_at"='0000-00-00 00:00:00',"updated_at"='2022-07-16 16:57:05.837',"name"='User 1' WHERE "user_id" = 1

2022/07/16 16:57:05 /Users/utsav.varia/open source/gorm-postgres-bug/main.go:32
[1.399ms] [rows:0] SELECT * FROM "users" WHERE "user_id" = 1 LIMIT 1

2022/07/16 16:57:05 /Users/utsav.varia/open source/gorm-postgres-bug/main.go:32
[1.284ms] [rows:1] INSERT INTO "users" ("created_at","updated_at","name","user_id") VALUES ('2022-07-16 16:57:05.84','2022-07-16 16:57:05.837','User 1',1) RETURNING "user_id"

2022/07/16 16:57:05 /Users/utsav.varia/open source/gorm-postgres-bug/main.go:37
[0.472ms] [rows:1] UPDATE "users" SET "created_at"='0000-00-00 00:00:00',"updated_at"='2022-07-16 16:57:05.844',"name"='User 2' WHERE "user_id" = 1

```