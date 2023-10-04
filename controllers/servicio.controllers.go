package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rez-dev/backend-go/app-go/models"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "mysql"
	Nombre := "test_usachatiendebd"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func GetTerms(c *gin.Context) {
	conexionEstablecida := conexionDB()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM wp_terms")
	if err != nil {
		// panic(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	term := Term{}
	terms := Terms{}
	for obtenerRegistros.Next() {
		var id, term_group int
		var name, slug string
		err = obtenerRegistros.Scan(&id, &name, &slug, &term_group)
		if err != nil {
			// panic(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		term.ID = id
		term.Name = name
		term.Slug = slug
		term.Term_group = term_group
		terms = append(terms, term)
	}
	// fmt.Println(terms)
	c.IndentedJSON(http.StatusOK, terms)
}
