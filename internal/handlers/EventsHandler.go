package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RaihanurRahman2022/simple-web-server/internal/helper"
	"github.com/RaihanurRahman2022/simple-web-server/internal/models"
)

func (h *Handler) EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// all methods also needs to make the method of
	// handler structure to access db
	case http.MethodGet:
		h.eventListHandler(w, r)
	case http.MethodPost:
		h.createEventHandler(w, r)
	}
}

func (h *Handler) eventListHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(r.Context(), "SELECT id, title, description, location, start_time, end_time, organizer, attendees, created_at, updated_at, deleted_at FROM events")

	// Create a helper function to centralize the error so that we can use it in any handler
	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var events []models.Event
	// Next prepares the next row for reading. It returns true if there is another row and false if no more rows are available or a fatal error has occurred. It automatically closes rows when all rows are read.

	// should check rows.Err() after rows.Next() returns false to detect whether result-set reading ended prematurely due to an error. See Conn.Query for details.

	for rows.Next() {
		var event models.Event

		// Scan reads the values from the current row into dest values positionally. dest can include pointers to core types, values implementing the Scanner interface, and nil. nil will skip the value entirely. It is an error to call Scan without first calling Next() and checking that it returned true.

		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.Organizer, &event.Attendees, &event.CreatedAt, &event.UpdatedAt, &event.DeletedAt); err != nil {

			helper.HandleError(w, err, http.StatusInternalServerError)
			return
		}
	}

	helper.SetHeader(w)

	if err := json.NewEncoder(w).Encode(events); err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		helper.HandleError(w, err, http.StatusBadRequest)
		return
	}

	var id int
	err := h.DB.QueryRow(r.Context(), "INSERT INTO events (title, description, location, start_time, end_time, organizer, attendees) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		event.Title, event.Description, event.Location, event.StartTime, event.EndTime, event.Organizer, event.Attendees).Scan(&id)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	// Updating th ID in event Object with the inserted id
	event.ID = id
	helper.SetHeader(w)

	if err := json.NewEncoder(w).Encode(event); err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}
