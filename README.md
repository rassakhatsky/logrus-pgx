# Logrus hook for Pgx library <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />[![Build Status](https://travis-ci.org/rassakhatsky/logrus-pgx.svg?branch=master)](https://travis-ci.org/rassakhatsky/logrus-pgx)
Allows to use [Logrus](https://github.com/sirupsen/logrus) logger with [Pgx](https://github.com/jackc/pgx) package.

## Installation
It is go gettable
  $ go get github.com/andygrunwald/go-trending

## Usage
Just import it.
```go
import "github.com/rassakhatsky/logrus-pgx"
```
And assign the logrus-pgx logger as default logger for Pgx.
```go
var dbConfig pgx.ConnPoolConfig
dbConfig.Host = "some_host"
dbConfig.User = "some_user"
dbConfig.Password = "some_password"
dbConfig.Database = "some_db"
dbConfig.Logger = (*logrus_pgx.PgxLogger)(logrus.StandardLogger())

pgxPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{})
```
