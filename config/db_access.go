package config

import (
	"../crypto"
	"errors"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


// ユーザの存在チェック
func (du *DummyUserModel) CheckUser() error {

  // データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	//データの検索
	 rows, err := db.Query("SELECT * FROM USER_MODEL WHERE USERNAME=?", du.Username)
	 if err != nil { return err }   // エラー時は、エラーを返却して抜ける
	 defer rows.Close()

	 if rows.Next() {
	 } else {
				return errors.New("No_Data_Found")
	 }

	for rows.Next() {
	      var id int
	      var s_username, s_password, s_email, s_role, s_remark string
	      err = rows.Scan(&id, &s_username, &s_password, &s_email, &s_role, &s_remark)
				if err != nil { return err }   // エラー時は、エラーを返却して抜ける
	}

  return nil

}

// ユーザの認証チェック
func (du *DummyUserModel) UserAuth() error {

	var s_username, s_password, s_email string

  // データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  defer db.Close()
  if err != nil { return err }   // エラー時は、エラーを返却して抜ける

	//データの検索
	 rows, err := db.Query("SELECT USERNAME, PASSWORD, EMAIL FROM USER_MODEL WHERE USERNAME=?", du.Username)
	 if err != nil { return err }   // エラー時は、エラーを返却して抜ける

	 defer rows.Close()

	 if rows.Next() {
		 	      err = rows.Scan(&s_username, &s_password, &s_email)
		 				log.Printf(" email is %s \n", s_email)
		 				if err != nil { return err }   // エラー時は、エラーを返却して抜ける
		 				if err = crypto.CompareHashAndPassword(s_password, du.Password); err != nil {
		 					return errors.New("Password of user \"" + du.Username + "\" doesn't exists")
		 				} else {
		 					du.Email = s_email
		 				}
	 } else {
				return errors.New("User \"" + du.Username + "\" doesn't exists")
	 }

//	log.Printf("Dummy user %v \n", du)

  return nil

}

// ユーザをDBに登録する。
func (du *DummyUserModel) InsertUpload() error {

  // データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

  // データの挿入
  res, err := db.Exec(
    `INSERT INTO USER_MODEL (USERNAME, PASSWORD, EMAIL) VALUES (?, ?, ?)`,
    du.Username, du.Password, du.Email)
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  // 挿入処理の結果からIDを取得
  id, err := res.LastInsertId()
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  log.Printf("id is %v \n", id)

  return nil
}

// ユーザ情報を変更する。
func (du *DummyUserModel) UpdateUpload() error {

  // データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける
		defer db.Close()

  // データの挿入
  res, err := db.Exec(
    `UPDATE USER_MODEL SET EMAIL=? WHERE USERNAME=?`,
    du.Email, du.Username)
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  // 更新処理の結果から、その件数を取得
  no, err := res.RowsAffected()
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  log.Printf("number of records is %v \n", no)

  return nil

}

// ユーザ情報を削除する。
func (du *DummyUserModel) DeleteUpload() error {

  // データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける
		defer db.Close()

  // データの挿入
  res, err := db.Exec(
    `DELETE FROM USER_MODEL WHERE USERNAME=?`,
     du.Username)
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  // 更新処理の結果から、その件数を取得
  no, err := res.RowsAffected()
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  log.Printf("number of records is %v \n", no)

  return nil

}

//////////////////////////////////////
//  QSL CARDS MODULE
//////////////////////////////////////


type QslCardsModel struct {
	ID int
	CALLSIGN string
	DATETIME string
	FILES string
	iud string
}

// slice of QSLCARDS
type S_QSLCARDS []QslCardsModel


func (qsl *QslCardsModel) QSLRecords() (S_QSLCARDS, error) {

	var sq S_QSLCARDS

	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return nil,  err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	//データの検索
	 rows, err := db.Query("SELECT * FROM QSLCARDS ORDER BY ID")
	 if err != nil { return nil,  err }   // エラー時は、エラーを返却して抜ける
	 defer rows.Close()

//	 if rows.Next() {
//	 } else {
//				return  nil, errors.New("No_Data_Found")
//	 }


	for rows.Next() {
	      var sz QslCardsModel
	      err = rows.Scan(&sz.ID, &sz.CALLSIGN, &sz.DATETIME, &sz.FILES)
				if err != nil { return nil, err }   // エラー時は、エラーを返却して抜ける
//				log.Printf("ID : %d, CALLSIGN : %s, DATETIME : %s, FILEs: %s \n", sz.ID, sz.CALLSIGN, sz.DATETIME, sz.FILES)
				sq = append(sq, sz)
	}

	if len(sq) == 0 {
						return  nil, errors.New("No_Data_Found")
	}

  return  sq, nil

}


func (qsl *QslCardsModel) QSLUpdate() error {

	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	// SQL UPDATE proc
  res, err := db.Exec(
    `UPDATE QSLCARDS SET CALLSIGN=? , DATETIME=?, FILES=? WHERE ID=?`,
     qsl.CALLSIGN,qsl.DATETIME,qsl.FILES,qsl.ID )
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける

  // 削除されたレコード数
  affect, err := res.RowsAffected()
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	log.Printf("affected by QSLUpdate : %d\n", affect)

	return nil
}

func (qsl *QslCardsModel) QSLDelete() error {

	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	// SQL DELETE proc
  res, err := db.Exec(
    `DELETE FROM QSLCARDS WHERE ID=?`,
    qsl.ID )
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける

  // 削除されたレコード数
  affect, err := res.RowsAffected()
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	log.Printf("affected by QSLDelete : %d\n", affect)

	return nil
}


func (qsl *QslCardsModel) QSLCheck() error {

	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	// SQL Select proc
	var cnt int = 0

	rows, err := db.Query("SELECT COUNT(*) as cnt FROM QSLCARDS WHERE ID=?", qsl.ID)
		for rows.Next() {
    		err = rows.Scan(&cnt)
				if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
    }
		if cnt > 0 {
								return  errors.New("Data_Found!")
		}
	return nil

}

func (qsl *QslCardsModel) QSLInsert() error {

	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "../qso/qsldb/qsldb_main.db")
  if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	defer db.Close()

	// SQL UPDATE proc
  res, err := db.Exec(
    `INSERT INTO QSLCARDS VALUES (?, ?, ?, ?)`,
     qsl.ID, qsl.CALLSIGN,qsl.DATETIME,qsl.FILES )
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける

  // 削除されたレコード数
  affect, err := res.RowsAffected()
	if err != nil { return  err }   // エラー時は、エラーを返却して抜ける
	log.Printf("affected by QSLInsert : %d\n", affect)

	return nil

}
