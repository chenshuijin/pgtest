# A simple webapi [![Build Status](https://assets-cdn.github.com/favicon.ico)](https://github.com/chenshuijin/pgtest)

# Table of contents

* [Installation](#installation)
* [Config](#config)
* [Example apis](#Example-apis)
  * [Init](#Init)
  * [Users](#Users)
  * [Relationships](#Relationships)
* [Howto](#howto)

## Installation

Install:

    go get gopkg.in/pg.v4
    go get github.com/gorilla/mux
    go get github.com/gorilla/context

## Config
config the postgresql server in \src\conf\conf.go
```go
package conf

import (
	"gopkg.in/pg.v4"
)

func GetDbConf() *pg.Options {
	return &pg.Options{
		Addr: "localhost:5432",
		User: "postgres",
		Password: "123456",
		Database: "webdb",
	}
}
```

## Example apis

### Init
Create the tables for the app, run once in your server at the first
```
GET /init
Example:
$curl -XGET http://localhost:8000/init
Datatable create ok!
```
### Users
list all users
```
GET /users
Example:
$curl -XGET http://localhost:8000/users
[{"Id":"1","Name":"user1","Type":"user"},{"Id":"2","Name":"user2","Type":"user"},{"Id":"3","Name":"user3","Type":"user"},{"Id":"4","Name":"user4","Type":"user"},{"Id":"5","Name":"user5","Type":"user"},{"Id":"6","Name":"user6","Type":"user"}]
```
Create a user
allowed fields:
name = string
```
POST /users
Example:
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:8000/users"
{"Id":"7","Name":"Alice","Type":"user"}
```
### Relationships
List a users all relationships
```
GET /users/{id}/relationships
Example:
$curl -XGET "http://localhost:8000/users/1/relationships"
[{"Id":"1","User_id":"2","State":"disliked","Type":"relationship"}]
```
Create/update relationship state to another user.
allowed fields:
state = "liked"|"disliked"
If two users have "liked" each other, then the state of the relationship is "matched"
```
PUT /users/{id}/relationships/{other_user_id}
Example:
$curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/1/relationships/2"
{"Id":"1","User_id":"2","State":"liked","Type":"relationship"}
```
## Howto

### windows 
	run make.bat in `cmd`
### *uix
	run make.bash in `terminal`