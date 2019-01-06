# togo

`Golang` generator handler for SQL, HTTP, i18n, SQL DDL etc.

# Dependencies

Install all dependencies before proceeding.

* `go get github.com/nicksnyder/go-i18n/v2/i18n`

## Installing `togo`

```bash

cd togo/
go install ./...

```

## Generating SQL binding objects from files

Ensure `togo` is installed as described above. Then using `samples/sql` as an example;

Open `sql.go` and ensure `go:generator` definition is specified as shown below:

> //go:generate togo sql

..then from the terminal do;

```bash

cd togo/samples/sql
rm sql_gen.go

go generate
```

The is ensures that the SQL definitions in `files/` are all loaded and converted into SQL bindings, in a new generated file called `sql_gen.go`
