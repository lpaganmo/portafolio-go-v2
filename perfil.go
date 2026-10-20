// perfil.go — Capa 2: los DATOS del portafolio.
//
// Por qué un archivo aparte: separamos QUÉ se muestra (datos) de CÓMO se
// muestra (plantilla HTML) y de CÓMO se sirve (main.go). Si mañana cambia
// un dato, solo se toca este archivo.
//
// Todos los datos son FICTICIOS (persona, empresas y correo).
package main

// Perfil agrupa todo lo que necesita la plantilla. Los campos empiezan por
// MAYÚSCULA porque en Go eso significa "exportado": la plantilla HTML solo
// puede leer campos exportados.
type Perfil struct {
	Nombre      string
	Puesto      string
	Lema        string // El mensaje que exige el enunciado.
	Ubicacion   string
	Email       string
	SobreMi     []string // Una entrada por párrafo.
	Experiencia []Empleo
	Formacion   []Estudio
	Habilidades []GrupoHabilidades
}

// Empleo es un puesto de la sección "Experiencia".
type Empleo struct {
	Cargo   string
	Empresa string
	Periodo string
	Logros  []string
}

// Estudio es una línea de la sección "Formación".
type Estudio struct {
	Titulo  string
	Centro  string
	Periodo string
}

// GrupoHabilidades agrupa habilidades por categoría (Sistemas, Redes...).
type GrupoHabilidades struct {
	Categoria string
	Items     []string
}

// perfilFicticio devuelve los datos de Álex Navarro Ruiz.
func perfilFicticio() Perfil {
	return Perfil{
		Nombre:    "Álex Navarro Ruiz",
		Puesto:    "Técnico de Soporte IT · Sistemas y Redes",
		Lema:      "Soy alumno de la UOC",
		Ubicacion: "Valencia, España",
		Email:     "alex.navarro@example.com", // example.com: dominio reservado para ejemplos (RFC 2606).
		SobreMi: []string{
			"Soy técnico de soporte IT de segundo nivel. Mi trabajo consiste en que los equipos, las cuentas y la red de una empresa funcionen sin que nadie tenga que pensar en ellos.",
			"Estudio el CFGS de Administración de Sistemas Informáticos en Red en la UOC. Este portafolio está escrito en Go y se ejecuta dentro de un contenedor Docker: forma parte de lo que estoy aprendiendo.",
		},
		Experiencia: []Empleo{
			{
				Cargo:   "Técnico de Soporte IT (N2)",
				Empresa: "Logivant Iberia S.L.",
				Periodo: "2023 – actualidad",
				Logros: []string{
					"Resolución de incidencias de segundo nivel para unos 300 usuarios en tres sedes.",
					"Automatización de altas y bajas de usuarios con PowerShell.",
					"Documentación de procedimientos en la base de conocimiento interna.",
				},
			},
			{
				Cargo:   "Técnico de microinformática",
				Empresa: "Kubo Sistemas S.L.",
				Periodo: "2021 – 2023",
				Logros: []string{
					"Montaje, despliegue y mantenimiento de equipos de puesto de trabajo.",
					"Soporte presencial y remoto a pequeñas empresas.",
				},
			},
		},
		Formacion: []Estudio{
			{Titulo: "CFGS Administración de Sistemas Informáticos en Red (ASIR)", Centro: "UOC", Periodo: "En curso"},
			{Titulo: "CFGM Sistemas Microinformáticos y Redes (SMR)", Centro: "IES ficticio de Valencia", Periodo: "2019 – 2021"},
		},
		Habilidades: []GrupoHabilidades{
			{Categoria: "Sistemas", Items: []string{"Windows Server", "Active Directory", "Linux (Debian, Ubuntu)", "Microsoft 365"}},
			{Categoria: "Redes", Items: []string{"TCP/IP", "DNS y DHCP", "VLAN", "VPN"}},
			{Categoria: "Cloud y contenedores", Items: []string{"Docker", "Go", "AWS (fundamentos)", "Git"}},
			{Categoria: "Soporte", Items: []string{"ITIL", "PowerShell", "Bash", "Gestión de incidencias"}},
		},
	}
}
