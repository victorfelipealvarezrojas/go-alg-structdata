# Go: Genéricos, Interfaces e Implementación Implícita

## Resumen de Conceptos Aprendidos

### 1. Tipos y Alias de Tipos

#### Tipo Nuevo (sin `=`)
```go
type IntegerType int  // Crea un NUEVO tipo basado en int
```
- `IntegerType` es **diferente** a `int`
- No son compatibles directamente
- Tienes que hacer conversión explícita

#### Alias de Tipo (con `=`)
```go
type IStringer = interface {
    fnString() string
}
```
- `IStringer` es **exactamente lo mismo** que la interfaz
- Son completamente compatibles
- No creas un tipo nuevo, solo un alias

---

### 2. Interfaces e Implementación Implícita

#### Definir una Interfaz
```go
type IStringer = interface {
    fnString() string
}
```

#### Implementar la Interfaz
```go
type ageType int

func (i ageType) fnString() string {
    return fmt.Sprintf("%d", i)
}
```

**Lo especial de Go: Implementación Implícita**

- No necesitas declarar `implements IStringer`
- Go **infiere automáticamente** que `ageType` implementa `IStringer`
- Solo necesitas tener el método con la firma exacta

**Comparación con otros lenguajes:**
```java
// Java - explícito
class ageType implements IStringer {
    public String fnString() { ... }
}
```

```go
// Go - implícito (Go lo infiere solo)
func (i ageType) fnString() string { ... }
```

---

### 3. Receptor del Método (Method Receiver)

```go
func (i ageType) fnString() string {
    return fmt.Sprintf("%d", i)
}
```

**El receptor `(i ageType)` es:**
- **No es un parámetro** de entrada
- Es lo que **vincula el método al tipo**
- `i` es como `this` en Java/C# (acceso a la instancia)
- `ageType` especifica a **cuál tipo pertenece** este método

**Acceso a propiedades:**
```go
type StudentType struct {
    Name string
    ID   int
    Age  float64
}

func (s StudentType) fnString() string {
    return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
    // s tiene acceso a todas las propiedades
}
```

---

### 4. Genéricos con Restricción de Tipo

```go
func addStudent[T IStringer](students []T, student T) []T {
    return append(students, student)
}
```

**Significa:**
- `T` es un tipo genérico
- `T` **debe implementar** `IStringer`
- Solo tipos que tengan el método `fnString() string` pueden usarse

**Uso:**
```go
students := []ageType{}
result := addStudent(students, 45)  // Funciona porque ageType implementa IStringer
```

---

### 5. Genéricos sin Restricción (`any`)

```go
func addStudent[T any](students []T, student T) []T {
    return append(students, student)
}
```

**Significa:**
- `T` acepta **cualquier tipo**, sin restricciones
- `string`, `int`, `StudentType`, lo que sea

**Ventaja:** Más flexible
**Desventaja:** Pierdes seguridad de tipo — no sabes qué puedes hacer con `T`

---

### 6. Go No es Orientado a Objetos

Go es un lenguaje **funcional con características de objetos**:

- **No tiene clases** — tiene tipos
- **Los métodos son funciones** con receptor
- **Todo es una función** en el fondo
- Énfasis en **composición sobre herencia**

**Comparación:**
```java
// Java - Orientado a Objetos
class ageType {
    public String fnString() { ... }  // Método dentro de la clase
}
```

```go
// Go - Funcional
func (i ageType) fnString() string { ... }  // Función con receptor
```

---

### 7. Exportación de Tipos (Paquetes)

#### En el paquete (generic.go)
```go
package generic

type StudentType struct {  // Mayúscula = EXPORTADO
    Name string
    ID   int
    Age  float64
}

func AddStudent[T any](...) { ... }  // Mayúscula = EXPORTADO
```

#### En main.go
```go
import "algstruc/ch1/generic"

students := []generic.StudentType{}  // Acceso con prefijo
result := generic.AddStudent(...)    // Acceso con prefijo
```

**Regla:**
- **Mayúscula** → exportado (visible desde otros paquetes)
- **Minúscula** → no exportado (solo dentro del paquete)
- **Siempre usar prefijo** `paquete.Tipo` (es la convención de Go)

---

### 8. Struct Literals con Campos Nombrados

#### ❌ Sin nombres (no recomendado)
```go
generic.StudentType{"John", 213, 17.5}
```

#### ✅ Con nombres (recomendado)
```go
generic.StudentType{
    Name: "John",
    ID: 213,
    Age: 17.5,
}
```

**Por qué:**
- Más legible
- Si el struct cambia, no se rompe tu código
- Práctica estándar de Go

---

## Comparación: Go vs Otros Lenguajes

| Concepto | Go | Java | JavaScript |
|----------|----|----|-------------|
| Interfaces | Implícitas | Explícitas | No existen |
| Métodos | Funciones con receptor | Dentro de clase | Métodos en prototipo |
| Herencia | No existe | Sí | Prototipo |
| Genéricos | `[T constraint]` | `<T extends>` | No oficialmente |
| OOP | No puro | Puro | Prototipo |

---

## Lo Raro de Go (Viniendo de Java/C#/JS)

1. **Implementación implícita de interfaces** — No declaras `implements`
2. **Métodos fuera de la clase** — Receptores explícitos `(i ageType)`
3. **Tipos nuevos sin herencia** — `type X int` crea un tipo nuevo
4. **Alias de tipos con `=`** — Distinción entre alias y tipo nuevo
5. **No es OOP tradicional** — Es funcional + composición

---

## Código de Ejemplo Completo

```go
package generic

import "fmt"

// Interfaz (alias)
type IStringer = interface {
    fnString() string
}

// Tipos que implementan IStringer
type ageType int

func (i ageType) fnString() string {
    return fmt.Sprintf("%d", i)
}

type nameType string

func (s nameType) fnString() string {
    return string(s)
}

type StudentType struct {
    Name string
    ID   int
    Age  float64
}

func (s StudentType) fnString() string {
    return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

// Función genérica con restricción
func AddStudent[T IStringer](students []T, student T) []T {
    return append(students, student)
}
```

---

## Puntos Clave para Recordar

1. **Go infiere tipos** — Si tienes el método correcto, implementas la interfaz
2. **Receptor = `this`** — El receptor te da acceso a la instancia
3. **Genéricos = flexibilidad** — Reutiliza código para múltiples tipos
4. **Interfaz implícita = desacoplamiento** — Menos dependencias entre paquetes
5. **Go es funcional** — Piensa en transformaciones de datos, no en objetos
6. **Siempre usa mayúscula para exportar** — Es la convención
7. **Usa campos nombrados en structs** — Es la práctica recomendada

---

## Siguiente Paso: Rust

Una vez domines Go, aprende Rust para:
- Expandir tu mente con ownership y borrowing
- Aprender funcional más puro
- Tener oportunidades laborales reales

**Stack futuro:** Go (backend) + Rust (sistemas/performance)