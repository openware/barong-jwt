package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	uid   = flag.String("uid", "U487205863", "UID")
	email = flag.String("email", "admin@barong.io", "Email")
	role  = flag.String("role", "admin", "Role")
	level = flag.Int("level", 3, "Level")
)

func main() {
	flag.Parse()
	ks, err := LoadOrGenerateKeys("config/rsa-key", "config/rsa-key.pub")
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	t, err := ForgeToken(*uid, *email, *role, *level, ks.PrivateKey, nil)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println(t)
}
