package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	en := flag.String("encode", "", "Base64 Encode a string.")
	dc := flag.String("decode", "", "Decode a Base64 encoded string.")
	flag.Parse()
	if *en != "" {
		sEnc := b64.StdEncoding.EncodeToString([]byte(*en))
		fmt.Println(sEnc)
		os.Exit(0)
	}

	if *dc != "" {
		sDec, _ := b64.StdEncoding.DecodeString(*dc)
		fmt.Println(string(sDec))
	}
}
