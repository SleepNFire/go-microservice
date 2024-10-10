package gin

import "github.com/gin-gonic/gin"

const (
	Public    = "Public"
	Internal  = "Internal"
	Technical = "Technical"

	Port80 = ":8080"
	Port81 = ":8081"
	Port82 = ":8082"
)

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

func NewGinService() GinService {
	gs := GinService{
		Public: ApiGin{
			Engine: gin.New(),
			Name:   Public,
			Port:   Port80,
		},
		Internal: ApiGin{
			Engine: gin.New(),
			Name:   Internal,
			Port:   Port81,
		},
		Technical: ApiGin{
			Engine: gin.New(),
			Name:   Technical,
			Port:   Port82,
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
