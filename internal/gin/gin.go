package gin

import "github.com/gin-gonic/gin"

type ApiGin struct {
	Engine *gin.Engine
	Name   string
	Port   string
}

type GinService struct {
	Public    ApiGin
	Internal  ApiGin
	Technical ApiGin
}

func newGinService() GinService {
	gs := GinService{
		Public: ApiGin{
			Engine: gin.New(),
			Name:   "Public",
			Port:   "8080",
		},
		Internal: ApiGin{
			Engine: gin.New(),
			Name:   "Internal",
			Port:   "8081",
		},
		Technical: ApiGin{
			Engine: gin.New(),
			Name:   "Technical",
			Port:   "8082",
		},
	}
	gs.Public.Engine.Use(gin.Logger())
	gs.Public.Engine.Use(gin.Recovery())

	gs.Internal.Engine.Use(gin.Logger())
	gs.Internal.Engine.Use(gin.Recovery())

	gs.Technical.Engine.Use(gin.Logger())
	gs.Technical.Engine.Use(gin.Recovery())

	return gs
}
