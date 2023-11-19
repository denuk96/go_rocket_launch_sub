package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"

	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"

	"github.com/zhashkevych/scheduler"
)

func main() {
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, LaunchNotification, time.Second*5)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	worker.Stop()
}

func LaunchNotification(_ctx context.Context) {
	logrus.Info("Starting LaunchNotification")

	config.InitEnvs()
	db := config.InitDB()
	if db == nil {
		logrus.Error("Received nil database connection")
		return
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	services.Notification.NotifyAllWithin(1)

	logrus.Info("LaunchNotification completed successfully")
}
