package main

import (
	"github.com/gin-gonic/gin"
	Lib "github.com/justpoypoy/api/lib" // Lib custom
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	grouping := route.Group("/base/v1")
	{
		// Lib.VersionInfo => memanggil function Public yg ada di module Lib
		// public function diawali dengan huruf kapital ex: VersionInfo <= bisa dipanggil
		// private function diawali dengan huruf kecil ex: versionInfo <= tidak bisa dipanggil di function selain module itu sendiri
		grouping.POST("/info/version", Lib.VersionInfo)
	}

	// set running port
	route.Run(":8082")

	// ENDPOINT menjadi = http://localhost:8082/base/v1/info/version
}
