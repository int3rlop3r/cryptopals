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

// challenge: 3
func xorFreq(s string) ([]byte, error) {
	h1, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("xorFreq: %w", err)
	}
	var maxScore int
	var cypherKey []byte
	newBuff := make([]byte, len(h1), len(h1))
	for a := 65; a <= 122; a++ {
		// key len should be same as cypher len (bytes)
		key := bytes.Repeat([]byte{byte(a)}, len(h1))
		for i := range key {
			newBuff[i] = h1[i] ^ key[i]
		}
		score := freqCheck(newBuff)
		if score > maxScore {
			maxScore = score
			cypherKey = key
		}
	}
	return cypherKey, nil
}

func freqCheck(s []byte) int {
	commLetters := "etaoin shrdlu"
	score := 0
	var countL int // count of lower-case
	var countU int // count of upper-case
	for i := 0; i < len(commLetters); i++ {
		points := len(commLetters) - i
		countL = bytes.Count(s, []byte{commLetters[i]})
		if commLetters[i] != ' ' {
			countU = bytes.Count(s, []byte{commLetters[i] - byte(32)})
		} else {
			countU = 0
		}
		score += points * (countL + countU)
	}
	return score
}

func main() {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, err := xorFreq(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key:", string(key))
}
