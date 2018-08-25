package main

import (
	"github.com/fiscaluno/saga/server"
	"github.com/fiscaluno/saga/student"
)

func main() {
	student.Migrate()
	server.Listen()
}
