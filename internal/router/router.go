package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/factory"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/swag/docs"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(g *gin.Engine, config util.Config, token token.Maker, pgx *pgxpool.Pool) {
	fac, err := factory.NewFactory(pgx)
	if err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.BasePath = "/api"
	a := g.Group("/api")
	{
		Static(g)
		NewUserRouter(g, a, fac.UserService, token, config)
		NewAccountRouter(g, a, fac.AccountService, token)
		NewPostRouter(g, a, fac.PostService, token)
		NewCommentRouter(g, a, fac.CommentService, token)
		NewReactRouter(g, a, fac.ReactService)
		NewFollowRouter(g, a, fac.FollowService)
	}

	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
