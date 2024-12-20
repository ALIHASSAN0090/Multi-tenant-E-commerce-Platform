package cmd

import (
	"context"
	"ecommerce-platform/Dao"
	"ecommerce-platform/Validation"
	"ecommerce-platform/app"
	config "ecommerce-platform/configs"
	AuthServiceImpl "ecommerce-platform/controllers/auth_service/auth_service_impl"
	logger "ecommerce-platform/logger/log_service_impl"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var ApiServerCommand = &cobra.Command{
	Use:   "Api",
	Short: "Api Starts Server",
	Run:   ExecuteApi,
}

func ExecuteApi(cmd *cobra.Command, args []string) {

	logger := logger.New()

	logger.Info("Executing Api Command !")

	config.InitConfig()

	fmt.Println(config.AppConfig.DB_DATABASE, "dbname")

	postgresDB, err := app.ConnectToPostgres()
	if err != nil {
		logger.Fatal(err)
	} else {
		logger.Info("Connected to postgres!")
	}

	ValidationService := Validation.NewValidationService()

	AuthDao := Dao.NewAuthDao(postgresDB)

	logger.Info("Starting Api Server")

	AuthService := AuthServiceImpl.NewAuthService(AuthServiceImpl.NewAuthServiceImpl{
		Logger:  logger,
		AuthDao: AuthDao,
	})

	router := router.NewRouter(logger, AuthService, ValidationService)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)

	Server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.AppConfig.APP_ADDRESS),
		Handler: router.Engine,
	}

	go func() {
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop

	logger.Info("Stopping API server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := Server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	logger.Info("API server stopped.")
}
