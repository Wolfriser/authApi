package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var users map[email]User
var userTokens map[string]email

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
	userTokens[token] = data.Email
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var registrationData User
	err := json.Unmarshal(body, &registrationData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	user := users[registrationData.Email]
	u := User{}
	if user != u {
		http.Error(w, "This email is already registered", http.StatusConflict)
		return
	}
	users[registrationData.Email] = registrationData
	response := Response{
		Success: true,
		Message: "Registration successful",
	}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(j)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := cookie.Value
	email := userTokens[token]
	if email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var UpdateData User
	err = json.Unmarshal(body, &UpdateData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	
}
