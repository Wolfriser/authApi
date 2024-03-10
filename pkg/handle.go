package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

var temp = make(map[email]User)
var userTokens = make(map[string]email)
var films = make(map[string]Film)
var users = Admin(temp)

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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if user == u || err != nil {
		http.Error(w, "Wrong email or password", http.StatusUnauthorized)
		// log.Println(err.Error())
		return
	}
	token, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: token.String(),
		Path:  "/",
	})
	userTokens[token.String()] = data.Email
	user.Password = ""
	log.Println(token)
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
	if registrationData.IsAdmin {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	user := users[registrationData.Email]
	u := User{}
	if user != u {
		http.Error(w, "This email is already registered", http.StatusConflict)
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(registrationData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	registrationData.Password = string(hashedPass)
	users[registrationData.Email] = registrationData
	registrationData.Password = ""
	response := Response{
		Success: true,
		Message: "Registration successful",
		User:    registrationData,
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
	email, ok := userTokens[cookie.Value]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var updateData User
	err = json.Unmarshal(body, &updateData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	users[email] = updateData
	user := users[email]
	user.Password = ""
	response := Response{
		Success: true,
		Message: "User data updated successfully",
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

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	email, ok := userTokens[cookie.Value]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if !users[email].IsAdmin {
		http.Error(w, "Forbidden: User does not have admin access", http.StatusForbidden)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var filmData Film
	err = json.Unmarshal(body, &filmData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	films[filmData.Name] = filmData
	response := Response{
		Success: true,
		Message: "Film data uploaded successfuly",
	}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(j)
}

func Admin(m map[email]User) map[email]User {
	pass, err := bcrypt.GenerateFromPassword([]byte("0000"), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	adminData := User{
		Email:    "admin",
		Password: string(pass),
		IsAdmin:  true,
	}
	m["admin"] = adminData
	return m
}
