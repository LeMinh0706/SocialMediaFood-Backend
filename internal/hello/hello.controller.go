package hello

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	helloService *HelloService
}

func NewHelloController() *HelloController {
	return &HelloController{
		helloService: NewHelloService(),
	}
}

// HelloExample godoc
// @Summary hello example
// @Schemes
// @Description check param
// @Tags example
// @Accept json
// @Produce json
// @Param name path string true "Name"
// @Router /example/hi/{name} [get]
func (hc *HelloController) GetHelloParam(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		response.ErrorNonKnow(g, 400, "Input your name on param")
		return
	}
	data := hc.helloService.GetName(name)
	response.SuccessResponse(g, 201, data)
}

// HelloExample godoc
// @Summary hello example
// @Schemes
// @Description check query
// @Tags example
// @Accept json
// @Produce json
// @Param name query string true "Name in query"
// @Router /example/hello [get]
func (hc *HelloController) GetHelloQuery(g *gin.Context) {
	name := g.Query("name")

	if name == "" {
		response.ErrorNonKnow(g, 400, "Name is required")
		return
	}
	// id := g.Query("id")
	// g.JSON(http.StatusOK, gin.H{
	// 	"message": "Hello " + hc.helloService.GetName(name),
	// 	"status":  "200",
	// })
	data := hc.helloService.GetName(name)
	response.SuccessResponse(g, 200, data)
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
// @Param body body Hello true "ID and Name here"
// @Router /example/hello [post]
func (hc *HelloController) PostHelloBody(g *gin.Context) {
	var body Hello
	if err := g.ShouldBindJSON(&body); err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}

	hello := hc.helloService.GetInfo(body.Name, body.Id)

	response.SuccessResponse(g, 200, hello)
	// g.JSON(http.StatusOK, gin.H{
	// 	"Message": "Hello " + hello.Name,
	// 	"Id":      hello.Id,
	// 	"status":  200,
	// })
}
