### API - Curso GO HIQ - 

### Instalar
```shell
git clone git@github.com:filipefernandes007/curso-go-api.git
cd curso-go-api.git
```

### Correr aplicação
```shell
go run main.go
```

### Testar
```shell
curl -v -X POST -d '{"first_name":"Filipe","last_name":"Fernandes","age":48,"phone_number":"999999999"}' http://localhost:9000/api/create-client
curl http://localhost:9000/api/get-client\?uuid\=e777fd05-d9a0-42ec-b969-46036d4c791e
curl http://localhost:9000/api/list-client
```
