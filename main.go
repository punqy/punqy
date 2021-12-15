package main

import (
	"fmt"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app"
	logger "github.com/sirupsen/logrus"
	_ "net/http/pprof"
	"os"
)

func main() {
	println("===============================")
	println("=      LUCI© application      =")
	println("===============================")
	if err := punqy.Execute(app.Commands()); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			logger.Error(err)
			return
		}
		os.Exit(1)
	}
}
