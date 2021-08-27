package main

import (
	"goland-demo/demo"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// SET ENV
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	demo.Demo()
	//var (
	//	mainInit sync.Once
	//	ctx      context.Context
	//	cancel   context.CancelFunc
	//	s        *server.Server
	//)
	//
	//mainInit.Do(func() {
	//	// SET CONTEXT
	//	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//
	//	// SET SERVER
	//	if os.Getenv("APPS_ENV") == "prod" || os.Getenv("APPS_ENV") == "qa" {
	//		gin.SetMode(gin.ReleaseMode)
	//	}
	//
	//	var host = os.Getenv("APPS_HOST")
	//	var port = os.Getenv("APPS_PORT")
	//	s = server.NewServer(host, port)
	//})
	//
	//if err := s.ServeHTTP(); err != nil {
	//	log.Fatal("Failed starting server:", err)
	//}
	//
	//// GRACEFUL SHUTDOWN
	//{
	//	quit := make(chan os.Signal)
	//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//	<-quit
	//	log.Println("Shutting down server...")
	//
	//	defer cancel()
	//	if err := s.Shutdown(ctx); err != nil {
	//		log.Fatal("Server forced to shutdown: ", err)
	//	}
	//	log.Println("Server exiting")
	//}
}
