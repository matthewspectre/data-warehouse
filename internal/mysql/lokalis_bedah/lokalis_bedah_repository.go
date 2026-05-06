package lokalis_bedah

import (
	entity "rme/internal/entity/lokalis_bedah"
	model "rme/internal/model/lokalis_bedah"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	var list []model.LokalisBedahModel
	q := m.DB.Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.LokalisBedah, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.LokalisBedah{
			ID:             mm.ID,
			IDPasien:       mm.IDPasien,
			IDDokter:       mm.IDDokter,
			Tanggal:        mm.Tanggal,
			LokasiKelainan: mm.LokasiKelainan,
			JenisKelainan:  mm.JenisKelainan,
			Ukuran:         mm.Ukuran,
			Warna:          mm.Warna,
			NyeriTekan:     mm.NyeriTekan,
			Konsistensi:    mm.Konsistensi,
			Mobilitas:      mm.Mobilitas,
			TandaRadang:    mm.TandaRadang,
			Fluktuasi:      mm.Fluktuasi,
			Catatan:        mm.Catatan,
			Visible:        mm.Visible,
		})
	}
	return out, nil
}
