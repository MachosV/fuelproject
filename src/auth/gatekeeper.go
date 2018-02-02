package auth

import (
	"encoding/base64"
	"math/rand"
	"fmt"
	"storage"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"bytes"
	"models"
	"time"
	"log"
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
func DoLogin(username string, password string) string{
	var db = storage.GetDb()
	var form_password,err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var system_password []byte
	results, err := db.Query("SELECT password FROM accounts" +
		"WHERE" +
		"username="+username+" or " +
		"email="+username)
	//nomizw edw skaei
	if err == sql.ErrNoRows{
		return ""
	}
	for results.Next(){
		results.Scan(system_password)
	}
	if !bytes.Equal(form_password, system_password){
		return ""
	}
	return CreateSession()
}

func DoLogout(id string){
	delete(sessions,id)
}

func CreateSession() string{
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
        return bdecoded
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
		log.Println("Session GC started")
		for key,_ := range sessions{
			t1 := sessions[key].GetTimestamp()
			if time.Since(t1).Minutes() > 1{
				delete(sessions,key)
			}
		}
	}
}