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
