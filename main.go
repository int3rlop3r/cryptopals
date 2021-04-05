package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
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
	var maxScore float64
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
			//fmt.Println("score:", score, ", msg:", string(msg))
		}
	}
	return cypherKey, msg, nil
}

var freqs []float64 = []float64{0.37025, 0.067875, 0.123575, 0.19685, 0.5478, 0.105, 0.092325, 0.269875, 0.33295, 0.0047, 0.031425, 0.181325, 0.119025, 0.31665, 0.350075, 0.0829, 0.005125, 0.274425, 0.28625, 0.414675, 0.13115, 0.050475, 0.095475, 0.007875, 0.096325, 0.0032}

func freqCheck(s []byte) float64 {
	var score float64
	charFreq := make(map[byte]float64)
	sLower := bytes.ToLower(s)

	for _, ch := range sLower {
		if int(ch) < 32 || int(ch) > 126 {
			//fmt.Println("broke:", string(ch), ", b:", int(ch))
			return score // which is 0 at this point
		}
		_, ok := charFreq[ch]
		if ok { // already counted this character
			continue
		}
		charFreq[ch] = float64(bytes.Count(sLower, []byte{ch})) / float64(len(s))
	}

	for _, ch := range sLower {
		var freq float64
		pos := int(ch) - 97
		if pos < 0 || pos > 25 {
			freq = 0
		} else {
			freq = freqs[pos] // get freq of the character
		}
		score += math.Sqrt(charFreq[ch] * freq)
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
		isValid := checkMsg(m)
		if !isValid {
			continue
		}
		fmt.Println("line trimmed:", string(line[:len(line)-1]))
		fmt.Println("key:", string(k))
		fmt.Println("msg:", string(m))
		fmt.Println(strings.Repeat("=", 20))
	}
}
