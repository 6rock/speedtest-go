package main

import (
	"flag"
	"github.com/librespeed/speedtest/results"
	_ "time/tzdata"

	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/database"
	"github.com/librespeed/speedtest/web"

	_ "github.com/breml/rootcerts"
	log "github.com/sirupsen/logrus"
)

//var (
//	optConfig = flag.String("c", "", "config file to be used, defaults to settings.toml in the same directory")
//)

func main() {
	//flag.Parse()
	//conf := config.Load(*optConfig)
	var (
		addr string
		port int
	)
	flag.StringVar(&addr, "addr", "", "IP address to listen on")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Parse()
	config.SetConfig(addr, port)
	web.SetServerLocation(config.LoadedConfig())
	results.Initialize(config.LoadedConfig())
	database.SetDBInfo(config.LoadedConfig())
	log.Fatal(web.ListenAndServe(config.LoadedConfig()))
}
