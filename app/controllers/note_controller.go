package controllers

import (
	"josekron/go-revel-notes/app/models"

	"github.com/revel/revel"
)

type NoteController struct {
	*revel.Controller
}

var notes = []models.Note{
	models.Note{1, "Shopping list", "Milk, eggs, coffee"},
	models.Note{2, "Password router", "12345"},
	models.Note{3, "John´s phone", "612345567"},
}

func (c NoteController) List() revel.Result {
	return c.RenderJSON(notes)
}

func (c NoteController) Show(noteId int) revel.Result {
	var res models.Note

	for _, note := range notes {
		if note.Id == noteId {
			res = note
		}
	}

	if res.Id == 0 {
		return c.NotFound("Could not find note")
	}

	return c.RenderJSON(res)
}
