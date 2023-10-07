package controllers

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backend-go/models"
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
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM wp_terms")
	if err != nil {
		// panic(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	term := models.Term{}
	terms := models.Terms{}
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

func GetTests(c *gin.Context) {
	conexionEstablecida := conexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT wp_terms.name AS 'Nombre del Término', wp_term_taxonomy.taxonomy AS 'Taxonomía', wp_posts.post_title AS 'Título de la Publicación', wp_posts.post_content AS 'Contenido de la Publicación', wp_postmeta.meta_key AS 'Clave del Meta', wp_postmeta.meta_value AS 'Valor del Meta' FROM wp_terms INNER JOIN wp_term_taxonomy ON wp_terms.term_id = wp_term_taxonomy.term_id INNER JOIN wp_term_relationships ON wp_term_taxonomy.term_taxonomy_id = wp_term_relationships.term_taxonomy_id INNER JOIN wp_posts ON wp_term_relationships.object_id = wp_posts.ID LEFT JOIN wp_postmeta ON wp_posts.ID = wp_postmeta.post_id WHERE wp_postmeta.meta_key IS NOT NULL")
	if err != nil {
		// panic(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	test := models.Test{}
	tests := models.Tests{}
	for obtenerRegistros.Next() {
		var name, taxonomy, post_title, post_content, meta_key, meta_value string
		err = obtenerRegistros.Scan(&name, &taxonomy, &post_title, &post_content, &meta_key, &meta_value)
		if err != nil {
			// panic(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		test.Name = name
		test.Taxonomy = taxonomy
		test.Post_title = post_title
		test.Post_content = post_content
		test.Meta_key = meta_key
		test.Meta_value = meta_value
		tests = append(tests, test)
	}

	c.IndentedJSON(http.StatusOK, tests)
}
