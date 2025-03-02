package main

import (
	"flag"
	"mygrpcp_project/cmd"
	"mygrpcp_project/config"
	"mygrpcp_project/gRPC/server"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "config file")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)
	if err := server.NewGRPCServer(cfg); err != nil {
		panic(err)
	} else {
		time.Sleep(1e9)
		cmd.NewApp(cfg)
	}

}
