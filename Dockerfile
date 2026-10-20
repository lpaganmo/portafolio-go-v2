# Dockerfile — versión 1 (sencilla, de una sola etapa).
# Es la "receta" con la que Docker construye la IMAGEN de nuestra web.
# Cada instrucción crea una CAPA; Docker reutiliza (cachea) las capas que no cambian.

# 1) Imagen base: un Linux Alpine que ya trae Go instalado.
#    Su versión (1.24) debe ser >= la línea "go 1.24" de go.mod.
FROM golang:1.24-alpine

# 2) Carpeta de trabajo DENTRO de la imagen (se crea si no existe).
WORKDIR /app

# 3) Copiamos primero solo go.mod: si mañana cambias main.go pero no go.mod,
#    Docker reutiliza esta capa en vez de rehacerla (builds más rápidos).
COPY go.mod ./

# 4) Ahora copiamos el resto del código del proyecto.
COPY . .

# 5) Compilamos: genera el ejecutable /app/portafolio dentro de la imagen.
RUN go build -o portafolio .

# 6) Documenta que la app escucha en el 8080 (NO abre el puerto por sí solo:
#    eso lo hace "-p" en docker run).
EXPOSE 8080

# 7) Qué se ejecuta al arrancar el contenedor. Forma JSON ("exec"): el binario
#    es el proceso principal y recibe directamente la señal de "docker stop".
CMD ["./portafolio"]
