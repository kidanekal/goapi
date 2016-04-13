package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/kidanekal/goapi/api"
	"github.com/kidanekal/goapi/logger"
)

var packageLogger = logger.CLI()
var port = 4200

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := api.NewRouter()

	packageLogger.Info("goapi listenting", "port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), router); err != nil {
		packageLogger.Crit("Faild to start service", "Error", err)
	}

}
