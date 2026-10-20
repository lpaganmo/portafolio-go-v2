// Capa 2 — El servidor ahora muestra una página HTML con el perfil, un CSS y una imagen.
//
// Novedades respecto a la capa 1:
//  1. go:embed     -> la plantilla, el CSS y la imagen viajan DENTRO del binario.
//  2. html/template -> genera el HTML a partir de los datos de perfil.go.
//  3. Una segunda ruta, /static/, que entrega los ficheros estáticos (CSS e imagen).
package main

import (
	"embed"         // incrusta ficheros en el binario al compilar
	"html/template" // plantillas HTML con escapado automático (seguridad)
	"io/fs"         // para "entrar" en la subcarpeta static/
	"log"
	"net/http"
)

// La directiva //go:embed (sin espacio tras //) le dice al compilador:
// "mete las carpetas templates y static dentro del ejecutable".
// Ventaja: el binario es autosuficiente. En la capa 4 podremos usar una imagen
// Docker casi vacía, porque no habrá que copiar ninguna carpeta aparte.
//
//go:embed templates static
var archivos embed.FS

// La plantilla se lee y se comprueba UNA sola vez, al arrancar.
// template.Must hace que el programa se detenga si la plantilla tiene un error:
// mejor fallar al arrancar que devolver errores a cada visitante.
var plantilla = template.Must(template.ParseFS(archivos, "templates/index.html"))

// inicio rellena la plantilla con los datos del perfil y la envía al navegador.
func inicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := plantilla.Execute(w, perfilFicticio()); err != nil {
		log.Printf("error al generar la página: %v", err)
		http.Error(w, "Error interno", http.StatusInternalServerError)
	}
}

func main() {
	// fs.Sub crea una vista que empieza DENTRO de static/.
	// Así la URL /static/css/estilo.css apunta al fichero static/css/estilo.css.
	estaticos, err := fs.Sub(archivos, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", inicio)

	// "GET /static/" (con barra final) casa con todo lo que empiece por /static/.
	// StripPrefix quita "/static/" de la URL antes de buscar el fichero.
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(estaticos)))

	direccion := ":8080"
	// Mensaje corregido: escuchamos en TODAS las interfaces, no solo en localhost.
	log.Printf("Servidor escuchando en el puerto 8080 (todas las interfaces) → http://localhost:8080")
	log.Fatal(http.ListenAndServe(direccion, mux))
}
