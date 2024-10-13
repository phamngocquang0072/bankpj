package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/bankpj/initializers"
	"github.com/phamngocquang0072/bankpj/controllers"
)

func init() {
	initializers.GetEnvVariable()
	initializers.ConnectDB()
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	UserRouter(r)
	r.Run(":"+port)
}

func UserRouter(r *gin.Engine) {
	r.GET("/user", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.PostUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
}