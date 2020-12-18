package routers

import (
	"PROYECTintegrador/ProyectoGOI/app"
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDBmysql()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}
	// Migrate the schema
	conn.AutoMigrate(
		&models.Person{},
		&models.User{},
		&models.Rol{},
		&models.Sesiones{},
		&models.Tareas{},
		&models.Nivel{},
		&models.Grado{},
		&models.Alumno{},
		&models.Docente{},
	)

	r := gin.Default()

	//config := cors.DefaultConfig() https://github.com/rs/cors
	//config.AllowOrigins = []string{"http://localhost", "http://localhost:8086"}

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", app.ItemsIndex)

		v1.GET("/personas", app.PersonsIndex)
		v1.POST("/personas", authMiddleWare(), app.PersonsCreate)
		v1.GET("/personas/:id", app.PersonsGet)
		v1.PUT("/personas/:id", app.PersonsUpdate)
		v1.DELETE("/personas/:id", app.PersonsDelete)

		v1.GET("/users", app.UsersIndex)
		v1.POST("/users", app.UsersCreate)
		v1.GET("/users/:id", app.UsersGet)
		v1.PUT("/users/:id", app.UsersUpdate)
		v1.DELETE("/users/:id", app.UsersDelete)
		v1.POST("/login", app.UsersLogin)
		v1.POST("/logout", app.UsersLogout)

		v1.GET("/rol", app.RolLista)
		v1.POST("/rol", authMiddleWare(), app.RolCreate)
		v1.GET("/rol/:id", app.RolGetID)
		v1.PUT("/rol/:id", app.RolUpdate)
		v1.DELETE("/rol/:id", app.RolDelete)

		v1.GET("/sesiones", app.SesionIndex)
		v1.POST("/sesiones", authMiddleWare(), app.SesionCreate)
		v1.GET("/sesiones/:id", app.SesionGet)
		v1.PUT("/sesiones/:id", app.SesionUpdate)
		v1.DELETE("/sesiones/:id", app.SesionDelete)

		v1.GET("/tareas", app.TareaIndex)
		v1.POST("/tareas", authMiddleWare(), app.TareaCreate)
		v1.GET("/tareas/:id", app.TareaGet)
		v1.PUT("/tareas/:id", app.TareaUpdate)
		v1.DELETE("/tareas/:id", app.TareaDelete)

		v1.GET("/nivel", app.NivelIndex)
		v1.POST("/nivel", authMiddleWare(), app.NivelCreate)
		v1.GET("/nivel/:id", app.NivelGet)
		v1.PUT("/nivel/:id", app.NivelUpdate)
		v1.DELETE("/nivel/:id", app.NivelDelete)

		v1.GET("/grado", app.GradoIndex)
		v1.POST("/grado", authMiddleWare(), app.GradoCreate)
		v1.GET("/grado/:id", app.GradoGet)
		v1.PUT("/grado/:id", app.GradoUpdate)
		v1.DELETE("/grado/:id", app.GradoDelete)

		v1.GET("/alumno", app.AlumnoIndex)
		v1.GET("/alumno/:id", app.AlumnoGETID)
		v1.POST("/alumno", authMiddleWare(), app.AlumnoCreate)
		v1.PUT("/alumno/:id", app.AlumnoUpdate)
		v1.DELETE("/alumno/:id", app.AlumnoDelete)

		v1.GET("/docente", app.DocenteIndex)
		v1.GET("/docente/:id", app.DocenteGETID)
		v1.POST("/docente", authMiddleWare(), app.DocenteCreate)
		v1.PUT("/docente/:id", app.DocenteUpdate)
		v1.DELETE("/docente/:id", app.DocenteDelete)
	}

	return r
}

func connectDBmysql() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/wagner?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "http://localhost, http://localhost:8086,")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE ")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//https://dev.to/stevensunflash/a-working-solution-to-jwt-creation-and-invalidation-in-golang-4oe4

//https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
