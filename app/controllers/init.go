package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(NoteController.checkUser, revel.BEFORE)
}
