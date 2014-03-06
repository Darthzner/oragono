package main

import (
	"flag"
	"fmt"
	"github.com/jlatt/ergonomadic/irc"
	"log"
	"os"
	"path/filepath"
)

func main() {
	conf := flag.String("conf", "ergonomadic.json", "ergonomadic config file")
	initdb := flag.Bool("initdb", false, "initialize database")
	passwd := flag.String("genpasswd", "", "bcrypt a password")
	flag.Parse()

	if *passwd != "" {
		encoded, err := irc.GenerateEncodedPassword(*passwd)
		if err != nil {
			log.Fatal("encoding error: ", err)
		}
		fmt.Println(encoded)
		return
	}

	config, err := irc.LoadConfig(*conf)
	if err != nil {
		log.Fatal("error loading config: ", err)
	}
	err = os.Chdir(filepath.Dir(*conf))
	if err != nil {
		log.Fatal("chdir error: ", err)
	}

	if *initdb {
		irc.InitDB(config.Server.Database)
		log.Println("database initialized: " + config.Server.Database)
		return
	}

	// TODO move to data structures
	irc.DEBUG_NET = config.Debug.Net
	irc.DEBUG_CLIENT = config.Debug.Client
	irc.DEBUG_CHANNEL = config.Debug.Channel
	irc.DEBUG_SERVER = config.Debug.Server

	irc.NewServer(config).Run()
}
