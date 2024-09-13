package controllers

import (
	"net/http"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/services"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	helloService *services.HelloService
}

func NewHelloController() *HelloController {
	return &HelloController{
		helloService: services.NewHelloService(),
	}
}

// HelloExample godoc
// @Summary hello example
// @Schemes
// @Description just hello
// @Tags example
// @Accept json
// @Produce json
// @Param name path string true "Name"
// @Router /example/hello/{name} [get]
func (hc *HelloController) GetHelloParam(g *gin.Context) {
	name := g.Param("name")
	// id := g.Query("id")
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello " + hc.helloService.GetName(name),
		// "id":      id,
		"status": "200",
	})
}

// HelloExample godoc
// @Summary hello example
// @Schemes
// @Description just hello
// @Tags example
// @Accept json
// @Produce json
// @Param name query string true "Name in query"
// @Router /example/hello [get]
func (hc *HelloController) GetHelloQuery(g *gin.Context) {
	name := g.Query("name")

	if name == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "name is required",
			"code":    400,
		})
		return
	}
	// id := g.Query("id")
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello " + hc.helloService.GetName(name),
		"status":  "200",
	})
}

type BodyRequest struct {
	Name string `json:"name" binding:"required" example:"Hiro"`
	Id   int    `json:"id" binding:"required"`
}

// HelloExample godoc
// @Summary hello example
// @Schemes
// @Description Check Hello struct on repo
// @Tags example
// @Accept json
// @Produce json
// @Param body body repo.Hello true "ID and Name here"
// @Router /example/hello [post]
func (hc *HelloController) PostHelloBody(g *gin.Context) {
	var body repo.Hello
	if err := g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": 400,
		})
		return
	}

	hello := hc.helloService.GetInfo(body.Name, body.Id)

	g.JSON(http.StatusOK, gin.H{
		"Message": "Hello " + hello.Name,
		"Id":      hello.Id,
		"status":  200,
	})
}
