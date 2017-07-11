package main
// only need mysql OR sqlite
// both are included here for reference
import (
    "github.com/gin-gonic/gin"
    "github.com/korolr/gingorm/db"
    "github.com/korolr/gingorm/controllers"
    "github.com/korolr/gingorm/middlewares"
)

func main() {
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function
    //db, err = gorm.Open("sqlite3", "./gorm.db")
    db.Init()

    r := gin.Default()

    v1 := r.Group("/v1")
    v1.Use(middlewares.AuthHandler())
    {
        v1.GET("/people", controllers.GetPeople)
        v1.GET("/people/:id", controllers.GetPerson)
        v1.POST("/people", controllers.CreatePerson)
        v1.PUT("/people/:id", controllers.UpdatePerson)
        v1.DELETE("/people/:id", controllers.DeletePerson)
    }

    r.GET("/auth", controllers.Auth)

    r.Run(":8080")

    defer db.CloseDB()
}


