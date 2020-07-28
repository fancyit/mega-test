package controller

import (
	h "../helpers"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

type Pswd struct {
	Phrase       string `json:"phrase"`
	HexedPhrase  string
	HashedPhrase uint64
}

var MainMethod = func(w http.ResponseWriter, r *http.Request) {
	var pass Pswd
	var tmpVal map[string]interface{} // проместим входящую переменную в тип мап(ключ: значение), нужный нам ключ мы знаем
	json.NewDecoder(r.Body).Decode(&tmpVal)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	pass.Phrase = fmt.Sprintf("%v", tmpVal["phrase"]) //если на вход нам дали число(float64) - парсим его в строку
	if pass.Phrase == "<nil>" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(w)
	} else {
		sha256bytes := sha256.Sum256([]byte(pass.Phrase)) //получаем хэш в виде массива байт
		pass.HexedPhrase = h.B2hex(sha256bytes[:])        //хексим слайс массива
		pass.HashedPhrase = h.Hex2int(pass.HexedPhrase)
		encoder.Encode(&pass)
	}
}
