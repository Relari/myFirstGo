# myFirstGo

Para crear un proyecto en GO con un repositorio en GitHub se hace de la siguiente manera:

```bash
go mod init github.com/Relari/myFirstGo
```

En la creacion de la clase principal se tiene que crear una clase con la extension .go

Adicionar que es un paquete principal.

```go
package main
```

Luego crear la funcion main.

```go
func main() {
    //...
}
```

Para imprimir un mensaje en consola.

```go
println("Hello World")
```

Para ejecutar el archivo en GO

```bash
go run .\main.go
```

Para descargar una libreria de GO de un repositorio

```bash
go get github.com/google/uuid
```

Para crear el archivo ejecutable 

```bash
go build .
```

Cuando se declaran variables y no se usan, no se va a poder ejecutar la aplicacion.