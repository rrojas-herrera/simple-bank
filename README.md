# 🦫 Guía Completa para Aprender Go (Golang)

## 📘 Introducción

Go (o Golang) es un lenguaje de programación creado por Google, enfocado en la **simplicidad**, **eficiencia**, y **concurrencia**.  
Es ampliamente utilizado en el desarrollo de **microservicios**, **herramientas CLI**, y **backend de alto rendimiento**.

---

## 🚀 Instalación y Configuración

### 🔹 Instalar Go

Descarga Go desde la página oficial:  
👉 [https://go.dev/dl/](https://go.dev/dl/)

### 🔹 Verificar instalación

```bash
go version
```

### 🔹 Configurar el entorno de trabajo

Asegúrate de tener configuradas las variables de entorno:

```bash
go env
```

Tu proyecto debe tener esta estructura:

```
mi-proyecto/
├── go.mod
├── main.go
└── pkg/
```

Crea el módulo:

```bash
go mod init mi-proyecto
```

---

## 🧱 Conceptos Básicos

### ✳️ Declaración de variables

```go
var nombre string = "Roberto"
edad := 30 // inferencia de tipo
```

### ✳️ Tipos de datos primitivos

| Tipo        | Descripción                 | Ejemplo           |
|--------------|----------------------------|--------------------|
| `string`     | Cadenas de texto           | `"Hola"`          |
| `int`        | Enteros                    | `10`              |
| `float64`    | Números decimales          | `3.14`            |
| `bool`       | Booleanos (true/false)     | `true`            |

---

## 🧩 Estructuras de Control

### 🔹 Condicionales

```go
if edad >= 18 {
    fmt.Println("Eres mayor de edad")
} else {
    fmt.Println("Eres menor de edad")
}
```

### 🔹 Bucles

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 🔹 Switch

```go
switch dia := 3; dia {
case 1:
    fmt.Println("Lunes")
case 2:
    fmt.Println("Martes")
default:
    fmt.Println("Otro día")
}
```

---

## 🧱 Funciones

```go
func Sumar(a int, b int) int {
    return a + b
}
```

### 🔸 Retornos múltiples

```go
func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return a / b, nil
}
```

---

## 🧍‍♂️ Estructuras (Structs)

```go
type Persona struct {
    Nombre string
    Edad   int
}
```

### 🔹 Crear e imprimir

```go
p := Persona{Nombre: "Roberto", Edad: 32}
fmt.Printf("Persona: %+v\n", p)
```

---

## 🪄 Métodos

```go
func (p Persona) Saludar() {
    fmt.Println("Hola, soy " + p.Nombre)
}
```

---

## 🧰 Paquetes e Importaciones

```go
import (
    "fmt"
    "time"
)
```

Usar una librería externa:

```bash
go get github.com/stretchr/testify
```

---

## 🪵 Conversión de Tipos

```go
import "strconv"

i := int64(42)
s := strconv.FormatInt(i, 10)
fmt.Println("El número es " + s)
```

Otros ejemplos:

| De → A | Función |
|---------|----------|
| `int` → `string` | `strconv.Itoa(x)` |
| `string` → `int` | `strconv.Atoi(s)` |
| `float64` → `string` | `strconv.FormatFloat(f, 'f', -1, 64)` |

---

## 🖨️ Imprimir en consola

```go
fmt.Println("Hola mundo")
fmt.Printf("Hola, %s\n", nombre)
```

Ejemplo con formato legible:

```go
func PrintStructColor(title string, data any) {
    fmt.Printf("\033[1;36m%s\033[0m\n", title)
    if data != nil {
        b, _ := json.MarshalIndent(data, "", "  ")
        fmt.Println(string(b))
    }
}
```

Resultado:

```
detalle:
{
  "ID": 87,
  "Nombre": "Roberto"
}
```

---

## ⚙️ Concatenación de cadenas

```go
saludo := "Hola, soy " + p.Nombre
```

También se puede usar `fmt.Sprintf`:

```go
mensaje := fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
```

---

## 🧠 Concurrencia

Go destaca por su modelo de concurrencia con **goroutines** y **canales**.

```go
func decirHola() {
    fmt.Println("Hola desde goroutine")
}

func main() {
    go decirHola()
    fmt.Println("Hola desde main")
    time.Sleep(time.Second)
}
```

---

## 🧪 Pruebas (Testing)

Crea un archivo `main_test.go`:

```go
package main

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestSumar(t *testing.T) {
    resultado := Sumar(2, 3)
    require.Equal(t, 5, resultado)
}
```

Ejecuta:

```bash
go test ./...
```

---

## 🐳 Go con Docker

Ejemplo de `Dockerfile`:

```dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
```

---

## 📚 Recursos Recomendados

- [Tour of Go (oficial)](https://go.dev/tour/)
- [Documentación oficial](https://pkg.go.dev/)
- [Go by Example](https://gobyexample.com/)
- Libro: *The Go Programming Language* (Alan Donovan & Brian Kernighan)

---

## 🧭 Siguientes pasos

1. Practica con ejercicios simples (`FizzBuzz`, `calculadora`, etc.)
2. Aprende a manejar **errores** (`error handling`)
3. Usa **interfaces** y **structs anidados**
4. Implementa **API REST** con Go (por ejemplo con *Fiber* o *Gin*)
5. Aprende sobre **concurrencia avanzada** (canales, select, mutex)

---

# 🧩 Curso Paso a Paso con Ejercicios

## 🏁 Nivel 1: Fundamentos

**Ejercicio 1:**  
Crea un programa que imprima “Hola, Go!” en la consola.

**Ejercicio 2:**  
Declara una variable `nombre` y muestra un mensaje que diga “Hola, [nombre]”.

**Ejercicio 3:**  
Crea una función `Sumar(a, b int)` que retorne la suma de dos números.

---

## 🧱 Nivel 2: Structs y Métodos

**Ejercicio 1:**  
Define un struct `Persona` con los campos `Nombre` y `Edad`.

**Ejercicio 2:**  
Agrega un método `Saludar()` que imprima un saludo personalizado.

---

## ⚙️ Nivel 3: Concurrencia

**Ejercicio 1:**  
Crea una función que imprima números del 1 al 5.  
Ejecuta dos goroutines simultáneamente.

**Ejercicio 2:**  
Utiliza canales (`chan`) para comunicar goroutines entre sí.

---

## 🧪 Nivel 4: Pruebas Unitarias

**Ejercicio:**  
Crea un archivo de prueba que valide el resultado de tu función `Sumar` usando `testify`.

---

## 🐘 Nivel 5: Base de Datos

Conéctate a PostgreSQL usando `database/sql` y el driver `pq`.

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres password=123 dbname=test sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
}
```

---

## 🧠 Nivel 6: API REST con Gin

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.Run(":8080")
}
```
