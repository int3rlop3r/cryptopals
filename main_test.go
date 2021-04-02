package main

import (
	"fmt"
	"os"
	"testing"
)

func TestChallenge1(t *testing.T) {
	expect := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	str, err := hexTo64(s)
	if err != nil {
		t.Errorf("ch1, err: %s", err)
	}

	if str != expect {
		t.Errorf("expected: %s, got: %s", expect, str)
	}
}

func TestChallenge2(t *testing.T) {
	expect := "746865206b696420646f6e277420706c6179"
	x := "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"
	z, err := xor(x, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if z != expect {
		t.Errorf("expected: %s, got: %s", expect, z)
	}
}

func TestChallenge3(t *testing.T) {
	expect := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	cypher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, err := xorFreq(cypher)
	if err != nil {
		t.Error(err)
	}
	if string(key) != expect {
		t.Errorf("expected: %s, got: %s", expect, key)
	}
}
