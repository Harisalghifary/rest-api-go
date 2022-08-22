package apps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Harisalghifary/rest-api-go/models"
	"github.com/Harisalghifary/rest-api-go/utils"
	"github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Language using Gorilla Mux and Cassandra DB!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var Newuser models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user name and password")
	}
	json.Unmarshal(reqBody, &Newuser)
	if err := utils.Session.Query("INSERT INTO users (name, username, password, email) VALUES (?, ?, ?, ?)", Newuser.Name, Newuser.Username, Newuser.Password, Newuser.Email).Exec(); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(Newuser, "", " ")

	fmt.Fprintf(w, "%s", string(Conv))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	m := map[string]interface{}{} // map[string]interface{} is a map[string]interface{}

	iter := utils.Session.Query("SELECT * FROM users").Iter() // iteration
	for iter.MapScan(m) {                                     // scan the data
		Users = append(Users, models.User{
			ID:       m["id"].(int),
			Name:     m["name"].(string),
			Username: m["username"].(string),
			Password: m["password"].(string),
			Email:    m["email"].(string),
		})
		m = map[string]interface{}{}
	}

	if err := iter.Close(); err != nil {
		panic(err)
	}

	Conv, _ := json.MarshalIndent(Users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	UserID := mux.Vars(r)["id"] // get the id from the url

	var User []models.User

	m := map[string]interface{}{}

	iter := utils.Session.Query("SELECT * FROM users WHERE id = ?", UserID).Iter()
	for iter.MapScan(m) {
		User = append(User, models.User{
			ID:       m["id"].(int),
			Name:     m["name"].(string),
			Username: m["username"].(string),
			Password: m["password"].(string),
			Email:    m["email"].(string),
		})
		m = map[string]interface{}{}
	}

	if err := iter.Close(); err != nil {
		panic(err)
	}
	Conv, _ := json.MarshalIndent(User, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID := mux.Vars(r)["id"]

	var UpdateUser models.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with correct property")
	}
	json.Unmarshal(reqBody, &UpdateUser)
	if err := utils.Session.Query("UPDATE users SET name = ?, username = ?, password = ?, email = ? WHERE id = ?", UpdateUser.Name, UpdateUser.Username, UpdateUser.Password, UpdateUser.Email, UserID).Exec(); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	Conv, _ := json.MarshalIndent(UpdateUser, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func DeleteOneUser(res http.ResponseWriter, req *http.Request) {
	UserID := mux.Vars(req)["id"]
	if err := utils.Session.Query("DELETE FROM users WHERE id = ?", UserID).Exec(); err != nil {
		panic(err)
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "User with ID %s is deleted successfully", UserID)

}
