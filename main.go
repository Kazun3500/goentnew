package main

import (
	"context"
	"fmt"
	"goentnew/ent"
	"log"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"

	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}

	createUser := func(c *gin.Context, fruit *ent.User) (*ent.User, error) {
		u, err := client.User.Create().SetAge(fruit.Age).SetName(fruit.Name).Save(ctx)
		if err != nil {
			return nil, err
		}
		return u, nil

	}
	getUsers := func(c *gin.Context) ([]*ent.User, error) {
		u, err := client.User.Query().All(ctx)
		if err != nil {
			return nil, err
		}
		return u, nil

	}

	routes := func(grp *fizz.RouterGroup) {
		// Add a new fruit to the market.

		grp.POST("", []fizz.OperationOption{
			fizz.Summary("Add a fruit to the market"),
			fizz.Response("400", "Bad request", nil, nil,
				map[string]interface{}{"error": "fruit already exists"},
			),
		}, tonic.Handler(createUser, 200))

		// List all available fruits.
		grp.GET("", []fizz.OperationOption{
			fizz.Summary("List the fruits of the market"),
			fizz.Response("400", "Bad request", nil, nil, nil),
			fizz.Header("X-Market-Listing-Size", "Listing size", fizz.Long),
		}, tonic.Handler(getUsers, 200))
	}
	// NewRouter returns a new router for the
	// Pet Store.
	newRouter := func() (*fizz.Fizz, error) {
		engine := gin.New()
		engine.Use(cors.Default())
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
			c.URL = "../openapi.json"
		}, ginSwagger.DefaultModelsExpandDepth(1)))

		fizz := fizz.NewFromEngine(engine)

		// Override type names.
		// fizz.Generator().OverrideTypeName(reflect.TypeOf(Fruit{}), "SweetFruit")

		// Initialize the informations of
		// the API that will be served with
		// the specification.
		infos := &openapi.Info{
			Title:       "Fruits Market",
			Description: `This is a sample Fruits market server.`,
			Version:     "1.0.0",
		}
		// Create a new route that serve the OpenAPI spec.
		fizz.GET("/openapi.json", nil, fizz.OpenAPI(infos, "json"))

		// Setup routes.
		routes(fizz.Group("/market", "market", "Your daily dose of freshness"))

		if len(fizz.Errors()) != 0 {
			return nil, fmt.Errorf("fizz errors: %v", fizz.Errors())
		}
		return fizz, nil
	}

	router, err := newRouter()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    ":4242",
		Handler: router,
	}
	srv.ListenAndServe()
}
