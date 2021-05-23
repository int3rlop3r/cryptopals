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

var freqs []float64 = []float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, // A-G
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, // H-N
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, // O-U
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074, // V-Z
}

func isValidSymbol(c rune) bool {
	symbols := " .,;'"
	for _, s := range symbols {
		if c == s {
			return true
		}
	}
	//fmt.Printf("'%c' is not a vlid symbol.\n", c)
	return false
}

func freqCheck(s []byte) float64 {
	var score float64
	charFreq := make(map[byte]float64)
	sUpper := bytes.ToUpper(s)
	totalChars := float64(len(s))
	for _, ch := range sUpper {
		//if isValidSymbol(rune(ch)) {
		//continue
		//}

		//if ch < 'A' || ch > 'Z' {
		//return score // which is 0 at this point
		//}

		_, ok := charFreq[ch]
		if ok { // already counted this character
			continue
		}
		charFreq[ch] = float64(bytes.Count(sUpper, []byte{ch})) / totalChars
	}

	for ch, calcF := range charFreq {
		pos := ch - 'A'
		var freq float64
		if int(pos) > 25 { // coz 26 letters
			freq = 0
		} else {
			freq = freqs[pos] // get freq of the character
		}
		score += math.Sqrt(calcF * freq)
	}
	//fmt.Println("score:", score)
	return score
}

func checkMsg(m []byte) bool {
	for _, ch := range bytes.ToUpper(m) {
		if (ch < 'A' || ch > 'Z') && !isValidSymbol(rune(ch)) {
			fmt.Printf("%c is not valid\n", ch)
			return false
		}
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
