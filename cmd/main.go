package main

import (
	"github.com/anafro/cosmoport/internal/config"
	"github.com/anafro/cosmoport/internal/server"
)

func main() {
	conf, err := config.Load("./cosmoport.yaml")

	if err != nil {
		panic(err)
	}

	serv := server.NewServer(conf)
	_ = serv // TODO
}
