package api

import (
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var (
	DB   *gorm.DB
	REST *httprouter.Router
	Box  packr.Box
	pass string
)

func Up(rest *httprouter.Router, db *gorm.DB, pck packr.Box, p string) {
	DB = db
	REST = rest
	Box = pck
	pass = p
	CreateTableDB()
	Createcallhandle()
}
