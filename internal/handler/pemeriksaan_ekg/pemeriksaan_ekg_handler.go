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

// GetAllB GET /pemeriksaan_ekg/b
func (h *Handler) GetAllB(c *gin.Context) {
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

	rows, err := h.uc.GetAllB(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]ResponseB, 0, len(rows))
	for _, r := range rows {
		res = append(res, ResponseB{
			Source:             "rsB",
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
			PWave:              r.PWave,
			AVBlock:            r.AVBlock,
			InterpretasiDokter: r.InterpretasiDokter,
			DateMake:           r.DateMake,
			DateUpdate:         r.DateUpdate,
		})
	}
	c.JSON(http.StatusOK, res)
}

// ETLToWarehouse POST /pemeriksaan_ekg/etl
func (h *Handler) ETLToWarehouse(c *gin.Context) {
	res, err := h.uc.ETLToWarehouse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "etl completed",
		"inserted_rs_a": res.InsertedRSA,
		"inserted_rs_b": res.InsertedRSB,
	})
}

// GetWarehouse GET /pemeriksaan_ekg/warehouse
// Optional: ?NIK=...
func (h *Handler) GetWarehouse(c *gin.Context) {
	nik := c.Query("NIK")
	var nikPtr *string
	if nik != "" {
		n := nik
		nikPtr = &n
	}

	rows, err := h.uc.GetAllWarehouse(nikPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]WarehouseResponse, 0, len(rows))
	for _, r := range rows {
		res = append(res, WarehouseResponse{
			Source:             r.Source,
			ID:                 r.IDPemeriksaanEkg,
			IDPasien:           r.IDPasien,
			NamaPasien:         r.NamaPasien,
			Visible:            r.Visible,
			DetakJantung:       r.DetakJantung,
			Irama:              r.Irama,
			PRInterval:         r.PRInterval,
			QRSDuration:        r.QRSDuration,
			QTTcInterval:       r.QTTcInterval,
			AxisJantung:        r.AxisJantung,
			STElevationDepress: r.STElevationDepress,
			TWaveAbnormality:   r.TWaveAbnormality,
			PWave:              r.PWave,
			AVBlock:            r.AVBlock,
			InterpretasiDokter: r.InterpretasiDokter,
			DateMake:           r.DateMake,
			DateUpdate:         r.DateUpdate,
		})
	}

	c.JSON(http.StatusOK, GetWarehouseResponse{Data: res})
}
