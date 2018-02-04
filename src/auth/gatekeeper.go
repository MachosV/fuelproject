package auth

import (
	"encoding/base64"
	"math/rand"
	"fmt"
	"storage"
	"golang.org/x/crypto/bcrypt"
	"models"
	"time"
	"net/http"
)

var sessions map[string]*models.Session

func init(){
	sessions = make(map[string]*models.Session)
	go SessionGC()
}

func CheckAuth(sessionid string) bool {
	if _, exists := sessions[sessionid]; !exists{
			return false
		}
	return true
}

//Returns a session string.
//Otherwise returns empty string.
func DoLogin(username string, password string) *http.Cookie{
	var db = storage.GetDb()
	var system_password []byte
	var form_password = []byte(password)
	results, err := db.Query("SELECT password FROM accounts " +
		"WHERE " +
		"username=? or " +
		"email=?;",username,username)
	if err != nil{
		return nil
	}
	for results.Next(){
		results.Scan(&system_password)
	}
	err = bcrypt.CompareHashAndPassword(system_password,form_password)
	if err != nil{
		return nil
	}
	return CreateSession()
}

func DoLogout(id string){
	delete(sessions,id)
}

func CreateSession() *http.Cookie{
	var b []byte
	var bdecoded string
	for {
		b = make([]byte, 32)
        	_,err := rand.Read(b)
		if err != nil {
			fmt.Print(err)
		}
		bdecoded = base64.URLEncoding.EncodeToString(b)
		if _, exists := sessions[bdecoded]; !exists{
			break
		}
	}
	var session = new(models.Session)
	session.Init()
	session.SetId(bdecoded)
	session.SetTimestamp()
	sessions[bdecoded] = session
	var cookie *http.Cookie = &http.Cookie{
		Name: "sessionid",
		Value: bdecoded,
		Expires: time.Now().Add(2*time.Minute),
		HttpOnly: true,
	}
        return cookie
}

func SetSession(sessionid string, key string, value string){
	var session = sessions[sessionid]
	session.SetValue(key,value)
}

func GetSession(sessionid string) *models.Session{
	return sessions[sessionid]
}

func DeleteSession(sessionid string){
	delete(sessions,sessionid)
}

func SessionGC(){
	for{
		time.Sleep(10 * time.Second)
		for key,_ := range sessions{
			t1 := sessions[key].GetTimestamp()
			if time.Since(t1).Minutes() > 2{
				delete(sessions,key)
			}
		}
	}
}