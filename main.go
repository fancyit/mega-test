package main

import (
	"encoding/json"
	"hash/crc64"
	"net/http"
)

type Pswd struct {
	Phrase       string `json:"phrase"`
	HashedPhrase uint64
}

func main() {
	http.HandleFunc("/get-phrase-hash", func(w http.ResponseWriter, r *http.Request) {
		var pass Pswd
		json.NewDecoder(r.Body).Decode(&pass)
		crc64Table := crc64.MakeTable(0xC96C5795D7870F42)
		crc64Int := crc64.Checksum([]byte(pass.Phrase), crc64Table)
		pass.HashedPhrase = crc64Int
		json.NewEncoder(w).Encode(pass)
	})

	http.ListenAndServe(":8081", nil)
}
