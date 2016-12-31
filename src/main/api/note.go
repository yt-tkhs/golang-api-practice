package api

import (
	"github.com/labstack/echo"
	"net/http"
	"main/types"
	"strconv"
)

var notes map[int]types.Note

func init() {
	notes = map[int]types.Note {
		1: {
			Id: 1,
			Title: "Note1",
			Content: "Hoge",
		},
		2: {
			Id: 2,
			Title: "Note2",
			Content: "Fuga",
		},
	}
}

func GetAllNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		notesSlice := make([]types.Note, 0, len(notes))

		for _, value := range notes {
			notesSlice = append(notesSlice, value)
		}

		return c.JSON(http.StatusOK, notesSlice)
	}
}

func GetNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "{}")
	}

	return c.JSON(http.StatusOK, notes[id])
}

func CreateNote(c echo.Context) error {
	note := new(types.Note)
	if err := c.Bind(note); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, note)
}