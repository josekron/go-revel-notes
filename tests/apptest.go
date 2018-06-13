package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) TestNotesList() {
	t.Get("/notes")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestNotesElementOk() {
	t.Get("/notes/1")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestNotesElementFail() {
	t.Get("/notes/50")
	t.AssertNotFound()
	t.AssertContentType("text/html; charset=utf-8")
}
