package main

import (
	"crypto/sha256"

	"encoding/json"

	"net/http"
	h "./src/helpers"
)

type Pswd struct {
	Phrase       string `json:"phrase"`
	HexedS	string
	HashedPhrase uint64
}

func main() {
	http.HandleFunc("/get-phrase-hash", func(w http.ResponseWriter, r *http.Request) {
		var pass Pswd
		json.NewDecoder(r.Body).Decode(&pass)
		sha256bytes := sha256.Sum256([]byte(pass.Phrase))
		hexedStr := h.B2hex(sha256bytes[:])
		pass.HashedPhrase = h.Hex2int(hexedStr)
		//p := fmt.Sprintf("%s", pass.Phrase)
		json.NewEncoder(w).Encode(Pswd{
			pass.Phrase,
			hexedStr,
			pass.HashedPhrase,
		})
	})

	http.ListenAndServe(":8081", nil)
}
