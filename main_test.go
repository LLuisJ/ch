package main

import "testing"

func TestHashBytes(t *testing.T) {
	content := []byte("Test")
	algorithms := []string{"MD4", "MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512", "SHA3-224", "SHA3-256",
		"SHA3-384", "SHA3-512"}
	expected := []string{
		"5d9381b52cd41a87bc8fcc7726aec03a",
		"0cbc6611f5540bd0809a388dc95a615b",
		"640ab2bae07bedc4c163f679a746f7ab7fb5d1fa",
		"3606346815fd4d491a92649905a40da025d8cf15f095136b19f37923",
		"532eaabd9574880dbf76b9b8cc00832c20a6ec113d682299550d7a6e0f345e25",
		"7b8f4654076b80eb963911f19cfad1aaf4285ed48e826f6cde1b01a79aa73fadb5446e667fc4f90417782c91270540f3",
		"c6ee9e33cf5c6715a1d148fd73f7318884b41adcb916021e2bc0e800a5c5dd97f5142178f6ae88c8fdd98e1afb0ce4c8d2c54b5f37b30b7da1997bb33b0b8a31",
		"d40cc4f9630f21eef0b185bdd6a51eab1775c1cd6ae458066ecaf046",
		"c0a5cca43b8aa79eb50e3464bc839dd6fd414fae0ddf928ca23dcebf8a8b8dd0",
		"da73bfcba560692a019f52c37de4d5e3ab49ca39c6a75594e3c39d805388c4de9d0ff3927eb9e197536f5b0b3a515f0a",
		"301bb421c971fbb7ed01dcc3a9976ce53df034022ba982b97d0f27d48c4f03883aabf7c6bc778aa7c383062f6823045a6d41b8a720afbb8a9607690f89fbe1a7",
	}
	for index := range algorithms {
		result, err := hashBytes(content, algorithms[index])
		if err != nil {
			t.Fatal(err)
		}
		if expected[index] != *result {
			t.Fatalf("test %d has failed\nexpected: %s\ngot: %s\n", index+1, expected[index], *result)
		}
	}
}
