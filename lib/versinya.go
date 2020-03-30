package lib

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbport    = "3306"
	dbaddress = "localhost"
	userdb    = "root"
	passdb    = ""
	dbdata    = "suwondo"
)

func VersionInfo(c *gin.Context) {
	// set variable untuk menampung field mysql
	var versi string
	var os string
	var versiID string

	var InfoStruct InfoVersion

	// validasi request input sesuai yang didefinisikan di InfoVersion
	if err := c.ShouldBindJSON(&InfoStruct); err != nil {
		c.JSON(400, gin.H{"Error ": err.Error()})
		return
	}

	// open koneksi ke database
	db, err := sql.Open("mysql", userdb+":"+passdb+"@tcp("+dbaddress+":"+dbport+")/"+dbdata)

	// check error jika ada
	if err != nil {
		log.Println(err.Error())
	}

	// tutup koneksi database dengan defer,
	// defer akan tereksekusi pada saat kode terakhir jalan
	// wlwpun didefine sebelum query statement
	defer db.Close()

	// statement query ke database
	query, err := db.Query("SELECT id, version, os FROM version WHERE uuid = ?", InfoStruct.Uuid)

	// check error jika ada
	if err != nil {
		log.Println(err.Error())
	}

	// memproses hasil statement query
	if query.Next() {
		// variable yang diset untuk menampung field dipakai disini
		query.Scan(&versiID, &versi, &os)
	}
	query.Close()
	// tutup statement query

	// set json format
	// InfoVersionResponse => WAJIB didefine di struct jika ingin menampilkan data yang akan dikeluarkan
	// &ResultInfoVersion => WAJIB didefine di struct jika ingin menampilkan data yang akan dikeluarkan
	responseJson := InfoVersionResponse{
		Success: true,
		Message: "sukses",
		Code:    "200",
		Result: &ResultInfoVersion{
			Version:          versi,
			Operating_system: os,
		},
	}

	// set responseJson ke format JSON
	Respjson, _ := json.Marshal(responseJson)
	// tambahkan Content-Type application/json
	c.Header("Content-Type", "application/json; charset=utf-8")
	// set status header 200 jika data sukses
	c.String(200, string(Respjson))
}
