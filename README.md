[README.md](https://github.com/user-attachments/files/33710389/README.md)
# Portafolio IT en Go · versión 2 (capa 2)

Portafolio web de un técnico de soporte IT **ficticio** (Álex Navarro Ruiz), escrito en **Go** usando solo la biblioteca estándar y ejecutado dentro de un **contenedor Docker**.

Proyecto académico de la asignatura **FP.443 (P) Devops y cloud computing** (CFGS ASIR, UOC), Producto 1: *Preparación del entorno de desarrollo*.

> **Datos ficticios.** La persona, las empresas y el correo son inventados. El correo usa `example.com`, un dominio reservado para ejemplos (RFC 2606).

---

## Qué hace esta versión

| Requisito del enunciado | Cómo se cumple |
|---|---|
| Web en Go que muestre «Soy alumno de la UOC» | El lema aparece en la cabecera de la página |
| Mostrar una imagen | Avatar ilustrado propio (`static/img/avatar.svg`) |
| Contenedor Docker con la aplicación | `Dockerfile` (versión 1, de una etapa) |

La página tiene seis bloques: cabecera con avatar, **Sobre mí**, **Experiencia**, **Formación**, **Habilidades** y **Contacto**. Se adapta a pantallas de móvil.

---

## Estructura del proyecto

```
portafolio-go/
├── go.mod                 # Módulo Go y versión mínima del lenguaje (1.24)
├── main.go                # Servidor: rutas, plantilla y ficheros estáticos
├── perfil.go              # Datos del perfil (structs + contenido ficticio)
├── Dockerfile             # Receta de la imagen Docker
├── README.md              # Este archivo
├── templates/
│   └── index.html         # Plantilla HTML (html/template)
└── static/
    ├── css/estilo.css     # Estilos: paleta Google Drive
    └── img/avatar.svg     # Imagen que exige el enunciado
```

Cada archivo tiene un único trabajo:

- **Datos** → `perfil.go`
- **Presentación** → `templates/index.html` y `static/css/estilo.css`
- **Servidor** → `main.go`

Por eso el rediseño visual de esta versión no ha necesitado cambiar ninguna línea de Go.

---

## Requisitos

| Herramienta | Versión usada | Para qué |
|---|---|---|
| Go | 1.27.1 (mínimo 1.24) | Ejecutar y compilar en local |
| Docker | Docker Desktop 4.81.0 (motor 29.6.1) | Construir la imagen y ejecutar el contenedor |
| Entorno | Windows 11 + WSL 2 (Ubuntu) | El proyecto vive en `~/proyectos/portafolio-go`, dentro de Linux |

> **Importante en WSL:** el proyecto debe estar en el sistema de archivos de Linux (`~`) y no en `/mnt/c/...`. Así se evita la lentitud de acceder a los archivos de Windows desde Linux.

---

## Ejecutar en local (sin Docker)

```bash
cd ~/proyectos/portafolio-go
go run .
```

Abre <http://localhost:8080> y para el servidor con **Ctrl+C**.

> Prueba siempre en local antes de contenerizar. Si la aplicación falla aquí, también fallará dentro del contenedor.

---

## Ejecutar con Docker

```bash
# 1. Construir la imagen (-t = nombre:etiqueta, "." = contexto del build)
docker build -t portafolio-go:v2 .

# 2. Eliminar el contenedor anterior si existe
docker rm -f portafolio

# 3. Arrancar el contenedor (-d = segundo plano, -p = puerto del equipo:puerto del contenedor)
docker run -d --name portafolio -p 8080:8080 portafolio-go:v2

# 4. Comprobar
docker ps
docker logs portafolio
```

Abre <http://localhost:8080>.

| Problema | Causa probable |
|---|---|
| `port is already allocated` | Sigue activo un `go run` u otro contenedor en el puerto 8080 |
| `container name "/portafolio" is already in use` | Ejecuta `docker rm -f portafolio` |
| `docker` no encontrado en Ubuntu | Docker Desktop está parado o la integración WSL está desactivada |

---

## Rutas del servidor

| Método y ruta | Respuesta |
|---|---|
| `GET /` | Página del portafolio (HTML generado con la plantilla) |
| `GET /static/...` | CSS e imagen (`text/css`, `image/svg+xml`) |
| Otro método en `/` (p. ej. `POST`) | `405 Method Not Allowed` |
| Ruta inexistente | `404 Not Found` |

---

## Decisiones técnicas

| Decisión | Motivo |
|---|---|
| Solo biblioteca estándar (`net/http`) | Sin dependencias externas: build más simple y reproducible |
| Patrón `GET /{$}` (Go 1.22+) | Solo acepta GET y solo la raíz exacta |
| Escuchar en `:8080` (todas las interfaces) | Imprescindible para que Docker pueda redirigir el tráfico al contenedor |
| `html/template` | Escapa automáticamente los datos y previene ataques XSS |
| La plantilla se procesa una vez al arrancar (`template.Must`) | Si tiene un error, el programa se detiene al arrancar, no ante cada visitante |
| `//go:embed templates static` | Plantilla, CSS e imagen van dentro del binario. Prepara una imagen Docker mínima en la capa 4 |
| Imagen en SVG | Ocupa 1,5 KB y se ve nítida a cualquier tamaño |
| Fuentes del sistema, sin Google Fonts | La página no depende de Internet para verse bien |

---

## Diseño: paleta Google Drive

Cada sección es una «carpeta» de color. Su color se guarda en una variable CSS (`--color`) que heredan todos sus elementos.

| Color | Hex | Uso |
|---|---|---|
| Azul | `#0066DA` | Experiencia, botón de contacto |
| Verde | `#00AC47` | Formación, lema |
| Amarillo | `#FFBA00` | Habilidades (solo fondo y decoración) |
| Rojo | `#EA4335` | Contacto (solo fondo y decoración) |

Para el **texto** se usan versiones oscuras de cada color (`#0057B8`, `#00732A`, `#8A5A00`, `#C5221F`), que cumplen el contraste mínimo recomendado (AA) sobre blanco. Solo se usa la paleta de colores, no el logotipo de Google Drive.

Técnicas CSS empleadas:
- Variables CSS
- `conic-gradient` (anillo del avatar)
- `clip-path` (icono de carpeta)
- `color-mix` (fondos suaves)
- Pseudo-elementos `::before` / `::after` (decoración)
- *Media queries* (adaptación al móvil)

> La web tiene un estilo propio a propósito. Los documentos de la asignatura siguen usando la plantilla de estilos UOC.

---

## Hoja de ruta

| Capa | Contenido | Estado |
|---|---|---|
| 1 | Servidor mínimo en Go + Dockerfile sencillo | ✅ Hecha |
| 2 | Perfil, plantilla HTML, CSS e imagen | ✅ **Esta versión** |
| 3 | Varios tipos de petición: API JSON, `/healthz`, formulario `POST` | ⏳ Pendiente |
| 4 | Dockerfile multietapa, `.dockerignore`, usuario sin privilegios | ⏳ Pendiente |
| 5 | Documentación y PDF de entrega | ⏳ Pendiente |

**Limitación conocida:** todavía no hay `.dockerignore`, así que `COPY . .` también copia este README dentro de la imagen. No afecta al funcionamiento y se resolverá en la capa 4.

---

## Uso de IA

Proyecto desarrollado con el apoyo de un asistente de IA (Claude, de Anthropic) como compañero de estudio. El alumno ha ejecutado, probado y documentado cada paso en su propio equipo.

## Referencias

- Amsterdam, J. (2024). *Routing enhancements for Go 1.22*. The Go Blog. <https://go.dev/blog/routing-enhancements>
- The Go Authors. *Package net/http*. <https://pkg.go.dev/net/http>
- The Go Authors. *Package embed*. <https://pkg.go.dev/embed>
- The Go Authors. *Package html/template*. <https://pkg.go.dev/html/template>
- Docker Inc. *Dockerfile reference*. <https://docs.docker.com/reference/dockerfile/>
- Docker Inc. *Docker Desktop WSL 2 backend on Windows*. <https://docs.docker.com/desktop/features/wsl/>
- Microsoft. *Working across Windows and Linux file systems*. <https://learn.microsoft.com/en-us/windows/wsl/filesystems>
- W3C. *Web Content Accessibility Guidelines (WCAG) 2.2*. <https://www.w3.org/TR/WCAG22/>
