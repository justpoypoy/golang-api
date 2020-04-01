module checkversi.go

go 1.14

replace github.com/justpoypoy/api/lib => ./lib

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/justpoypoy/api/lib v0.0.0-00010101000000-000000000000
)
