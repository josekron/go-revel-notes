package controllers

import (
	"josekron/go-revel-notes/app/models"
	"log"
	"os"
	"strings"

	"github.com/revel/revel"
)

var logger *log.Logger

type NoteController struct {
	*revel.Controller
}

var notes = []models.Note{
	models.Note{1, "Shopping list", "Milk, eggs, coffee", []string{"cooking", "product", "shopping"}},
	models.Note{2, "Password router", "12345", []string{"security", "password"}},
	models.Note{3, "John´s phone", "612345567", []string{"phone", "contact", "personal"}},
}

func init() {
	file, _ := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	logger = log.New(file, "noteController: ", log.Ldate|log.Ltime|log.Lshortfile)
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
	logger.Println("Filter tags: %s", tags)
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
	logger.Println("Check user...")
	return nil
}
