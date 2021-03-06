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
		_, err = svc.CreateNote(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note created",
			"status":  http.StatusOK,
		})
	})
}

func viewAllNotes(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		notes, err := svc.GetAllNotes()
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Notes fetched",
			"status":  http.StatusOK,
			"notes":   notes,
		})
	})
}

func updateNote(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		note := &entities.Note{}
		err := json.NewDecoder(r.Body).Decode(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		err = svc.UpdateNote(note)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note updated",
			"status":  http.StatusOK,
		})
	})
}

func deleteNote(svc note.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		id := r.URL.Query().Get("id")
		err := svc.DeleteNote(id)
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Note Deleted",
			"status":  http.StatusNoContent,
		})
	})
}

func MakeNoteHandler(r *http.ServeMux, svc note.Service) {
	r.Handle("/api/notes/create", createNote(svc))
	r.Handle("/api/notes", viewAllNotes(svc))
	r.Handle("/api/notes/update", updateNote(svc))
	r.Handle("/api/notes/delete", deleteNote(svc))
}
