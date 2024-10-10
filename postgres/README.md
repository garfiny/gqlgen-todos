
## Create DB

createdb command, which is from Postgres installation, will create a db on the server running on localhost:5432, or otherwise specified.
```
createdb todos_dev -h localhost -U postgres
```

## Create migration files

```
 migrate create -ext sql -dir postgres/migrations create_actions
```

## To run the migration

### Migrate

```
migrate --path "postgres/migrations" --database "$POSTGRESSQL_URL" up
```

### Rollback
```
migrate --path "postgres/migrations" --database "$POSTGRESSQL_URL" down
```


## To seed the data

```
psql -U postgres -d todos -h localhost -a -f ./postgres/seeds.sql
```
