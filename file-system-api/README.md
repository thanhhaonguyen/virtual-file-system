# file-system-api 

This project uses Go, Gin Web Framework and Ginkgo to implement Unit Test.

If you want to learn more about Gin, please visit its website: https://github.com/gin-gonic/gin .

## Install Go and setup Go workspace first

Docs: https://go.dev/doc

## Code structure
This project includes:
- ```controllers:``` handle APIs logic
- ```middlewares:``` enable CORS with various options 
- ```models:``` include Postgres DB connection and models
- ```utils:``` handle general errors in the app
- ```tests:``` implement Unit Test

## Using PostgreSQL
```
folder(id, name, parent_id, created_at)
file(id, name, data, folder_id, created_at)
```

## APIs
BaseURL: https://virtual-file-system-v1.herokuapp.com/
### Folder
```
[GET]     "/folder"
[GET]     "/folder-by-parent/:id"
[GET]     "/folder/:id"
[POST]    "/folder"
[PUT]     "/folder/:id"
[DELETE]  "/folder/:id"
```
### File
```
[GET]     "/file"
[GET]     "/file/:id"
[POST]    "/file"
[PUT]     "/file/:id"
[DELETE]  "/file/:id"
```

## Running the application and visit "localhost:8080" on browser
```shell script
go run main.go
```

## Running Unit Test
```shell script
go test -v -cover ./tests/... ./controllers/v1/...
```


