# GoLang Service
## (GoLang , CHI, GORM, SQLite3)
This is REST API service that overs the following concepts
- setup a CHI REST API service sertup (with/without Docker)
- Basic Model, Schema creation in GORM
- Basic CRUD operations using ORM objects
- Scanning a directory with subdirectories and files to capture keywords

**NOTE** : This codebase was developed on go1.18.1

## Technologies
[GoLang](https://golangbyexample.com)

[CHI REST Framework](https://go-chi.io/#/README)

[GORM](https://gorm.io/index.html)

## Local Setup
- initiate modules , run `go mod init <example.com/reponame>`
- install dependencies from go.mod `go mod download`
- install dependancies individually , example run `go get -u github.com/go-chi/chi/v5`
- create .env file from .env-dist with appropriate local values
- run the service `go run .`
- run unit tests `go test ./... -v -cover`

## File Structure

**main.go** - this is the entry point. handler mounted here will serve the REST APIS exposed by this service

**handler** - All the routes and controller methods to the API endpoints are defined here

**models** - All the Database entities are declared here

**db** - Database connection to sqlite3 and helper methods required to perform CRUD operations on DB entities are defined here

**temp** - temporary directory to place cloned public repos, each repo will cloned and remvoed programatically while scanning a public repository

**repos_local.db** - sqlite3 DB file

**mock_http** - mocks for unit testing

## References
- https://go-chi.io/#/pages/getting_started
- https://gorm.io/docs/

## API Documentation 
Postman Collection - GuardRails Service - APIs.postman_collection.json
PDF - API Documentation.pdf