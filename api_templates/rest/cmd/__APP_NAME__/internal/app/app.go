package app

import (
	"fmt"
	"net/http"

	"wwdb-queue-func/config"
	"wwdb-queue-func/internal/queue"
	"wwdb-queue-func/internal/router"
	"wwdb-queue-func/pkg/ctxlogger"
	"wwdb-queue-func/pkg/mongodb"
)

func Init() {

	//config has to be initialized before logger for debugFlag
	config := config.Load()
	fmt.Println(config)

	ctxlogger.Init(config.App.Debug)

	//startupscreen.Splash(config)

	mongodb.Init(config)

	go queue.Subscriber()

	router := router.New(config)
	fmt.Println(config.Server.Port)
	http.ListenAndServe(":"+config.Server.Port, router)
}
