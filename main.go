package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"

	"kietchung/controllers"
	"kietchung/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server       *gin.Engine
	cs           services.ChemistryService
	cc           controllers.ChemistryController
	ctx          context.Context
	chemistryCol *mongo.Collection
	refDocCol    *mongo.Collection
	mongoclient  *mongo.Client
	err          error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://tuankiet10022171:kietlu1712@cluster0.znigccy.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)

	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	chemistryCol = mongoclient.Database("chemistry").Collection("chemistry")
	refDocCol = mongoclient.Database("chemistry").Collection("ref_document")

	cs = services.NewUserService(chemistryCol, refDocCol, ctx)
	cc = controllers.New(cs)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	cc.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":3000"))

}
