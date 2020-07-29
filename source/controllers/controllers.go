package controller

import (
	h "../helpers"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

type Hash struct {
	Phrase       string `json:"phrase"`
	HashedPhrase int64
}
type DetailedHash struct {
	Phrase       string `json:"phrase"`
	HexedPhrase  string
	HexedCuttedPhrase string
	HashedPhrase int64
}

var MainMethod = func(w http.ResponseWriter, r *http.Request) {
	var hash Hash
	//temporary variable to contain the req field in case it is not a string
	var tmpVal map[string]interface{}
	json.NewDecoder(r.Body).Decode(&tmpVal)
	//setup encoder instance
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	//in case we got not a string but a float64 or an object of any kind we will convert it to a string
	hash.Phrase = fmt.Sprintf("%v", tmpVal["phrase"])
	// check if we got nil with request, then response will have an error of type mismatch
	if hash.Phrase == "<nil>" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(w)
	} else {
		//get the sha256 sum
		sha256bytes := sha256.Sum256([]byte(hash.Phrase))
		// convert the first 8 bytes to an int64, since int64 fits only 8*8 bits
		hash.HashedPhrase = h.ByteArrayToInt(sha256bytes[:8])
		//encode it to JSON response
		encoder.Encode(&hash)
	}
}
//the same as above but with couple of additional fields
var DetailedMethod = func(w http.ResponseWriter, r *http.Request) {
	var detailedHash DetailedHash
	var tmpVal map[string]interface{}
	json.NewDecoder(r.Body).Decode(&tmpVal)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	detailedHash.Phrase = fmt.Sprintf("%v", tmpVal["phrase"])
	if detailedHash.Phrase == "<nil>" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(w)
	} else {
		sha256bytes := sha256.Sum256([]byte(detailedHash.Phrase))
		//convert the checksum to hex string
		detailedHash.HexedPhrase = h.B2hex(sha256bytes[:])
		//the same as above but passing only first 8 bites with slice
		detailedHash.HexedCuttedPhrase = h.B2hex(sha256bytes[:8])
		//converting it to int64
		detailedHash.HashedPhrase = h.ByteArrayToInt(sha256bytes[:8])
		encoder.Encode(&detailedHash)
	}
}