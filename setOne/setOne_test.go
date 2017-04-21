package main

import "testing"

func TestHexToBase64(t *testing.T) {
	testCase := struct {
		in  string
		out string
	}{
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
	}
	got := HexToBase64(testCase.in)
	if got != testCase.out {
		t.Errorf("HexToBase64 returned %s, want %s", got, testCase.out)
	}
}

func TestFixedXOR(t *testing.T) {
	testCase := struct {
		in  string
		key string
		out string
	}{
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
		"746865206b696420646f6e277420706c6179",
	}
	got := FixedXOR(testCase.in, testCase.key)
	if got != testCase.out {
		t.Errorf("FixedXOR returned %s, want %s", got, testCase.out)
	}
}

func TestSingleByteXOR(t *testing.T) {
	testCase := struct {
		in  string
		out string
		key int
	}{
		"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
		"Cooking MC's like a pound of bacon",
		88,
	}
	gotKey, gotOut := SingleByteXOR(testCase.in)
	if gotKey != testCase.key || gotOut != testCase.out {
		t.Errorf("SingleByteXOR returned %d, %s, want %d, %s", gotKey, gotOut, testCase.key, testCase.out)
	}
}

func TestDetectSingleByteXOR(t *testing.T) {
	testCase := struct {
		filepath string
		out      string
		key      int
	}{
		"./4.txt",
		"Now that the party is jumping\n",
		53,
	}
	gotKey, gotOut := DetectSingleByteXOR(testCase.filepath)
	if gotKey != testCase.key || gotOut != testCase.out {
		t.Errorf("DetectSingleByteXOR returned %d, %s, want %d, %s", gotKey, gotOut, testCase.key, testCase.out)
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	testCase := struct {
		in  string
		key string
		out string
	}{
		"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
		"ICE",
		"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
	}
	got := RepeatingKeyXOR(testCase.key, testCase.in)
	if got != testCase.out {
		t.Errorf("RepeatingKeyXOR returned %s, want %s", got, testCase.out)
	}
}
