package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/loopfz/gadgeto/tonic"
	"goentnew/ent"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"

	_ "github.com/lib/pq"
)

func getDb() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, err
}

type DeleteResult struct {
	Count int
	Error string
}

type ApiServer struct {
	Db *ent.Client
}

func (s *ApiServer) CreateUser(c *gin.Context) {
	var u ent.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	us, err := s.Db.User.Create().SetAge(u.Age).SetName(u.Name).Save(context.Background())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := json.Marshal(us)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, string(data))
	return
}

func (s *ApiServer) GetUsers(c *gin.Context) {
	u, err := s.Db.User.Query().All(context.Background())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := json.Marshal(u)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, string(data))
	return
}

func (s *ApiServer) CreateUserFizz(c *gin.Context, fruit *ent.User) (*ent.User, error) {
	u, err := s.Db.User.Create().SetAge(fruit.Age).SetName(fruit.Name).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *ApiServer) GetUsersFizz(c *gin.Context) ([]*ent.User, error) {
	u, err := s.Db.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *ApiServer) DeleteUsersFizz(c *gin.Context) (*DeleteResult, error) {
	u, err := s.Db.User.Delete().Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &DeleteResult{Count: u, Error: ""}, nil
}

func InitFizzRoutes(grp *fizz.RouterGroup, server *ApiServer) {
	grp.POST("", []fizz.OperationOption{
		fizz.Summary("Add a fruit to the market"),
		fizz.Response("400", "Bad request", nil, nil,
			map[string]interface{}{"error": "fruit already exists"},
		),
	}, tonic.Handler(server.CreateUserFizz, 200))

	// List all available fruits.
	grp.GET("", []fizz.OperationOption{
		fizz.Summary("List the fruits of the market"),
		fizz.Response("400", "Bad request", nil, nil, nil),
		fizz.Header("X-Market-Listing-Size", "Listing size", fizz.Long),
	}, tonic.Handler(server.GetUsersFizz, 200))
}

func InitRoutes(grp *gin.RouterGroup, server *ApiServer) {
	grp.POST("", server.CreateUser)
	grp.GET("", server.GetUsers)
}

func runServer() {
	engine := gin.New()
	engine.Use(cors.Default())

	db, err := getDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	server := ApiServer{Db: db}
	// Setup routes.
	gr := engine.RouterGroup.Group("/market")
	InitRoutes(gr, &server)

	srv := &http.Server{
		Addr:    ":4242",
		Handler: engine,
	}
	srv.ListenAndServe()

}

func runFizzServer() {
	engine := gin.New()
	engine.Use(cors.Default())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.URL = "../openapi.json"
	}, ginSwagger.DefaultModelsExpandDepth(1)))

	fizzEngine := fizz.NewFromEngine(engine)

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
	fizzEngine.GET("/openapi.json", nil, fizzEngine.OpenAPI(infos, "json"))

	db, err := getDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	server := ApiServer{Db: db}
	// Setup routes.
	InitFizzRoutes(fizzEngine.Group("/market", "market", "Your daily dose of freshness"), &server)
	if len(fizzEngine.Errors()) != 0 {
		panic(fmt.Sprintf("fizz errors: %v", fizzEngine.Errors()))
	}
	srv := &http.Server{
		Addr:    ":4242",
		Handler: fizzEngine,
	}
	srv.ListenAndServe()

}

func main() {

	//runFizzServer()
	//runServer()
	db, err := getDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println(db.User.Query().FirstX(context.Background()))

}
