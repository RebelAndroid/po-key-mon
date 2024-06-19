package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/argon2"
)

func main() {
	fmt.Println("Hello world")

	b := argon2.IDKey([]byte(""), []byte("po-key-mon-salt"), 64, 128*1024, 8, 32)

	r := bytes.NewReader(b[:])
	pubkey, privkey, err := ed25519.GenerateKey(r)
	if err != nil {
		log.Fatal("unable to generate key: ", err)
	}
	fmt.Println(len(pubkey))
	fmt.Println(len(privkey))

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(privkey)
	encoder.Close()
}
