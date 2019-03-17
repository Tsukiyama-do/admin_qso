package routes

import (
	"../config"
	"../sessions"
	"../env"

	"net/http"
	"github.com/gin-gonic/gin"
	"log"
  "io/ioutil"
	"path/filepath"
	"encoding/json"


	"os"
	"strconv"
//  "github.com/bitly/go-simplejson"
	"fmt"
//	"bytes"
)

const up_path string = "../qso/downloads/"


func Qslmain(ctx *gin.Context) {

// Initial step : To confirm sessions
  var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
//		println("Unknown sesssion!  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": env.S_host(), })
		return
	}

	user = buffer.(*config.DummyUserModel)
//	println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

	// Initial step : End To confirm sessions

	ctx.HTML(http.StatusOK, "qslmain.html", gin.H{
		"isLoggedIn": true,
		"username": user.Username,
		"email": user.Email,
		"domainport":  env.S_host(),
	})

//	ctx.Redirect(http.StatusSeeOther, "/")

}

func Qslselectp(ctx *gin.Context) {

//	println("qsl/qslselectp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": env.S_host(), })
			return
		}

		user = buffer.(*config.DummyUserModel)

//		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

		if user.Username == "" {
		}

		// Initial step : End To confirm sessions
		// Retrieval step : To get records and convert them to JSON


	db := config.DummyDB()
	sl_json_cal, err := db.SelectQSL()
	if err != nil {
		println("Error: " + err.Error())
		return
	}


	ctx.JSON(http.StatusOK, gin.H{ "results" : sl_json_cal })
}
/*
type qsl2_rec struct{
	I     bool    `json:"I"`
	U     bool    `json:"U"`
	D     bool    `json:"D"`
	ID     int  `json:"ID"`
	CALLSIGN string `json:"CALLSIGN"`
	DATETIME string `json:"DATETIME"`
	FILES string `json:"FILES"`
}
*/

type Rec struct {
	mode string      `json:"mode"`
	id string        `json:"id"`
	callsign string  `json:"callsign"`
	datetime string  `json:"datetime"`
	files string     `json:"files"`
}


func Qslupddelp(ctx *gin.Context) {

//	println("qsl/qslupddelp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
//			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport":  env.S_host(), })
			return
		}

		user = buffer.(*config.DummyUserModel)
//		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)
		if user.Username == "" {
		}

		// Initial step : End To confirm sessions
		// Parsing step : To get records and convert them to JSON

			body := ctx.Request.Body
			x, _ := ioutil.ReadAll(body)


//		 	b := string(buf[0:n])
//		 	log.Println("Request body : " + string(x))



			var rec []interface{}
			var rec2 []Rec
			if err := json.Unmarshal(x, &rec); err != nil {
				log.Println(err)
				log.Println("Error occured at JSON conversion.")
				ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at JSON conversion")})
				return
			}

			for _ , p := range rec {
//				log.Printf(" E %v \n", p)
//				log.Printf("%d  %s %s %s %s %s \n", idx,
 //							p.(map[string]interface{})["mode"].(string),p.(map[string]interface{})["id"].(string), p.(map[string]interface{})["callsign"].(string), p.(map[string]interface{})["datetime"].(string), p.(map[string]interface{})["files"].(string) )

					wkr := Rec{p.(map[string]interface{})["mode"].(string), p.(map[string]interface{})["idno"].(string),  p.(map[string]interface{})["callsign"].(string), p.(map[string]interface{})["datetime"].(string),  p.(map[string]interface{})["files"].(string) }

						rec2 = append(rec2, wkr)

			}


		// Run sql proc
		db := config.DummyDB()

    for _ , q := range rec2 {
//				log.Printf("json( %d ) is %s %s %s \n", idx, q.mode, q.id, q.callsign  )

				if q.mode == "U" {
						err := db.UpdateQSL(q.id, q.callsign, q.datetime, q.files)
						if err != nil {
							log.Println("Error is " + err.Error())
							ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at Update QSL")})
							return
						}
				}
				if q.mode == "D" {
						err := db.DeleteQSL(q.id)
						if err != nil {
							log.Println("Error is " + err.Error())
							ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at Update QSL")})
							return
						}
				}
			}


			ctx.JSON(http.StatusOK, gin.H{ "results" : "Update Delete success!" })

}


