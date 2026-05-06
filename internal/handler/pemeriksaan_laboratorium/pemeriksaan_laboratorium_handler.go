package pemeriksaan_laboratorium

import (
	"net/http"
	usecase "rme/internal/usecase/pemeriksaan_laboratorium"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

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
			ID:           r.ID,
			IDPasien:     r.IDPasien,
			Visible:      r.Visible,
			NamaPasien:   r.NamaPasien,
			IDDokter:     r.IDDokter,
			NamaDokter:   r.NamaDokter,
			Hb:           r.Hb,
			Ht:           r.Ht,
			Leukosit:     r.Leukosit,
			Trombosit:    r.Trombosit,
			GulaPuasa:    r.GulaPuasa,
			GulaSewaktu:  r.GulaSewaktu,
			HbA1c:        r.HbA1c,
			Kolesterol:   r.Kolesterol,
			HDL:          r.HDL,
			LDL:          r.LDL,
			Trigliserida: r.Trigliserida,
			SGOT:         r.SGOT,
			SGPT:         r.SGPT,
			Ureum:        r.Ureum,
			Kreatinin:    r.Kreatinin,
			AsamUrat:     r.AsamUrat,
			Natrium:      r.Natrium,
			Kalium:       r.Kalium,
			Klorida:      r.Klorida,
			Waktu:        r.Waktu,
		})
	}
	c.JSON(http.StatusOK, res)
}

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
			Source:         "rsB",
			ID:             r.ID,
			IDPasien:       r.IDPasien,
			Visible:        r.Visible,
			NamaPasien:     r.NamaPasien,
			IDDokter:       r.IDDokter,
			NamaDokter:     r.NamaDokter,
			Hb:             r.Hb,
			Ht:             r.Ht,
			Leukosit:       r.Leukosit,
			Trombosit:      r.Trombosit,
			GulaPuasa:      r.GulaPuasa,
			GulaSewaktu:    r.GulaSewaktu,
			HbA1c:          r.HbA1c,
			Kolesterol:     r.Kolesterol,
			HDL:            r.HDL,
			LDL:            r.LDL,
			Trigliserida:   r.Trigliserida,
			SGOT:           r.SGOT,
			SGPT:           r.SGPT,
			Ureum:          r.Ureum,
			Kreatinin:      r.Kreatinin,
			AsamUrat:       r.AsamUrat,
			Natrium:        r.Natrium,
			Kalium:         r.Kalium,
			Klorida:        r.Klorida,
			Waktu:          r.Waktu,
			LajuEndapDarah: r.LajuEndapDarah,
			Albumin:        r.Albumin,
		})
	}
	c.JSON(http.StatusOK, res)
}

// ETLToWarehouse menangani POST /pemeriksaan_laboratorium/etl
// Menarik data dari rsA dan rsB lalu menyimpan ke database `data_warehouse`.
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

// GetWarehouse menangani GET /pemeriksaan_laboratorium/warehouse
// Mengambil data hasil ETL dari database `data_warehouse`.
// Optional: filter berdasarkan NIK.
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
			Source:         r.Source,
			ID:             r.IDPemeriksaanLaboratorium,
			IDPasien:       r.IDPasien,
			NamaPasien:     r.NamaPasien,
			Hb:             r.Hb,
			Ht:             r.Ht,
			Leukosit:       r.Leukosit,
			Trombosit:      r.Trombosit,
			GulaPuasa:      r.GulaPuasa,
			GulaSewaktu:    r.GulaSewaktu,
			HbA1c:          r.HbA1c,
			Kolesterol:     r.Kolesterol,
			HDL:            r.HDL,
			LDL:            r.LDL,
			Trigliserida:   r.Trigliserida,
			SGOT:           r.SGOT,
			SGPT:           r.SGPT,
			Ureum:          r.Ureum,
			Kreatinin:      r.Kreatinin,
			AsamUrat:       r.AsamUrat,
			Natrium:        r.Natrium,
			Kalium:         r.Kalium,
			Klorida:        r.Klorida,
			Waktu:          r.Waktu,
			LajuEndapDarah: r.LajuEndapDarah,
			Albumin:        r.Albumin,
			Visible:        r.Visible,
		})
	}

	c.JSON(http.StatusOK, GetWarehouseResponse{Data: res})
}
