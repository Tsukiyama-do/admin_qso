package routes

import (
	"../config"
	"../sessions"

	"net/http"
	"github.com/gin-gonic/gin"
	"log"
//	"strconv"
  "github.com/bitly/go-simplejson"
	"fmt"

)

func Qslmain(ctx *gin.Context) {

// Initial step : To confirm sessions
  var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unknown sesssion!  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": c_host, })
		return
	}

	user = buffer.(*config.DummyUserModel)
	println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

	// Initial step : End To confirm sessions

	ctx.HTML(http.StatusOK, "qslmain.html", gin.H{
		"isLoggedIn": true,
		"username": user.Username,
		"email": user.Email,
		"domainport": c_host,
	})

//	ctx.Redirect(http.StatusSeeOther, "/")

}

func Qslselectp(ctx *gin.Context) {

	println("qsl/qslselectp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": c_host, })
			return
		}

		user = buffer.(*config.DummyUserModel)
		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

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

type qsl2_rec struct{
	I     bool    `json:"I"`
	U     bool    `json:"U"`
	D     bool    `json:"D"`
	ID     int  `json:"ID"`
	CALLSIGN string `json:"CALLSIGN"`
	DATETIME string `json:"DATETIME"`
	FILES string `json:"FILES"`
}


func Qslupddelp(ctx *gin.Context) {

	println("qsl/qslupddelp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": c_host, })
			return
		}

		user = buffer.(*config.DummyUserModel)
		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

		// Initial step : End To confirm sessions
		// Parsing step : To get records and convert them to JSON

		buf := make([]byte, 2048)
		num, _ := ctx.Request.Body.Read(buf)
		log.Printf("request - num of buf : %d\n", num)

		bjson,err := simplejson.NewJson(buf)    // generate json object by simplejson
		if err != nil {
			log.Println("Error occured at JSON conversion.")
			ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at JSON conversion")})
			return
		}

    // map json text to records
//    for idx, p := range bjson.MustMap() {
//        log.Printf("Result (%d) : %v, %v, %v, %d, %s, %s, %s\n", idx, p.Get("I").Bool(), p.Get("U").Bool(), p.Get("D").Bool(), p.Get("ID").Int(), p.Get("CALLSIGN").String(), p.Get("DATETIME").String(), p.Get("FILES").String())
//    }

		// Run sql proc

//		db := config.DummyDB()
    for idx, p := range bjson.MustMap() {
				log.Printf("json( %d ) is %v \n", idx,  p)

/*				if p.U == true {
						err := db.UpdateQSL(p.Get("ID").Int(), p.Get("CALLSIGN").String(), p.Get("DATETIME").String(), p.Get("FILES").String())
						if err != nil { log.Println("Error is " + err.Error())}

				}
				if p.D == true {
						err := db.DeleteQSL(p.ID)

				}
*/
    }

		ctx.HTML(http.StatusOK, "qslmain.html", gin.H{
			"isLoggedIn": true,
			"username": user.Username,
			"email": user.Email,
			"domainport": c_host,
			"sql_result": "SQL success!",
		})


}


func Qslinsertp(ctx *gin.Context) {

	println("qsl/qslinsertp")

	// Initial step : To confirm sessions
	  var user *config.DummyUserModel

		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unknown sesssion!  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": c_host, })
			return
		}

		user = buffer.(*config.DummyUserModel)
		println("Session/User is OK, seID: " + session.ID + " , username: " + user.Username)

		// Initial step : End To confirm sessions
		// Parsing step : To get records and convert them to JSON

		buf := make([]byte, 2048)
		num, _ := ctx.Request.Body.Read(buf)
		log.Printf("request - num of buf : %d\n", num)

		bjson,err := simplejson.NewJson(buf)    // generate json object by simplejson
		if err != nil {
			log.Println("Error occured at JSON conversion.")
			ctx.JSON(http.StatusBadRequest, gin.H{ "results" :  fmt.Sprintf("Error occured at JSON conversion")})
			return
		}

    // map json text to records
//    for idx, p := range bjson.MustMap() {
//        log.Printf("Result (%d) : %v, %v, %v, %d, %s, %s, %s\n", idx, p.Get("I").Bool(), p.Get("U").Bool(), p.Get("D").Bool(), p.Get("ID").Int(), p.Get("CALLSIGN").String(), p.Get("DATETIME").String(), p.Get("FILES").String())
//    }

		// Run sql proc

//		db := config.DummyDB()
    for idx, p := range bjson.MustMap() {
				log.Printf("json( %d ) is %v \n", idx,  p)

/*				if p.U == true {
						err := db.UpdateQSL(p.Get("ID").Int(), p.Get("CALLSIGN").String(), p.Get("DATETIME").String(), p.Get("FILES").String())
						if err != nil { log.Println("Error is " + err.Error())}

				}
				if p.D == true {
						err := db.DeleteQSL(p.ID)

				}
*/
    }

		ctx.HTML(http.StatusOK, "qslmain.html", gin.H{
			"isLoggedIn": true,
			"username": user.Username,
			"email": user.Email,
			"domainport": c_host,
			"sql_result": "SQL success!",
		})


}
