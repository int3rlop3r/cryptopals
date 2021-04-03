package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestChallenge1(t *testing.T) {
	expect := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	x, err := hexTo64(s)
	if err != nil {
		t.Errorf("ch1, err: %s", err)
	}

	if !bytes.Equal(expect, x) {
		t.Errorf("expected: '%s', got: '%s'", expect, x)
	}
}

func TestChallenge2(t *testing.T) {
	expect := []byte("746865206b696420646f6e277420706c6179")
	x := "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"
	z, err := fixedXOR(x, y)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for i := range expect {
		if expect[i] != z[i] {
			t.Errorf("expected: %c, got: %c", expect[i], z[i])
		}
	}
	if !bytes.Equal(expect, z) {
		t.Errorf("expected: '%x', got: '%x'", expect, z)
	}
}

func TestChallenge3(t *testing.T) {
	expect_key := []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	expect_msg := []byte("Cooking MC's like a pound of bacon")
	cypher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, msg, err := singleByteXOR(cypher)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(expect_key, key) || !bytes.Equal(expect_msg, msg) {
		t.Errorf("expected key: '%x', got key: '%x',\nexpected msg: '%x', got msg: '%x'\n", expect_key, key, expect_msg, msg)
	}
}
