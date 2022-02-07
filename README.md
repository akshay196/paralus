# rcloud-base
rcloud-base

## Setting up the database

### Create the initial db/user

Example for `admindb`:

``` sql
create database admindb;
create user admindbuser;
```

Now in the newly created db:

``` sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
grant execute on function uuid_generate_v4() to admindbuser;
```

*This will grant the necessary permission to the newly created user to run uuid_generate_v4()*

### Run application migrations

We use [`golang-migrate`](https://github.com/golang-migrate/migrate) to perform migrations.

#### Install [`golang-migrate`](https://github.com/golang-migrate/migrate)

``` shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

*`-tags 'postgres'` is important as otherwise it compiles without postgres support*

You can refer to the [guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for full details.

#### Run migrations

Example for `admindb`:

``` shell
export POSTGRESQL_URL='postgres://<user>:<pass>@<host>:<port>/admindb?sslmode=disable'
migrate -path ./persistence/migrations/admindb -database "$POSTGRESQL_URL" up
```

See [cli-usage](https://github.com/golang-migrate/migrate#cli-usage) for more info.