package anamnesis

// HTTP handlers for anamnesis endpoints.

import (
	"net/http"
	"sort"

	usecase "rme/internal/usecase/anamnesis"

	"github.com/gin-gonic/gin"
)

// Handler membungkus usecase dan menyediakan handler HTTP.
type Handler struct {
	uc usecase.Usecase
}

// NewHandler membuat instance Handler baru.
func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// GetAll menangani GET /anamnesis (Gin handler)
func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	listB, err := h.uc.GetAllB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rsA := make([]AnamnesisResponse, 0, len(list))
	for _, data := range list {
		rsA = append(rsA, AnamnesisResponse{
			Source:                "rsA",
			ID:                    data.ID,
			NamaPasien:            data.NamaPasien,
			IDPasien:              data.IDPasien,
			IDDokter:              data.IDDokter,
			Text:                  data.Text,
			DateMake:              data.DateMake,
			DateUpdate:            data.DateUpdate,
			IDDataKlinik:          data.IDDataKlinik,
			RiwayatPengobatan:     data.RiwayatPengobatan,
			RiwayatKeluarga:       data.RiwayatKeluarga,
			RiwayatPenyakitDahulu: data.RiwayatPenyakitDahulu,
			RiwayatPenyakitLain:   data.RiwayatPenyakitLain,
			StatusKehamilan:       data.StatusKehamilan,
			KeluhanTambahan:       data.KeluhanTambahan,
			Visible:               data.Visible,
		})
	}

	sort.Slice(rsA, func(i, j int) bool {
		return rsA[i].DateMake.After(rsA[j].DateMake)
	})

	rsB := make([]AnamnesisBResponse, 0, len(listB))
	for _, data := range listB {
		rsB = append(rsB, AnamnesisBResponse{
			Source:                "rsB",
			ID:                    data.ID,
			NamaPasien:            data.NamaPasien,
			IDPasien:              data.IDPasien,
			IDDokter:              data.IDDokter,
			Text:                  data.Text,
			DateMake:              data.DateMake,
			DateUpdate:            data.DateUpdate,
			IDDataKlinik:          data.IDDataKlinik,
			RiwayatPengobatan:     data.RiwayatPengobatan,
			RiwayatKeluarga:       data.RiwayatKeluarga,
			RiwayatPenyakitDahulu: data.RiwayatPenyakitDahulu,
			RiwayatPenyakitLain:   data.RiwayatPenyakitLain,
			RiwayatAlergi:         data.RiwayatAlergi,
			StatusKehamilan:       data.StatusKehamilan,
			KeluhanUtama:          data.KeluhanUtama,
			KeluhanTambahan:       data.KeluhanTambahan,
			Visible:               data.Visible,
		})
	}

	sort.Slice(rsB, func(i, j int) bool {
		return rsB[i].DateMake.After(rsB[j].DateMake)
	})

	resp := GetAllResponse{
		RSA: rsA,
		RSB: rsB,
	}

	c.JSON(http.StatusOK, resp)
}

// ETLToWarehouse menangani POST /anamnesis/etl
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

// GetWarehouse menangani GET /anamnesis/warehouse
// Mengambil data hasil ETL dari database `data_warehouse`.
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

	rows := make([]AnamnesisWarehouseResponse, 0, len(list))
	for _, data := range list {
		rows = append(rows, AnamnesisWarehouseResponse{
			Source:                data.Source,
			IDAnamnesis:           data.IDAnamnesis,
			IDPasien:              data.IDPasien,
			NamaPasien:            data.NamaPasien,
			Text:                  data.Text,
			DateMake:              data.DateMake,
			DateUpdate:            data.DateUpdate,
			IDDataKlinik:          data.IDDataKlinik,
			RiwayatPengobatan:     data.RiwayatPengobatan,
			RiwayatKeluarga:       data.RiwayatKeluarga,
			RiwayatPenyakitDahulu: data.RiwayatPenyakitDahulu,
			RiwayatPenyakitLain:   data.RiwayatPenyakitLain,
			RiwayatAlergi:         data.RiwayatAlergi,
			StatusKehamilan:       data.StatusKehamilan,
			KeluhanUtama:          data.KeluhanUtama,
			KeluhanTambahan:       data.KeluhanTambahan,
			Visible:               data.Visible,
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].DateMake.After(rows[j].DateMake)
	})

	c.JSON(http.StatusOK, GetWarehouseResponse{Data: rows})
}
