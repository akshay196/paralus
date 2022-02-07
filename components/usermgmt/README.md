# User management

This module is responsible for manging users/groups/roles as well as being a frontend for the casbin internal service.
All user/auth related requests go through here.


## Development

### Run kratos migrations

``` shele
export DSN='postgres://postgres@localhost:5432/admindb?sslmode=disable'
kratos -c kratos.yaml migrate sql -e --yes
```

### Start kratos

``` shell
cd components/usermgmt/_kratos
kratos serve -c kratos.yml
```

### Run usermgmt server

``` shell
cd components/usermgmt/_kratos
go run main.go
```