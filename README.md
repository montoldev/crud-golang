Step to run :
1. use command ```go get ./...``` (optional ```go mod tidy ```)
2. use command ```run main.go```
3. import postman with crud-golang.postman_collection.json

Step to test :
1. use command ```go test ./...```
2. If it doesn't pass, you fix code or unit_test to pass.
3. remove file database/crud.db

Step create data with script sql :
1. use command ```sqlite3 crud.db```
2. use script sql ```INSERT INTO customers (name, age) VALUES ('John Doe', 30); ```
2. quit sqlite3 ```.exit```
