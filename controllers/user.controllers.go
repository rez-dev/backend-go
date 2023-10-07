package controllers

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backend-go/models"
)

func conexionDBTest() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "mysql"
	Nombre := "bd_test1"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/")
	if err != nil {
		panic(err.Error())
	}

	// crear base de datos
	_, err = conexion.Exec("CREATE DATABASE IF NOT EXISTS " + Nombre)

	// usar base de datos
	_, err = conexion.Exec("USE " + Nombre)

	// crear tabla
	_, err = conexion.Exec(`CREATE TABLE IF NOT EXISTS usuario (
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		email VARCHAR(50),
		password VARCHAR(50),
		rut VARCHAR(50),
		role VARCHAR(50)
	)`)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func GetUsers(c *gin.Context) {
	conexionEstablecida := conexionDBTest()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM usuario")
	if err != nil {
		// panic(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	user := models.User{}
	users := models.Users{}
	for obtenerRegistros.Next() {
		var id int64
		var name, email, password, rut, role string
		err = obtenerRegistros.Scan(&id, &name, &email, &password, &rut, &role)
		if err != nil {
			// panic(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Rut = rut
		user.Role = role
		users = append(users, user)
	}
	// fmt.Println(users)
	c.IndentedJSON(http.StatusOK, users)
}
