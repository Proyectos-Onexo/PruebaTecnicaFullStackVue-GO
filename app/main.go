package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	usuario := "root"
	Contrasenia := ""
	Nombre := "restaurante"

	conexion, err := sql.Open(Driver, usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}

	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/detalles", detalles)
	http.HandleFunc("/filtros", filtros)

	log.Println("Esta cosa esta jalando")
	http.ListenAndServe(":3060", nil)
}

type Platillos struct {
	IDPlatillo int
	Nombre     string
	Precio     float32
	Imagen     string
	Detalles   string
	Tipo       int
}

func inicio(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hola mundo")
	// plantillas.ExecuteTemplate(w, "inicio", nil)

	con := conexionBD()
	registro, err := con.Query("SELECT * FROM platillos")
	if err != nil {
		panic(err)
	}

	platillo := Platillos{}
	arregloplatillo := []Platillos{}

	for registro.Next() {
		var tipo, idplatillo int
		var precio float32
		var nombre, imagen, detalles string

		err = registro.Scan(&idplatillo, &nombre, &precio, &imagen, &detalles, &tipo)

		if err != nil {
			panic(err)
		}
		platillo.IDPlatillo = idplatillo
		platillo.Nombre = nombre
		platillo.Precio = float32(precio)
		platillo.Imagen = imagen
		platillo.Detalles = detalles
		platillo.Tipo = tipo

		arregloplatillo = append(arregloplatillo, platillo)

	}

	//fmt.Println(arregloplatillo)
	plantillas.ExecuteTemplate(w, "inicio", arregloplatillo)

}
func detalles(w http.ResponseWriter, r *http.Request) {
	idplatillo := r.URL.Query().Get("id")
	con := conexionBD()
	registro, err := con.Query("SELECT * FROM platillos WHERE idplatillo=?", idplatillo)

	if err != nil {
		panic(err)
	}
	platillo := Platillos{}
	arregloplatillo := []Platillos{}

	for registro.Next() {
		var tipo, idplatillo int
		var precio float32
		var nombre, imagen, detalles string

		err = registro.Scan(&idplatillo, &nombre, &precio, &imagen, &detalles, &tipo)

		if err != nil {
			panic(err)
		}
		platillo.IDPlatillo = idplatillo
		platillo.Nombre = nombre
		platillo.Precio = float32(precio)
		platillo.Imagen = imagen
		platillo.Detalles = detalles
		platillo.Tipo = tipo

		arregloplatillo = append(arregloplatillo, platillo)

	}
	plantillas.ExecuteTemplate(w, "detalles", arregloplatillo)

}

func filtros(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tipo := r.FormValue("tipo")
		con := conexionBD()
		registro, err := con.Query("SELECT * FROM platillos WHERE tipo=?", tipo)
		if err != nil {
			panic(err)
		}
		platillo := Platillos{}
		arregloplatillo := []Platillos{}

		for registro.Next() {
			var tipo, idplatillo int
			var precio float32
			var nombre, imagen, detalles string

			err = registro.Scan(&idplatillo, &nombre, &precio, &imagen, &detalles, &tipo)

			if err != nil {
				panic(err)
			}
			platillo.IDPlatillo = idplatillo
			platillo.Nombre = nombre
			platillo.Precio = float32(precio)
			platillo.Imagen = imagen
			platillo.Detalles = detalles
			platillo.Tipo = tipo

			arregloplatillo = append(arregloplatillo, platillo)

		}
		plantillas.ExecuteTemplate(w, "filtros", arregloplatillo)

	}
}