func Qslinsertp(ctx *gin.Context) {

//	println("qsl/qslinsertp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport":  env.S_host(), })
			return
		}

		user = buffer.(*config.DummyUserModel)
//		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)
		if user.Username == "" {
		}

		// Initial step : End To confirm sessions
		// Parsing step : To get records and convert them to JSON

		body := ctx.Request.Body
		x, _ := ioutil.ReadAll(body)

//		log.Println("Request body : " + string(x))

		var rec []interface{}
		var rec2 []Rec
		if err := json.Unmarshal(x, &rec); err != nil {
			log.Println(err)
			log.Println("Error occured at JSON conversion.")
			ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at JSON conversion")})
			return
		}

		for _ , p := range rec {
//			log.Printf("%d  %s %s %s %s %s \n", idx,
//						p.(map[string]interface{})["mode"].(string),p.(map[string]interface{})["id"].(string), p.(map[string]interface{})["callsign"].(string), p.(map[string]interface{})["datetime"].(string), p.(map[string]interface{})["files"].(string) )

					rec2 = append(rec2,Rec{p.(map[string]interface{})["mode"].(string),p.(map[string]interface{})["idno"].(string), p.(map[string]interface{})["callsign"].(string), p.(map[string]interface{})["datetime"].(string), p.(map[string]interface{})["files"].(string)})

		}


	// Run sql proc
	db := config.DummyDB()

	for _ , q := range rec2 {
//			log.Printf("json( %d ) is %s %s %s \n", idx, q.mode, q.id, q.callsign  )

			if q.mode == "I" {
					err := db.InsertQSL(q.id, q.callsign, q.datetime, q.files)
					if err != nil {
						log.Println("Error is " + err.Error())
						ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at Insert QSL")})
						return
					}
			}
		}


		ctx.JSON(http.StatusOK, gin.H{ "results" : "Insert success!" })


}


func Qsluploads(ctx *gin.Context) {

// Initial step : To confirm sessions
  var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unknown sesssion!  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport":  env.S_host(), })
		return
	}

	user = buffer.(*config.DummyUserModel)
//	println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)
	if user.Username == "" {
	}

	// Initial step : End To confirm sessions

	ctx.HTML(http.StatusOK, "uploads.html", gin.H{
		"isLoggedIn": true,
		"username": user.Username,
		"email": user.Email,
		"domainport":  env.S_host() ,
	})

//	ctx.Redirect(http.StatusSeeOther, "/")

}


func Fileselectp(ctx *gin.Context) {

// Initial step : To confirm sessions
  var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unknown sesssion!  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport":  env.S_host(), })
		return
	}

	user = buffer.(*config.DummyUserModel)
//	println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)
	if user.Username == "" {
	}

	// Initial step : End To confirm sessions
	// Retrieval step : To get file name list and convert them to JSON


	df := config.DummyF()
			sl_json_cal, err := df.UploadExists()
			if err != nil {
				println("Error: " + err.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at File dir.")})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{ "results" : sl_json_cal })


}


func Fileuploadp(ctx *gin.Context) {

// Initial step : To confirm sessions
  var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unknown sesssion!  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport":  env.S_host(), })
		return
	}

	user = buffer.(*config.DummyUserModel)
//	println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)
	if user.Username == "" {
	}

	// Initial step : End To confirm sessions
	// uploaded file handling

		var targetPath string
		file, err := ctx.FormFile("qslupload")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{ "results" : fmt.Sprintf("File Transfer Error: %s", err.Error()) })
			log.Printf("File Transfer Error: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		targetPath = up_path + filename

		// ファイル存在チェック　IsExist check to filename.
		for n:=0 ; n < 10 ;n++ {
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				break    //  ループを抜ける
			}else{
				s_n := strconv.Itoa(n)
				targetPath = targetPath + ".bk-" + string(s_n)
				log.Printf("targetPath : %s\n　", targetPath)
			}
		}

		// Save an uploaded file to designated directory.
		if err = ctx.SaveUploadedFile(file, targetPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{ "results" :  fmt.Sprintf("File Saving Error: %s", err.Error()) } )
			log.Printf("File Saving Error: %s", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{ "results" : "File upload success!" })


}
