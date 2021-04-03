package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// challenge: 1
func hexTo64(s string) ([]byte, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("hexto64: %w", err)
	}
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)
	return encoded, nil
}

// challenge: 2
func fixedXOR(buff1, buff2 string) ([]byte, error) {
	if len(buff1) != len(buff2) {
		return nil, errors.New("xor: strings must be of equal lengths")
	}
	hex1, err := hex.DecodeString(buff1)
	if err != nil {
		return nil, fmt.Errorf("xor: %w", err)
	}
	hex2, err := hex.DecodeString(buff2)
	if err != nil {
		return nil, fmt.Errorf("xor: %w", err)
	}
	xorBytes(hex1, hex2)
	encoded := make([]byte, hex.EncodedLen(len(hex2)))
	hex.Encode(encoded, hex1)
	return encoded, nil
}

func xorBytes(x, y []byte) {
	for i := 0; i < len(x); i++ {
		x[i] ^= y[i]
	}
}

// challenge: 3
func singleByteXOR(s string) ([]byte, []byte, error) {
	h1, err := hex.DecodeString(s)
	if err != nil {
		return nil, nil, fmt.Errorf("xorFreq: %w", err)
	}
	var maxScore int
	var cypherKey []byte
	msg := make([]byte, len(h1))
	for a := 32; a <= 126; a++ {
		// key len should be same as cypher len (bytes)
		key := bytes.Repeat([]byte{byte(a)}, len(h1))
		xorBytes(h1, key)
		score := freqCheck(h1)
		if score > maxScore {
			maxScore = score
			cypherKey = key
			copy(msg, h1)
		}
	}
	return cypherKey, msg, nil
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

func checkMsg(m []byte) bool {
	for _, i := range m {
		if 32 <= int(i) && int(i) <= 126 {
			continue
		}
		return false
	}
	return true
}

func main() {
	f, err := os.Open("./challenge4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("err:", err)
			return
		}
		k, m, err := singleByteXOR(string(line[:len(line)-1])) // remove \n
		if err != nil {
			fmt.Println("err:", err)
			continue
			//return
		}
		//isValid := checkMsg(m)
		//if !isValid {
		//continue
		//}
		fmt.Println("line trimmed:", string(line[:len(line)-1]))
		fmt.Println("key:", string(k))
		fmt.Println("msg:", string(m))
		fmt.Println(strings.Repeat("=", 20))
	}
}
