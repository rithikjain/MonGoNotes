package handler

import (
	"encoding/json"
	"github.com/rithikjain/MongoNotes/api/view"
	"github.com/rithikjain/MongoNotes/pkg/entities"
	"github.com/rithikjain/MongoNotes/pkg/note"
	"net/http"
)

func createNote(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		note := &entities.Note{}
		err := json.NewDecoder(r.Body).Decode(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		n, err := svc.CreateNote(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note create",
			"status":  http.StatusOK,
			"note":    n,
		})
	})
}

func MakeNoteHandler(r *http.ServeMux, svc note.Service) {
	r.Handle("/api/notes/create", createNote(svc))
}
