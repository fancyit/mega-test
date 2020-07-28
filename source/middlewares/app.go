package app

import (
	res "../models"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

var UrlCheker = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		urls := []string{"/get-phrase-hash"} //Список доступных методов (на вырост ;-))
		requestPath := r.URL.Path //текущий путь запроса

		/* проверяем, существует ли запрашиваемая страница
		при обращении через браузер пользователь получит заглушку и возможность почитать
		описание существующих на данный момент методов */
		for _, value := range urls {

			if value != requestPath {
				fs := http.FileServer(http.Dir("../static/oops/"))
				response := make(map[string] interface{})
				response = res.Message(404, "No such page")
				w.WriteHeader(http.StatusNotFound)
				w.Header().Add("Content-Type", "application/json")
				res.Respond(w, response)
				http.Handle("/static/", http.StripPrefix("/static/", fs))
				return
			}
		}
		next.ServeHTTP(w, r)
	});
}

var ReqParser = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var val map[string]interface{}
		var err string = ""
		phrase := val["phrase"]
		json.NewDecoder(r.Body).Decode(&val)
		inputType := reflect.TypeOf(phrase).Name()
		switch inputType {
		case "float64":
			phrase = fmt.Sprintf("%v", phrase)
		case "string":
			phrase = fmt.Sprintf("%v", phrase)
		default:
			err = fmt.Sprintf("Such data type as %v not allowed", reflect.TypeOf(phrase).Name())
		}
		response := make(map[string] interface{})
		response = res.Message(415, err)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		res.Respond(w, response)
		return
	});
}
