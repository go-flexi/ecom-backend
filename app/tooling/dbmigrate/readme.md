## Library to migrate database

[https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)

command to create migration: run this command inside the tooling/dbmigration

```
migrate create -ext sql -dir migrations -seq create_users_table
```