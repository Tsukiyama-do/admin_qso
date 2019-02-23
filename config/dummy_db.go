package config

import (
	"../crypto"
//	"errors"
//	"log"

	"strconv"
)

func NewDummyUser(username, email string) *DummyUserModel {
	return &DummyUserModel{
		Username: username,
		Email: email,
	}
}

type DummyUserModel struct {
	Username string
	Password string
	Email string
	authenticated bool
}

func (u *DummyUserModel) SetPassword(password string) error {
	hash, err := crypto.PasswordEncrypt(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *DummyUserModel) Authenticate() {
	u.authenticated = true
}

type DummyDatabase struct {
	database map[string]interface{}
}

var store DummyDatabase

func init() {
	store.database = map[string]interface{}{}
}

func DummyDB() *DummyDatabase {
	return &store
}

func (db *DummyDatabase) Exists(username string) bool {

//	_, r := db.database[username]

	mm :=  DummyUserModel{
    Username : username,
   }

	if err := mm.CheckUser() ; err != nil { return false }

	return true
}

func (db *DummyDatabase) SaveUser(username, email, password string) error {
//	if db.Exists(username) {
//		return errors.New("user \"" + username + "\" already exists")
//	}

	mm := DummyUserModel{
      Username : username,
   }

	if err := mm.CheckUser() ; err != nil {
		if err.Error() != "No_Data_Found" {
			return err
		}
	}

	user := NewDummyUser(username, email)
	if err := user.SetPassword(password); err != nil {
		return err
	}
	if err := user.InsertUpload() ; err != nil {  return err}

	return nil
}

func (db *DummyDatabase) GetUser(username, password string) (*DummyUserModel, error) {
//	buffer, exists := db.database[username]
//	if !exists {
//		return nil, errors.New("user \"" + username + "\" doesn't exists")
//	}

  mm := DummyUserModel{
     Username : username,
		 Password : password,
   }

	if err := mm.CheckUser() ; err != nil {
		return &mm, err
	}

//	if  err := mm.SetPassword(password); err != nil {
//		return &mm, err
//	}

	if err := mm.UserAuth() ; err != nil {  return &mm, err }

	return &mm, nil
}

/////////////////////
///  QSLCARDS methods
/////////////////////



func (db *DummyDatabase) SelectQSL() (string, error) {
//	buffer, exists := db.database[username]
//	if !exists {
//		return nil, errors.New("user \"" + username + "\" doesn't exists")
//	}

  var mm QslCardsModel
//	var mj S_QSLCARDS

	var s_json_cal string

	mj, err := mm.QSLRecords()
	if err != nil {  return s_json_cal, err 	}

//	log.Printf("QSLRec len is %d", len(mj))
//	log.Printf("QSLRec is %v", mj[0])

	s_json_cal = `[ `
	for _, items := range mj {
			s_json_cal = s_json_cal + `{ "ID" : "` + strconv.Itoa(items.ID) + `" , "CALLSIGN" :  "` + items.CALLSIGN + `" , "DATETIME":  "` + items.DATETIME + `" , "FILES":  "` + items.FILES + `" },`
	}
/*
	for _, items := range mj {
			log.Printf( `{ "ID" : "` + strconv.Itoa(items.ID) + `" , "CALLSIGN":  "` + items.CALLSIGN + `" , "DATETIME":  "` + items.DATETIME + `" , "FILES":  "` + items.FILES + `" },` )
	}
*/
	if len(s_json_cal) > 0 {        //  文字列末尾の　, を削除している。
		 s_json_cal = string(s_json_cal[:(len(s_json_cal)-1)])
	}

	s_json_cal = s_json_cal + `] `

	return s_json_cal, nil
}


func (db *DummyDatabase) UpdateQSL(id string, callsign string, datetime string, files string) error {

	i_id, err := strconv.Atoi(id)
	err = nil

  mm := QslCardsModel{
			ID : i_id ,
			CALLSIGN : callsign ,
			DATETIME :  datetime ,
			FILES : files ,
			iud : "U" ,
	}

	err = mm.QSLUpdate()
	if err != nil {  return err 	}

	return nil

}


func (db *DummyDatabase) DeleteQSL(id string) error {

	i_id, err := strconv.Atoi(id)
	err = nil

  mm := QslCardsModel{
			ID : i_id ,
			iud : "D" ,
	}

	err = mm.QSLDelete()
	if err != nil {  return err 	}

	return nil

}


func (db *DummyDatabase) InsertQSL(id string, callsign string, datetime string, files string) error {

	i_id, err := strconv.Atoi(id)
	err = nil

  mm := QslCardsModel{
			ID : i_id ,
			CALLSIGN : callsign ,
			DATETIME :  datetime ,
			FILES : files ,
			iud : "I" ,
	}

	err = mm.QSLCheck()
	if err != nil {  return err 	}

	err = mm.QSLInsert()
	if err != nil {  return err 	}


	return nil

}
