## inicializacion 

primer levantar el proyecto de getway y agregar
el servicename dentro de .env SERVICENAME=SECURITY

## instalacion 

```
go mod tidy
```

## compilacion

```
go build security-go
```

## prerequisitos
```
activar consul, rabbitmq y postgress a nivel o en docker
```

## instalar el swagger, se escribe el comando en consola
```
go install github.com/swaggo/swag/cmd/swag@latest
```

## instalar wire para la inyeccion de dependencias

```
go install github.com/google/wire/cmd/wire@latest
```
## recompilar el wire dentro de module

```
cd module && wire && cd ..
```

## recompilar el swagger cuando se hace cambios en los controladores
```
swag init
```

una vez compilado por primera vez importar el docs dentro de httpserver

_ moduleName/docs donde moduleName es el nombre del modulo de go que esta monejando los paquetes

si sale el error de leftDelim ejecutar

```
go get -u github.com/swaggo/swag
```


la url donde se encuentra el swager es la
```
http://localhost:port/swagger/index.html
```
## compilar el proyecto
```
go run main.go
```
## Air
## Le reconozca los cambios del proyecto
```
go install github.com/cosmtrek/air@v1.40.4
air -v
```
## ejecutar air
```
air
```

