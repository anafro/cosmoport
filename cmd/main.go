package main

import (
	"fmt"

	"github.com/anafro/cosmoport/internal/config"
	"github.com/anafro/cosmoport/internal/server"
	"github.com/anafro/cosmoport/internal/support"
)

func main() {
	conf := support.Panics(config.Load("./cosmoport.yaml"))
	serv := server.NewServer(conf)

	go serv.Start()

	for {
		fmt.Printf(<-serv.Logs)
	}
}
