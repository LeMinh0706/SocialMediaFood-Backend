package menu

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service IMenuService
	token   token.Maker
}

func NewMenuController(service IMenuService, token token.Maker) *MenuController {
	return &MenuController{
		service: service,
		token:   token,
	}
}

// Menu godoc
// @Summary      Add food
// @Description  Add food
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Param        request body AddFood true "request"
// @Success      201  {object}  MenuResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /menu [post]
func (m *MenuController) CreateNewFood(g *gin.Context) {
	var req AddFood
	if err := g.ShouldBindJSON(&req); err != nil {
		// response.ErrorResponse(g, 40000)
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	dish, err := m.service.CreateDish(g, req.AccountID, req.Quantity, req.DishName, req.Img, req.Price)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, dish)
}

// Menu godoc
// @Summary      Add food
// @Description  Add food
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Param        account_id query int false "AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      201  {object}  MenuResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /menu [get]
func (m *MenuController) GetMenu(g *gin.Context) {
	accStr := g.Query("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := post.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := m.service.GetMenu(g, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}
