# Usage

- First, build and generate executable binary named as app.

`go build -o app main.go`

- Second, execute binary with different ports

`./app -listen=:8080 & ./app -listen=:8081 & ./app -listen=:8082 &`

Because we are using & operator, we are delegating this process to background. 
In order to stop running processes you can use `kill -9 {process_id}` like `kill -9 55494`