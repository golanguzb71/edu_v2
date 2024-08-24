package main

import (
	"context"
	"edu_v2/graph"
	database "edu_v2/internal/config"
	"edu_v2/internal/repository"
	"edu_v2/internal/service"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	database.ConnectPostgres()
	database.ConnectRedis()
	db := database.DB
	rdb := database.RDB
	groupRepo := repository.NewGroupRepository(db)
	collRepo := repository.NewCollectionRepository(db)
	answerRepo := repository.NewAnswerRepository(db, rdb)
	userCollRepo := repository.NewUserCollectionRepository(db, rdb)

	answerService := service.NewAnswerService(answerRepo)
	groupService := service.NewGroupService(groupRepo)
	collService := service.NewCollectionService(collRepo)
	userService := service.NewUserCollectionUserService(userCollRepo)

	server := startServer(port, groupService, collService, answerService, userService)

	waitForShutdown(server)
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func startServer(port string, groupService *service.GroupService, collService *service.CollectionService, answerService *service.AnswerService, userCollService *service.UserCollectionService) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	mux.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			GroupService:    groupService,
			CollService:     collService,
			AnswerService:   answerService,
			UserCollService: userCollService,
		},
	})))

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/images/" {
			http.NotFound(w, r)
			return
		}
		imagePath := filepath.Join("question_images", r.URL.Path[len("/images/"):])

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, imagePath)
	})

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"}),
	)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler(mux),
	}

	go func() {
		log.Printf("Server is starting on port %s...", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	return server
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited")
}
