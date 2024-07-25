package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
)

type createRoomRequest struct {
	RoomNameSr         string `json:"room_name_sr" binding:"required,alphanum"`
	RoomNameEn         string `json:"room_name_en" binding:"required,alphanum"`
	RoomNameBg         string `json:"room_name_bg" binding:"required,alphanum"`
	RoomShortDesSr     string `json:"room_short_des_sr" binding:"required,alphanum"`
	RoomShortDesEn     string `json:"room_short_des_en" binding:"required,alphanum"`
	RoomShortDesBg     string `json:"room_short_des_bg" binding:"required,alphanum"`
	RoomDescriptionSr  string `json:"room_description_sr" binding:"required,alphanum"`
	RoomDescriptionEn  string `json:"room_description_en" binding:"required,alphanum"`
	RoomDescriptionBg  string `json:"room_description_bg" binding:"required,alphanum"`
	RoomPicturesFolder string `json:"room_pictures_folder" binding:"required,alphanum"`
	RoomGuestNumber    int32  `json:"room_guest_number" binding:"required"`
	RoomPriceEn        int32  `json:"room_price_en" binding:"required"`
}

func (server *Server) createRoom(ctx *gin.Context) {
	var req createRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRoomParams{
		RoomNameSr:         req.RoomNameSr,
		RoomNameEn:         req.RoomNameEn,
		RoomNameBg:         req.RoomNameBg,
		RoomShortDesSr:     req.RoomShortDesSr,
		RoomShortDesEn:     req.RoomShortDesEn,
		RoomShortDesBg:     req.RoomShortDesBg,
		RoomDescriptionSr:  req.RoomDescriptionSr,
		RoomDescriptionEn:  req.RoomDescriptionEn,
		RoomDescriptionBg:  req.RoomDescriptionBg,
		RoomPicturesFolder: req.RoomPicturesFolder,
		RoomGuestNumber:    req.RoomGuestNumber,
		RoomPriceEn:        req.RoomPriceEn,
	}

	room, err := server.store.CreateRoom(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, room)
}

type getRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRoom(ctx *gin.Context) {
	var req getRoomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	room, err := server.store.GetRoom(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, room)
}

type listRoomsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listRooms(ctx *gin.Context) {
	var req listRoomsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRoomsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	rooms, err := server.store.ListRooms(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}
