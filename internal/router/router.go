package router

import (
	"github.com/dargasht/gocrud"
	"github.com/dargasht/gocrud/internal/cfg"
	"github.com/dargasht/gocrud/internal/database/repo"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	loggerMW "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func SetupApp(app *fiber.App, db *repo.Queries, logger *zap.Logger) *fiber.App {

	//Middlewares

	app.Use(healthcheck.New(healthcheck.Config{LivenessEndpoint: "/health"}))
	app.Use(favicon.New())
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New())
	// app.Use(limiter.New(limiter.Config{
	// 	Max:               100,
	// 	Expiration:        30 * time.Second,
	// 	LimiterMiddleware: limiter.SlidingWindow{},
	// }))
	app.Use(loggerMW.New(loggerMW.Config{
		TimeZone:   "Asia/Tehran",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	//FOR DEVELOPMENT
	app.Get("/api/v1/metrics", monitor.New())
	app.Use(pprof.New())

	//---------------------------------------------------------------
	//Main Routes

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	//---------------------------------------------------------------
	//Handler configs including database and logger

	handlerConfig := gocrud.NewHandlerConfig(db, logger)

	//---------------------------------------------------------------
	//JWT middleware
	jwt := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.AUTHSECRET)}, TokenLookup: "header:Authorization", AuthScheme: "Bearer",
	})

	//---------------------------------------------------------------
	//Routes

	config := routeConfig{
		r:             v1,
		handlerConfig: *handlerConfig,
		jwt:           jwt,
	}

	setProductRoutes(config)

	//TODO: USE IT LATER
	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// fmt.Println(string(data))

	return app
}

func setProductRoutes(c routeConfig) {

}
