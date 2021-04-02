package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// challenge: 1
func hexTo64(s string) (string, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("hexto64: %w", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// challenge: 2
func xor(buff1, buff2 string) (string, error) {
	if len(buff1) != len(buff2) {
		return "", errors.New("xor: strings must be of equal lengths")
	}
	hex1, err := hex.DecodeString(buff1)
	if err != nil {
		return "", fmt.Errorf("xor: %w", err)
	}
	hex2, err := hex.DecodeString(buff2)
	if err != nil {
		return "", fmt.Errorf("xor: %w", err)
	}
	newBuff := make([]byte, len(hex1), len(hex2))
	for i := 0; i < len(hex1); i++ {
		newBuff[i] = hex1[i] ^ hex2[i]
	}
	return hex.EncodeToString(newBuff), nil
}

func nextKey(l int) <-chan []byte {
	kCh := make(chan []byte)
	go func() {
		defer close(kCh)
		for a := 97; a <= 122; a++ {
			rep := bytes.Repeat([]byte{byte(a)}, l)
			kCh <- rep
		}
	}()
	return kCh
}

func xorFreq(s string) {
	h1, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("xorFreq: %s\n", err)
		return
	}
	newBuff := make([]byte, len(h1), len(h1))
	for key := range nextKey(len(h1)) {
		//fmt.Println("key:", string(key))
		for i := range key {
			newBuff[i] = h1[i] ^ key[i]
		}
		fmt.Println("xor'ed:", string(newBuff))
	}
}

func main() {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	xorFreq(s)
}
