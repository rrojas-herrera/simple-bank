1. Tipos de datps

    #### Define nuevos tipos basados en los existentes, o estructuras personalizadas.

    ```
    type ID int
    type Nombre string
    ```

    ```
    type Persona struct {
        Nombre string
        Edad   int
    }
    ```

    ```
    type Lector interface {
        Leer() string
    }
    ```

    ```
    type ListaDeCadenas = []string // alias de tipo
    ```

1. Var
    #### Se usa para declarar una o más variables, con o sin valor inicial.
    ```
    var nombre string
    var edad int = 30
    var activo = true // el tipo se infiere
    ```
    ```
    var (
        usuario = "roberto"
        puntos  = 100
    )
    ```
    ```
    mensaje := "Hola mundo" // equivalente a var mensaje = "Hola mundo"
    ```

2. Const
    #### Las constantes no pueden cambiar de valor después de ser declaradas.
    ```
    const Pi = 3.1416
    const (
        Puerto = 8080
        Host   = "localhost"
    )
    ```


4. Declaración de funciones (func)

    #### Define funciones o métodos.
    ```
    func Sumar(a int, b int) int {
        return a + b
    }
    ```
    ```
    func (p Persona) Saludar() string {
        return "Hola, soy " + p.Nombre
    }
    ```

5. Imports (import)

    #### Importa paquetes del estándar o externos.

    ```
    import "fmt"
    ```

    ```
    import (
        "os"
        "strings"
    )
    ```

    ```
    import io "io/ioutil" // alias
    import _ "github.com/lib/pq" // solo ejecutar init()
    ```


6. Packages (package)
    #### Cada archivo Go comienza con una declaración de paquete:

    ```
    package main
    ```

7. Visibilidad en Go
    #### La primera letra del nombre (de una función, variable, tipo, struct, etc.) determina si es exportado o no desde su paquete.


    - 🔹 Mayúscula inicial	✅ Sí (exportado)	Puede ser usado por otros paquetes.
        ```
        package utils

        func Sumar(a, b int) int {
            return a + b
        }
        ```
    - 🔸 Minúscula inicial	🚫 No (privado)	Solo puede ser usado dentro del mismo paquete.
        ```
        package utils

        func restar(a, b int) int {
            return a - b
        }
        ```

    - En otro archivo:
        ```
        package main

        import "miapp/utils"

        func main() {
            utils.Sumar(3, 4) // ✅ OK
            utils.restar(3, 4) // ❌ Error: restar no está exportada
        }
        ```
    - Esto aplica también a:
        - Variables:
            ```
            var Nombre = "Roberto" // exportada
            var edad = 30           // privada
            ```
        - Structs y sus campos:
            ```
            type Persona struct {
                Nombre string // exportado
                edad   int    // privado
            }
            ```
        - Interfaces, tipos y constantes:
            ```
            type Lector interface { Leer() error } // exportado
            type escritor interface { escribir() } // privado
            ```
        - En resumen
            ```
            Prefijo	Ejemplo	Nivel de acceso
            Mayúscula	Sumar, Usuario, ConfigGlobal	Público (exportado)
            Minúscula	sumar, usuario, configGlobal	Privado (interno al paquete)
            ```



