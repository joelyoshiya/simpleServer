// simpleServer is a simple web server that serves requests as a RESTful API
// it does not use any third party libraries
// it is a simple example of how to use the standard library when http is needed
// the server stores basic user info without passwords in memory

package main

// only standard library imports
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// define structures in memory on the server
type User struct {
	UserID  int    `json:"userid"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Users []User

func userInfo(w http.ResponseWriter, r *http.Request) {
	// get user info from database
	// return user info as json
	users := Users{
		User{UserID: 1, Name: "John", Surname: "Smith"},
		User{UserID: 2, Name: "Jane", Surname: "Doe"},
	}
	// get user id from request
	userId, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		fmt.Println("Error: ", err)
		// return error as json
		json.NewEncoder(w).Encode(err)
	}
	// find the user in the database
	for _, user := range users {
		if user.UserID == userId {
			// return user info as json
			fmt.Println("Endpoint Hit: getUserInfo")
			json.NewEncoder(w).Encode(user)
		}
	}
}

func allUserInfo(w http.ResponseWriter, r *http.Request) {
	// get all user info from database
	// return all user info as json
	users := Users{
		User{UserID: 1, Name: "John", Surname: "Smith"},
		User{UserID: 2, Name: "Jane", Surname: "Doe"},
	}
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(users)
}

// handler function to return server status
func serverStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running")
}

// main function to start the server
func handleRequests() {
	http.HandleFunc("/", serverStatus)
	http.HandleFunc("/userinfo", userInfo)
	http.HandleFunc("/alluserinfo", allUserInfo)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
