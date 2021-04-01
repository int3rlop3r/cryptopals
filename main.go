package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

// challenge: 1
func hexTo64(s string) (string, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("hexto64: %w", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func xor(buff1, buff2 string) (string, error) {
	if len(buff1) != len(buff2) {
		return "", errors.New("strings must be of equal lengths")
	}
	hex1, err := hex.DecodeString(buff1)
	if err != nil {
		return "", fmt.Errorf("hexto64: %w", err)
	}
	hex2, err := hex.DecodeString(buff2)
	if err != nil {
		return "", fmt.Errorf("hexto64: %s", err)
	}
	newBuff := make([]byte, len(hex1), len(hex2))
	for i := 0; i < len(hex1); i++ {
		newBuff[i] = hex1[i] ^ hex2[i]
	}
	return hex.EncodeToString(newBuff), nil
}

func main() {
	//s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	//str, err := hexTo64(s)
	//if err != nil {
	//fmt.Fprintln(os.Stderr, err)
	//return
	//}
	//fmt.Println(str)

	x := "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"
	z, err := xor(x, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(z)
}
