package cmd

import (
	"context"
	"ecommerce-platform/Dao"
	"ecommerce-platform/Validation"
	"ecommerce-platform/app"
	config "ecommerce-platform/configs"
	AdminControllerImpl "ecommerce-platform/controllers/admin_controller/admin_controller_impl"
	AuthServiceImpl "ecommerce-platform/controllers/auth_service/auth_service_impl"
	SellerControllerImpl "ecommerce-platform/controllers/seller_controller/seller_controller_impl"
	UserControllerImpl "ecommerce-platform/controllers/user_controller/user_controller_impl"
	"ecommerce-platform/db_migrations"
	"log"

	logger "ecommerce-platform/logger/log_service_impl"
	"ecommerce-platform/middleware"
	"ecommerce-platform/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var ApiServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Api Starts Server",
	Run:   ExecuteApi,
}

func ExecuteApi(cmd *cobra.Command, args []string) {

	logger := logger.New()

	logger.Info("Executing Api Command !")

	if err := config.LoadConfig(); err != nil {
		logger.Fatal("Failed to load config: ", err)
	}

	postgresDB, err := app.ConnectToPostgres()
	if err != nil {
		logger.Fatal(err)
	} else {
		logger.Info("Connected to postgres!")
	}
	defer func() {
		logger.Info("Closing database connection")
		postgresDB.Close()
	}()

	migration := db_migrations.NewMigration()
	if err := migration.RunMigrations(postgresDB); err != nil {
		log.Fatalf("Migrations Failed %v", err)
	}

	ValidationService := Validation.NewValidationService()

	AuthDao, AdminDao, UserDao, SellerDao := Dao.NewAuthDao(postgresDB), Dao.NewAdminDao(postgresDB), Dao.NewUserDao(postgresDB), Dao.NewSellerDao(postgresDB)

	middleware.Db = SellerDao

	logger.Info("Starting Api Server")

	AuthService := AuthServiceImpl.NewAuthService(AuthServiceImpl.NewAuthServiceImpl{
		Logger:    logger,
		AuthDao:   AuthDao,
		DB:        postgresDB,
		SellerDao: SellerDao,
	})
	AdminController := AdminControllerImpl.NewAdminController(AdminControllerImpl.NewAdminControllerImpl{
		Logger:   logger,
		AuthDao:  AuthDao,
		AdminDao: AdminDao,
	})
	Usercontroller := UserControllerImpl.NewUserImpl(UserControllerImpl.UserControllerConfig{
		UserDao: UserDao,
		DB:      postgresDB,
	})

	SellerController := SellerControllerImpl.NewSellerImpl(SellerControllerImpl.SellerController{
		SellerDao: SellerDao,
		DB:        postgresDB,
	})

	router := router.NewRouter(logger, AuthService, ValidationService, AdminController, Usercontroller, SellerController)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)

	Server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg.Port),
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
