package gapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
)

func (server *Server) UploadRoomImageHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.FormValue("room_id")
	fmt.Println("roomID", roomID)
	
	idInt, err := strconv.Atoi(roomID)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	int32RoomID := int32(idInt)

	// parsiraj multipart formu (ograničenje 10MB)
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Cannot read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadPath := fmt.Sprintf("./uploads/rooms/%s", roomID)
	os.MkdirAll(uploadPath, os.ModePerm)

	dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
	if err != nil {
		http.Error(w, "Cannot save image", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Cannot save image", http.StatusInternalServerError)
		return
	}

	folderPath := fmt.Sprintf("/uploads/rooms/%s/%s", roomID, handler.Filename)
	folderPG := pgtype.Text{String: folderPath, Valid: true}

	arg := db.UpdateRoomParams{
		ID:                 int32RoomID,
		RoomPicturesFolder: folderPG,
	}
	// Ovde ažuriraj room_pictures_folder u bazi
	room, err := server.store.UpdateRoom(r.Context(), arg)
	if err != nil {
		http.Error(w, "Cannot update DB", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, "Cannot encode response", http.StatusInternalServerError)
		return
	}
}
