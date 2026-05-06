package lokalis_bedah

import (
	"net/http"
	"strconv"
	"time"

	usecase "rme/internal/usecase/lokalis_bedah"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetAll(c *gin.Context) {
	var idDokterPtr *int
	var idPasienPtr *int
	idDokterStr := c.Query("idDokter")
	idPasienStr := c.Query("idPasien")
	if idDokterStr != "" {
		idd, err := strconv.Atoi(idDokterStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
			return
		}
		idDokterPtr = &idd
	}
	if idPasienStr != "" {
		idp, err := strconv.Atoi(idPasienStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
			return
		}
		idPasienPtr = &idp
	}

	list, err := h.uc.GetAll(idDokterPtr, idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]LokalisBedahResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, LokalisBedahResponse{
			ID:             d.ID,
			IDPasien:       d.IDPasien,
			IDDokter:       d.IDDokter,
			Tanggal:        d.Tanggal.Format(time.RFC3339),
			LokasiKelainan: d.LokasiKelainan,
			JenisKelainan:  d.JenisKelainan,
			Ukuran:         d.Ukuran,
			Warna:          d.Warna,
			NyeriTekan:     d.NyeriTekan,
			Konsistensi:    d.Konsistensi,
			Mobilitas:      d.Mobilitas,
			TandaRadang:    d.TandaRadang,
			Fluktuasi:      d.Fluktuasi,
			Catatan:        d.Catatan,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllB menangani GET /lokalis_bedah/b (data dari database rme-b)
func (h *Handler) GetAllB(c *gin.Context) {
	var idDokterPtr *int
	var idPasienPtr *int
	idDokterStr := c.Query("idDokter")
	idPasienStr := c.Query("idPasien")
	if idDokterStr != "" {
		idd, err := strconv.Atoi(idDokterStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
			return
		}
		idDokterPtr = &idd
	}
	if idPasienStr != "" {
		idp, err := strconv.Atoi(idPasienStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
			return
		}
		idPasienPtr = &idp
	}

	list, err := h.uc.GetAllB(idDokterPtr, idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]LokalisBedahBResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, LokalisBedahBResponse{
			Source:         "rsB",
			ID:             d.ID,
			IDPasien:       d.IDPasien,
			IDDokter:       d.IDDokter,
			Tanggal:        d.Tanggal.Format(time.RFC3339),
			LokasiKelainan: d.LokasiKelainan,
			JenisKelainan:  d.JenisKelainan,
			Ukuran:         d.Ukuran,
			Warna:          d.Warna,
			NyeriTekan:     d.NyeriTekan,
			Konsistensi:    d.Konsistensi,
			Mobilitas:      d.Mobilitas,
			TandaRadang:    d.TandaRadang,
			Fluktuasi:      d.Fluktuasi,
			Catatan:        d.Catatan,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// ETLToWarehouse menangani POST /lokalis_bedah/etl
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

// GetWarehouse menangani GET /lokalis_bedah/warehouse
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

	rows := make([]LokalisBedahWarehouseResponse, 0, len(list))
	for _, d := range list {
		rows = append(rows, LokalisBedahWarehouseResponse{
			Source:         d.Source,
			NamaPasien:     d.NamaPasien,
			ID:             d.ID,
			IDPasien:       d.IDPasien,
			IDDokter:       d.IDDokter,
			Tanggal:        d.Tanggal.Format(time.RFC3339),
			LokasiKelainan: d.LokasiKelainan,
			JenisKelainan:  d.JenisKelainan,
			Ukuran:         d.Ukuran,
			Warna:          d.Warna,
			NyeriTekan:     d.NyeriTekan,
			Konsistensi:    d.Konsistensi,
			Mobilitas:      d.Mobilitas,
			TandaRadang:    d.TandaRadang,
			Fluktuasi:      d.Fluktuasi,
			Catatan:        d.Catatan,
		})
	}

	c.JSON(http.StatusOK, GetWarehouseResponse{Data: rows})
}
