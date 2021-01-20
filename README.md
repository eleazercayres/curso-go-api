### API - Curso GO HIQ - 

### Instalar
```shell
git clone git@github.com:filipefernandes007/curso-go-api.git
cd curso-go-api.git
```

### Correr aplicação
```shell
go run main
```

### Testar
```shell
curl -v -X POST -d '{"first_name":"Filipe","last_name":"Fernandes","age":48,"phone_number":"999999999"}' http://localhost:9000/api/create-client
```
