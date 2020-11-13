package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/openware/rango/pkg/auth"
)

var (
	uid        = flag.String("uid", "U487205863", "UID")
	email      = flag.String("email", "admin@barong.io", "Email")
	role       = flag.String("role", "admin", "Role")
	level      = flag.Int("level", 3, "Level")
	configPath = "config"
)

func checkConfigDirectory() {
	i, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		os.Mkdir(configPath, 0750)
	} else if err != nil {
		panic(err.Error())
	}
	if i != nil && !i.IsDir() {
		panic(configPath + " exists and is not a directory!")
	}
}

func main() {
	flag.Parse()
	checkConfigDirectory()

	ks, err := auth.LoadOrGenerateKeys(configPath+"/rsa-key", configPath+"/rsa-key.pub")
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	t, err := auth.ForgeToken(*uid, *email, *role, *level, ks.PrivateKey, nil)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println(t)
}
