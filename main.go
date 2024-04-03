package main

import (
	"cooking-receipt/connector/sqliteConnector"
	"cooking-receipt/route"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteConnector.GetInstance()
	route.HandleRequest()
}
