package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/argon2"
)

func get_key(input []byte) string {
	b := argon2.IDKey(input, []byte("po-key-mon-salt"), 32, 64*1024, 8, 32)

	r := bytes.NewReader(b[:])
	_, privkey, err := ed25519.GenerateKey(r)
	if err != nil {
		log.Fatal("unable to generate key: ", err)
	}

	var output_buffer bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &output_buffer)
	encoder.Write(privkey)
	encoder.Close()
	return output_buffer.String()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	body := new(strings.Builder)
	io.Copy(body, r.Body)
	fmt.Fprintln(w, get_key([]byte(body.String())))
}

func main() {
	fmt.Println(get_key([]byte("hello world")))

	http.HandleFunc("/get_key", testHandler)
	log.Fatal(http.ListenAndServe(":47001", nil))
}
