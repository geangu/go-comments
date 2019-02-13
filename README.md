Requisitos:

    - go
    - mongodb

Ejecutar en la Terminal:

```
$ go get -u github.com/gorilla/mux
$ go get gopkg.in/mgo.v2
$ go get gopkg.in/go-playground/validator.v9
```

Test

```
$ go test -timeout 30s -coverprofile=/tmp/coverage
```
