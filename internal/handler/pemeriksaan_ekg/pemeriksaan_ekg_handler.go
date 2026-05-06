package pemeriksaan_ekg

import (
	"net/http"
	"strconv"

	usecase "rme/internal/usecase/pemeriksaan_ekg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// GetAll GET /pemeriksaan_ekg
func (h *Handler) GetAll(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	idDokterStr := c.Query("idDokter")
	var idPasienPtr *int
	var idDokterPtr *int
	if idPasienStr != "" {
		if v, err := strconv.Atoi(idPasienStr); err == nil {
			idPasienPtr = &v
		}
	}
	if idDokterStr != "" {
		if v, err := strconv.Atoi(idDokterStr); err == nil {
			idDokterPtr = &v
		}
	}

	rows, err := h.uc.GetAll(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]Response, 0, len(rows))
	for _, r := range rows {
		res = append(res, Response{
			ID:                 r.ID,
			IDPasien:           r.IDPasien,
			NamaPasien:         r.NamaPasien,
			IDDokter:           r.IDDokter,
			NamaDokter:         r.NamaDokter,
			Visible:            r.Visible,
			DetakJantung:       r.DetakJantung,
			Irama:              r.Irama,
			PRInterval:         r.PRInterval,
			QRSDuration:        r.QRSDuration,
			QTTcInterval:       r.QTTcInterval,
			AxisJantung:        r.AxisJantung,
			STElevationDepress: r.STElevationDepress,
			TWaveAbnormality:   r.TWaveAbnormality,
			InterpretasiDokter: r.InterpretasiDokter,
			DateMake:           r.DateMake,
		})
	}
	c.JSON(http.StatusOK, res)
}
