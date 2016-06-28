package users

import (
	"models"
	"strings"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type UserName struct{
	Name string
}
type RshipState struct{
	State string
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Select all users.
	users, err := models.GetUsers()

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte("json error!\n"))
	}
	fmt.Println("GetAllUsers")
	w.Write(b)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var name UserName
	
	postData,err:=ioutil.ReadAll(r.Body)
    fmt.Println("CreateUser")
    if err != nil {
    	fmt.Println("err" + err.Error())
    }
    
	fmt.Println(string(postData))
	err = json.Unmarshal(postData, &name)
	if err != nil {
		w.Write([]byte("UnMarshal json error!\n"))
		return
	}
	u := &models.User{
		Name: name.Name,
		Type: "user",
	}
	err = models.CreateUser(u)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else{
		b,_:=json.Marshal(u)
		w.Write(b)
	}
}

func GetRelationships(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRelationships")
	fmt.Println(r.URL)
	strs := strings.Split(r.URL.String(), "/")
	uid := strings.TrimSpace(strs[2])
	fmt.Println(uid)
	rships, err:=models.GetRshipById(uid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b, err:=json.Marshal(rships)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func CreateRelationships(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rstate RshipState
	fmt.Println("CreateRelationships")
	// get the uid and the other uid from url
	strs := strings.Split(r.URL.String(), "/")
	uid := strings.TrimSpace(strs[2])
	oid := strings.TrimSpace(strs[4])
	// read the post body
	postData,err:=ioutil.ReadAll(r.Body)
    if err != nil {
    	fmt.Println("err" + err.Error())
    }
    // print the post body
	fmt.Println(string(postData))
	err = json.Unmarshal(postData, &rstate)
	if err != nil {
		w.Write([]byte("UnMarshal json error!\n"))
		return
	}
	reqship := &models.Relationship {
			Id:uid,
			User_id:oid,
			State:rstate.State,
			Type:"relationship",
	}

	tmpRRship, _ := models.GetRshipByRid(oid, uid)

	if tmpRRship != nil {
		// if all liked, all 'matched'
		if reqship.State == "liked" && tmpRRship.State == "liked" {
			tmpRRship.State = "matched"
			err = models.CreateOrUpdteRelationship(tmpRRship)
			reqship.State = "matched"
		} else if reqship.State == "disliked" && tmpRRship.State == "matched" {
			tmpRRship.State = "liked"
			err = models.CreateOrUpdteRelationship(tmpRRship)
		}
	}
	
	err = models.CreateOrUpdteRelationship(reqship)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else{
		b,_:=json.Marshal(reqship)
		w.Write(b)
	}
}