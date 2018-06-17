package controllers

import (
	"fmt"
	"josekron/go-revel-notes/app/models"
	"strings"

	"github.com/revel/revel"
)

type NoteController struct {
	*revel.Controller
}

var notes = []models.Note{
	models.Note{1, "Shopping list", "Milk, eggs, coffee", []string{"cooking", "product", "shopping"}},
	models.Note{2, "Password router", "12345", []string{"security", "password"}},
	models.Note{3, "John´s phone", "612345567", []string{"phone", "contact", "personal"}},
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

func (c NoteController) FilterTags(tags string) revel.Result {
	fmt.Println("FilterTags - tags: %s", tags)
	var tagList = strings.Split(tags, ",")
	var result = []models.Note{}

	if len(tagList) > 0 {
		for _, tag := range tagList {
			for _, note := range notes {
				if note.IsTagContained(tag) {
					result = append(result, note)
				}
			}
		}
	}
	return c.RenderJSON(result)
}

func (c NoteController) checkUser() revel.Result {
	fmt.Println("Check user...")
	return nil
}
