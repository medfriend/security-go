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

## compilar el proyecto
```
go run main.go
```