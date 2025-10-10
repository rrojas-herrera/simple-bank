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