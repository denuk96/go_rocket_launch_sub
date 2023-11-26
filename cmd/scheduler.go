package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"

	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"

	"github.com/zhashkevych/scheduler"
)

func main() {
	log.Info("Starting scheduler...")

	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx,
		func(ctx context.Context) {
			LaunchNotification(ctx)
		},
		time.Second*5)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	worker.Stop()
}

func LaunchNotification(_ctx context.Context) {
	log.Info("Starting LaunchNotification")

	config.InitEnvs()
	db := config.InitDB()
	if db == nil {
		log.Error("Received nil database connection")
		return
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	services.Notification.NotifyAll()

	log.Info("LaunchNotification completed successfully")
}
