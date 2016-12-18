# kipapi

<p align="center">
<a href="https://travis-ci.org/fulldump/kipapi"><img src="https://travis-ci.org/fulldump/kipapi.svg?branch=master"></a>
<a href="https://goreportcard.com/report/fulldump/kipapi"><img src="http://goreportcard.com/badge/fulldump/kipapi"></a>
<a href="https://godoc.org/github.com/fulldump/kipapi"><img src="https://godoc.org/github.com/fulldump/kipapi?status.svg" alt="GoDoc"></a>
</p>

Kipapi is a CRUD HTTP API.


<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [How to use](#how-to-use)
	- [Create and expose CRUD via HTTP API](#create-and-expose-crud-via-http-api)
- [Kipapi developer](#kipapi-developer)
	- [Run tests](#run-tests)
	- [Coverage](#coverage)

<!-- /MarkdownTOC -->


# How to use

## Create and expose CRUD via HTTP API

```go
mydao := NewDaoUsers()  // Build DAO with Kip
myapi := golax.NewApi() // Create golax api
kipapi.New(myapi.Root, mydao) // Expose
```


# Kipapi developer


## Run tests

```sh
make setup && make test
```

## Coverage

```sh
make setup && make coverage
```

