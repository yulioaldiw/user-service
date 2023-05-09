package bootstrap

import (
	"context"
	"fmt"
	"log"

	database "user-service/internal/adapters/driven"
	"user-service/internal/core/application"
	"user-service/internal/platform/server"
	mysqldb "user-service/internal/platform/storage/mysql"
	"user-service/pkg/config"
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	fmt.Println("Web Server ready!")

	ctx := context.Background()
	db, err := config.ConfigDb(ctx)

	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	userRepository := mysqldb.NewUserRepository(db, config.Cfg.DbTimeout)
	userAdapter := database.NewUserAdapter(userRepository)
	userService := application.NewUserService(userAdapter)

	ctx, srv := server.NewServer(context.Background(), config.Cfg.DbHost, config.Cfg.Port, config.Cfg.DbTimeout, server.AppService{
		UserService: userService,
	})

	return srv.Run(ctx)
}
