package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var users map[email]User
var userTokens map[email]string

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)

	var data LoginData
	err := json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	user := users[data.Email]
	u := User{}
	if user == u || user.Password != data.Password {
		http.Error(w, "Wrong email or password", http.StatusUnauthorized)
		// log.Println(err.Error())
		return
	}
	token := "dfghjk"
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: token,
		Path:  "/",
	})
	userTokens[data.Email] = token
	response := Response{
		Success: true,
		Message: "Login successful",
		User:    user,
	}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(j)
}

func Register(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}
