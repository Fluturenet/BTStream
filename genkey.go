package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/fluturenet/ed25519"
	"os"
	"path/filepath"
)

func generateKey() {
	kpair, _ := ed25519.GenerateKey(nil)
	target := sha1.Sum(kpair.PublicKey())
	filename := filepath.Join(dataDir, hex.EncodeToString(target[:])+".key")
	fmt.Println("writing new key to : ",filename)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(kpair.PrivateKey())
}
