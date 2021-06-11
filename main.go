package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/openware/rango/pkg/auth"
)

const (
	defaultDir = "config"
)

var (
	uid   = flag.String("uid", "U487205863", "UID")
	email = flag.String("email", "admin@barong.io", "Email")
	role  = flag.String("role", "admin", "Role")
	level = flag.Int("level", 3, "Level")
)

func checkDefaultDir() {
	i, err := os.Stat(defaultDir)
	if os.IsNotExist(err) {
		os.Mkdir(defaultDir, 0750)
	} else if err != nil {
		panic(err.Error())
	}
	if i != nil && !i.IsDir() {
		panic("JWT_PRIVATE_KEY_PATH unset, " + defaultDir + " exists and is not a directory!")
	}
}

func main() {
	flag.Parse()

	keyPath := os.Getenv("JWT_PRIVATE_KEY_PATH")
	if keyPath == "" {
		checkDefaultDir()
		keyPath = defaultDir + "/rsa-key"
	}

	ks, err := auth.LoadOrGenerateKeys(keyPath, keyPath+".pub")
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	t, err := auth.ForgeToken(*uid, *email, *role, *level, ks.PrivateKey, nil)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println(t)
}
