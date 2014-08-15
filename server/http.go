package main

import (
	"bytes"
	"crypto/sha512"
	"log"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	mrand "math/rand"
	"time"
)

const CONFIG string = "applikitchen:meowmix@tcp(127.0.0.1:3306)/applikitchen?parseTime=true&collation=utf8_general_ci"

func Hex(chunks int) string {
	var buffer bytes.Buffer

	bytes := make([]byte, 4)
	for i := 0; i < chunks; i++ {
		binary.LittleEndian.PutUint32(bytes, mrand.Uint32())
		buffer.WriteString(hex.EncodeToString(bytes))
	}

	return buffer.String()
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	// pull email from route for username
	vars := mux.Vars(r)
	email := vars["email"]

	// pull password from POST param
	password := r.FormValue("password")
	if len(password) == 0 {
		fmt.Fprintf(w, "missing password parameter")
		return
	}

	hashPassBytes := sha512.Sum512([]byte(password))   // this gives us an array
	hashPassHex := hex.EncodeToString(hashPassBytes[:])      // here we take byte slice and hex it

	dbc, err := sql.Open("mysql", CONFIG)
	if err != nil {
		return
	}

	id_rows := dbc.QueryRow("SELECT id FROM users WHERE email = ? and password = ?", email, hashPassHex)

	var id int64

	err = id_rows.Scan(&id)

	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}

	// Get the expiration time if it exists for that user_id
	session_rows := dbc.QueryRow("SELECT expiration FROM sessions WHERE user_id = ?", id)

	var expiration time.Time

	err = session_rows.Scan(&expiration)

	var sessionToken string

	// Failure here indicates they have no active session
	if err != nil {
		// Session didn't exist, so create one
		sessionToken = Hex(4)  // Create new token
		_, err := dbc.Exec(
			"INSERT INTO sessions (session_id, user_id, expiration) values (?, ?, NOW())",
			sessionToken,
			id,
		)
		if err != nil {
			fmt.Fprintf(w, "failed to insert new session id")
			return
		}
	} else {
		//fmt.Print(w, expiration)	
	}

	// Check if the ID exists in the sessions table
	// If the ID exists, find out if it's expired
	// Iff it's expired, handle it
	// If the ID does not exist, create a new token and handle it

	response := struct {
		Success bool
		Token   string
	}{
		true,
		sessionToken,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		fmt.Fprintf(w, "could not write json")
		return
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/auth/{email}", authHandler)
	router.HandleFunc("/ping/", pingHandler)

	http.Handle("/", router)
	//if err := http.ListenAndServe("172.31.0.230:8080", nil); err != nil {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
