package main

import (
	"context"
	"edu_v2/graph"
	database "edu_v2/internal/config"
	"edu_v2/internal/repository"
	"edu_v2/internal/service"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	database.ConnectPostgres()
	database.ConnectRedis()
	db := database.DB

	//repos
	groupRepo := repository.NewGroupRepository(db)
	collRepo := repository.NewCollectionRepository(db)

	//services
	groupService := service.NewGroupService(groupRepo)
	collService := service.NewCollectionService(collRepo)

	//graphServer
	graphQLServer := startGraphQLServer(port, groupService, collService)

	//shut down
	waitForShutDown(graphQLServer)
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env files")
	}
}

func startGraphQLServer(port string, groupService *service.GroupService, collService *service.CollectionService) *http.Server {
	gqlMux := http.NewServeMux()

	gqlMux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	gqlMux.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			GroupService: groupService,
			CollService:  collService,
		},
	})))

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"}),
	)

	gqlSrv := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler(gqlMux),
	}

	go func() {
		if err := gqlSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting GraphQL server: %v", err)
		}
	}()
	log.Println("Server starting ...")
	return gqlSrv
}

func waitForShutDown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Servers exiting")
}
