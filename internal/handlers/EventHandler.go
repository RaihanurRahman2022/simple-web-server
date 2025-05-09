package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RaihanurRahman2022/simple-web-server/internal/helper"
	"github.com/RaihanurRahman2022/simple-web-server/internal/models"
)

func (h *Handler) EventHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/events/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		helper.HandleError(w, err, http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		h.getEventDetails(w, id)
	case http.MethodPut:
		h.UpdateEventDetails(w, r, id)
	case http.MethodDelete:
		h.DeleteEvent(w, r, id)
	}
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request, id int) {
	res, err := h.DB.Exec(r.Context(), `Delete from events where id=?`, id)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
	}

	count := res.RowsAffected()
	if count == 0 {

		helper.HandleError(
			w,
			errors.New("no events found with this ID"),
			http.StatusNotFound,
		)
		return
	}

	fmt.Fprintf(w, "Event is deleted successfully!!!")
}

func (h *Handler) UpdateEventDetails(w http.ResponseWriter, r *http.Request, id int) {
	event, err := h.getEventFromDB(id)

	if err != nil {
		helper.HandleError(w, err, http.StatusNotFound)
		return
	}

	var updatedEvent models.Event

	if err := json.NewDecoder(r.Body).Decode(&updatedEvent); err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	updatedEvent.ID = id
	updatedEvent.CreatedAt = event.CreatedAt
	updatedEvent.UpdatedAt = time.Now()

	_, err = h.DB.Exec(r.Context(), `UPDATE events SET title=?, description=?, location=?, start_time=?, end_time=?, updated_at=? where id=?`, updatedEvent.Title, updatedEvent.Description, updatedEvent.Location, updatedEvent.StartTime, updatedEvent.EndTime, updatedEvent.UpdatedAt, id)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	resp := helper.Response{
		Message: "Event updated successfully",
		Data:    updatedEvent,
	}
	helper.SetHeader(w)
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) getEventDetails(w http.ResponseWriter, id int) {
	event, err := h.getEventFromDB(id)

	if err != nil {
		helper.HandleError(w, err, http.StatusNotFound)
		return
	}
	helper.SetHeader(w)
	json.NewEncoder(w).Encode(event)
}

func (h *Handler) getEventFromDB(id int) (*models.Event, error) {
	var event models.Event

	err := h.DB.QueryRow(context.Background(), `Select * from events where id = ?`, id).Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.Organizer, &event.Attendees, &event.CreatedAt, &event.UpdatedAt, &event.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
