package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"

	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app"
	logger "github.com/sirupsen/logrus"
)

// @title           Punqy API
// @version         1.0
// @description     Punqy rest api.
// @BasePath  		/api/v1
func main() {
	println("===============================")
	println("=      Punqy© application     =")
	println("===============================")
	if err := punqy.Execute(app.Commands()); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			logger.Error(err)
			return
		}
		os.Exit(1)
	}
}
