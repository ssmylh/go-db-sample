# go-db-sample

This project is Go database/sql various samples which use SQLite 3.
The following operaions are written:

- create table
- insert with transaction
- select all
- select one by record's id

The definition of sample table is:

```
create table samples (id integer not null primary key, misc text);
```

## Dependency

- For all sample

```
go get github.com/mattn/go-sqlite3
```

- For genmai sample

```
go get github.com/naoina/genmai
```

- For gorp sample

```
go get gopkg.in/gorp.v1
```

- For argen sample

```
go get -u github.com/monochromegane/argen/...
```

## References

- [Package sql](http://golang.org/pkg/database/sql/)
- [go-sqlite3/_example/simple.go](https://github.com/mattn/go-sqlite3/blob/0cdea24bc72fac013abf416f27acd433e5906528/_example/simple/simple.go)
- [genmai](https://github.com/naoina/genmai)
- [gorp](https://github.com/go-gorp/gorp)
- [argen](https://github.com/monochromegane/argen)