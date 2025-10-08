🧩 1. Declaración de variables (var)

Se usa para declarar una o más variables, con o sin valor inicial.

var nombre string
var edad int = 30
var activo = true // el tipo se infiere

También puedes agruparlas:

var (
    usuario = "roberto"
    puntos  = 100
)


Y dentro de funciones, puedes usar la declaración corta:

mensaje := "Hola mundo" // equivalente a var mensaje = "Hola mundo"

🧠 2. Declaración de constantes (const)

Las constantes no pueden cambiar de valor después de ser declaradas.

const Pi = 3.1416
const (
    Puerto = 8080
    Host   = "localhost"
)

⚙️ 3. Declaración de tipos (type)

Define nuevos tipos basados en los existentes, o estructuras personalizadas.

type ID int
type Nombre string


También se usa para declarar structs, interfaces o alias:

type Persona struct {
    Nombre string
    Edad   int
}

type Lector interface {
    Leer() string
}

type ListaDeCadenas = []string // alias de tipo

🔁 4. Declaración de funciones (func)

Define funciones o métodos.

func Sumar(a int, b int) int {
    return a + b
}

func (p Persona) Saludar() string {
    return "Hola, soy " + p.Nombre
}

📦 5. Declaración de imports (import)

Importa paquetes del estándar o externos.

import "fmt"

import (
    "os"
    "strings"
)


También puedes importar con alias o anónimo:

import io "io/ioutil" // alias
import _ "github.com/lib/pq" // solo ejecutar init()

🧱 6. Declaración de paquetes (package)

Cada archivo Go comienza con una declaración de paquete:

package main

🧩 7. Declaración de bloques

Go permite agrupar declaraciones (var, const, type) en bloques:

const (
    Lunes = iota
    Martes
    Miercoles
)


🔠 Reglas de visibilidad en Go

En Go, la primera letra del nombre (de una función, variable, tipo, struct, etc.) determina si es exportado o no desde su paquete.

Nombre	Visible fuera del paquete	Descripción
🔹 Mayúscula inicial	✅ Sí (exportado)	Puede ser usado por otros paquetes.
🔸 Minúscula inicial	🚫 No (privado)	Solo puede ser usado dentro del mismo paquete.
🔍 Ejemplo
package utils

// Exportada — visible desde otros paquetes
func Sumar(a, b int) int {
    return a + b
}

// No exportada — solo visible dentro de utils
func restar(a, b int) int {
    return a - b
}


Y si en otro archivo tienes:

package main

import "miapp/utils"

func main() {
    utils.Sumar(3, 4) // ✅ OK
    utils.restar(3, 4) // ❌ Error: restar no está exportada
}

⚙️ Esto aplica también a:

Variables:

var Nombre = "Roberto" // exportada
var edad = 30           // privada


Structs y sus campos:

type Persona struct {
    Nombre string // exportado
    edad   int    // privado
}


Interfaces, tipos y constantes:

type Lector interface { Leer() error } // exportado
type escritor interface { escribir() } // privado

💡 En resumen
Prefijo	Ejemplo	Nivel de acceso
Mayúscula	Sumar, Usuario, ConfigGlobal	Público (exportado)
Minúscula	sumar, usuario, configGlobal	Privado (interno al paquete)