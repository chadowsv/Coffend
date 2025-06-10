package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

// Definicion de variable global para que sea accesible desde otros paquetes
var DB *sql.DB

func Connect() {
	// Cargar .env, godotenv utilizado para lectura de archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env:", err)
	}
	//Extraccion de variables de entorno
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	//URL de conexion
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, dbname)
	//sql.Opne se encarga de la preparacion de la conexion, si existe un error lo detiene
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error abriendo conexión:", err)
	}
	//Ping prueba que de vdd se pueda realizar la conexion a la base de datos
	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	log.Println("Conexión exitosa a SQL Server 🎉")
}
