package glacier

import (
	"bytes"
	"testing"
)

func TestTreeHash(t *testing.T) {
	var in bytes.Buffer
	in.WriteString("Hello World")
	out1 := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	tree, err := createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out1 {
		t.Fatal("wanted:", out1, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB; i++ {
		in.WriteByte('a')
	}
	out2 := "9bc1b2a288b26af7257a36277ae3816a7d4f16e89c1e7e77d0a5c48bad62b360"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out2 {
		t.Fatal("wanted:", out2, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB; i++ {
		in.WriteByte('a')
	}
	in.WriteString("Hello World")
	out3 := "7a398c79d8fc266cde4766b105d56a49361b22142aaa35a22ef505660c7edf59"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out3 {
		t.Fatal("wanted:", out3, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*2; i++ {
		in.WriteByte('a')
	}
	out4 := "3d95be8a6b55f83b93db329b8657ef6e8496361cb7b9a882b263fb2fb6197564"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out4 {
		t.Fatal("wanted:", out4, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*3; i++ {
		in.WriteByte('a')
	}
	out9 := "d96d417c167a3bef2eaab2e637095834c2023e867f8287b6e5aa4da66eb0a555"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out9 {
		t.Fatal("wanted:", out9, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*2; i++ {
		in.WriteByte('a')
	}
	in.WriteString("Hello World")
	out5 := "7cf44f7e83180f709ad6f8376dd704609d28a117f3a1878c301bc9e78c870344"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out5 {
		t.Fatal("wanted:", out5, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*4; i++ {
		in.WriteByte('a')
	}
	out6 := "ab7404d312438a79f8164a6714f2b7aa42b52fbfc8f3b2db62b372b182fc6619"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out6 {
		t.Fatal("wanted:", out6, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*5; i++ {
		in.WriteByte('a')
	}
	out7 := "ede4a38ec783a37bd243daa9693fbabaaf7ea2a370a511a72b02aaa8bfe0dda6"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out7 {
		t.Fatal("wanted:", out7, "got:", got)
	}

	in.Reset()
	for i := 0; i < MiB*7; i++ {
		in.WriteByte('a')
	}
	out8 := "3ff5623fb1ab0ef60adab85d698aa03560d24890e51d7d69d3296dccfb774ac7"
	tree, err = createTreeHash(&in)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(toHex(tree.Hash[:])); got != out8 {
		t.Fatal("wanted:", out8, "got:", got)
	}
}
