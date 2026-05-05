package pemeriksaanvital

import (
	"net/http"
	"strconv"

	usecase "rme/internal/usecase/pemeriksaan_vital"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// GetAll menangani GET /pemeriksaan_vital
func (h *Handler) GetAll(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	idDokterStr := c.Query("idDokter")
	var idPasienPtr *int
	var idDokterPtr *int
	if idPasienStr != "" {
		v, err := strconv.Atoi(idPasienStr)
		if err == nil {
			idPasienPtr = &v
		}
	}
	if idDokterStr != "" {
		v, err := strconv.Atoi(idDokterStr)
		if err == nil {
			idDokterPtr = &v
		}
	}
	list, err := h.uc.GetAll(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]PemeriksaanVitalResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, PemeriksaanVitalResponse{
			ID:             d.ID,
			IDPasien:       d.IDPasien,
			IDDokter:       d.IDDokter,
			DateMake:       d.DateMake,
			DateUpdate:     d.DateUpdate,
			TekananDarah:   d.TekananDarah,
			DenyutNadi:     d.DenyutNadi,
			SuhuTubuh:      d.SuhuTubuh,
			FrekuensiNapas: d.FrekuensiNapas,
			BeratBadan:     d.BeratBadan,
			TinggiBadan:    d.TinggiBadan,
			Visible:        d.Visible,
			NamaPasien:     d.NamaPasien,
			NamaDokter:     d.NamaDokter,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllB menangani GET /pemeriksaan_vital/b (data dari database rme-b)
func (h *Handler) GetAllB(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	idDokterStr := c.Query("idDokter")
	var idPasienPtr *int
	var idDokterPtr *int
	if idPasienStr != "" {
		v, err := strconv.Atoi(idPasienStr)
		if err == nil {
			idPasienPtr = &v
		}
	}
	if idDokterStr != "" {
		v, err := strconv.Atoi(idDokterStr)
		if err == nil {
			idDokterPtr = &v
		}
	}

	list, err := h.uc.GetAllB(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]PemeriksaanVitalBResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, PemeriksaanVitalBResponse{
			Source:           "rsB",
			ID:               d.ID,
			IDPasien:         d.IDPasien,
			IDDokter:         d.IDDokter,
			DateMake:         d.DateMake,
			DateUpdate:       d.DateUpdate,
			TekananDarah:     d.TekananDarah,
			DenyutNadi:       d.DenyutNadi,
			SuhuTubuh:        d.SuhuTubuh,
			FrekuensiNapas:   d.FrekuensiNapas,
			BeratBadan:       d.BeratBadan,
			TinggiBadan:      d.TinggiBadan,
			SaturasiOksigen:  d.SaturasiOksigen,
			TingkatKesadaran: d.TingkatKesadaran,
			IndeksMasaTubuh:  d.IndeksMasaTubuh,
			Visible:          d.Visible,
			NamaPasien:       d.NamaPasien,
			NamaDokter:       d.NamaDokter,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// ETLToWarehouse menangani POST /pemeriksaan_vital/etl
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

// GetWarehouse menangani GET /pemeriksaan_vital/warehouse
// Mengambil data hasil ETL dari database `data_warehouse`.
// Optional: filter berdasarkan NIK.
func (h *Handler) GetWarehouse(c *gin.Context) {
	nik := c.Query("NIK")
	var nikPtr *string
	if nik != "" {
		n := nik
		nikPtr = &n
	}

	list, err := h.uc.GetAllWarehouse(nikPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows := make([]PemeriksaanVitalWarehouseResponse, 0, len(list))
	for _, d := range list {
		rows = append(rows, PemeriksaanVitalWarehouseResponse{
			Source:             d.Source,
			IDPemeriksaanVital: d.IDPemeriksaanVital,
			IDPasien:           d.IDPasien,
			NamaPasien:         d.NamaPasien,
			DateMake:           d.DateMake,
			DateUpdate:         d.DateUpdate,
			TekananDarah:       d.TekananDarah,
			DenyutNadi:         d.DenyutNadi,
			SuhuTubuh:          d.SuhuTubuh,
			FrekuensiNapas:     d.FrekuensiNapas,
			BeratBadan:         d.BeratBadan,
			TinggiBadan:        d.TinggiBadan,
			SaturasiOksigen:    d.SaturasiOksigen,
			TingkatKesadaran:   d.TingkatKesadaran,
			IndeksMasaTubuh:    d.IndeksMasaTubuh,
			Visible:            d.Visible,
		})
	}

	c.JSON(http.StatusOK, GetWarehouseResponse{Data: rows})
}
