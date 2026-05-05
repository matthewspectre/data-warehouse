package anamnesis

// MySQL implementation of anamnesis repository.

import (
	"context"

	entity "rme/internal/entity/anamnesis"
	model "rme/internal/model/anamnesis"
	repo "rme/internal/repository/anamnesis"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.Anamnesis) *model.AnamnesisModel {
	if e == nil {
		return nil
	}
	return &model.AnamnesisModel{
		ID:                    e.ID,
		IDPasien:              e.IDPasien,
		IDDokter:              e.IDDokter,
		Text:                  e.Text,
		DateMake:              e.DateMake,
		DateUpdate:            e.DateUpdate,
		IDDataKlinik:          e.IDDataKlinik,
		RiwayatPengobatan:     e.RiwayatPengobatan,
		RiwayatKeluarga:       e.RiwayatKeluarga,
		RiwayatPenyakitDahulu: e.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   e.RiwayatPenyakitLain,
		StatusKehamilan:       e.StatusKehamilan,
		KeluhanTambahan:       e.KeluhanTambahan,
		Visible:               e.Visible,
	}
}

func toEntity(m *model.AnamnesisModel) *entity.Anamnesis {
	if m == nil {
		return nil
	}
	return &entity.Anamnesis{
		ID:                    m.ID,
		IDPasien:              m.IDPasien,
		IDDokter:              m.IDDokter,
		Text:                  m.Text,
		DateMake:              m.DateMake,
		DateUpdate:            m.DateUpdate,
		IDDataKlinik:          m.IDDataKlinik,
		RiwayatPengobatan:     m.RiwayatPengobatan,
		RiwayatKeluarga:       m.RiwayatKeluarga,
		RiwayatPenyakitDahulu: m.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   m.RiwayatPenyakitLain,
		StatusKehamilan:       m.StatusKehamilan,
		KeluhanTambahan:       m.KeluhanTambahan,
		Visible:               m.Visible,
	}
}

func toEntityB(m *model.AnamnesisBModel) *entity.AnamnesisB {
	if m == nil {
		return nil
	}
	return &entity.AnamnesisB{
		ID:                    m.ID,
		IDPasien:              m.IDPasien,
		IDDokter:              m.IDDokter,
		Text:                  m.Text,
		DateMake:              m.DateMake,
		DateUpdate:            m.DateUpdate,
		IDDataKlinik:          m.IDDataKlinik,
		RiwayatPengobatan:     m.RiwayatPengobatan,
		RiwayatKeluarga:       m.RiwayatKeluarga,
		RiwayatPenyakitDahulu: m.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   m.RiwayatPenyakitLain,
		RiwayatAlergi:         m.RiwayatAlergi,
		StatusKehamilan:       m.StatusKehamilan,
		KeluhanUtama:          m.KeluhanUtama,
		KeluhanTambahan:       m.KeluhanTambahan,
		Visible:               m.Visible,
	}
}

// GetAll mengambil semua data anamnesis yang masih visible.
// Jika `idDokter` atau `idPasien` tidak nil, hasil akan difilter berdasarkan kolom terkait.
func (r *RepositoryMySQL) GetAll() ([]*entity.Anamnesis, error) {
	ctx := context.Background()
	type modelWithName struct {
		model.AnamnesisModel
		NamaPasien string `gorm:"column:nama_pasien"`
	}
	var awls []modelWithName
	q := r.db.WithContext(ctx).
		Table("anamnesis a").
		Select("a.*, p.name AS nama_pasien").
		Joins("LEFT JOIN patients p ON p.id = a.id_pasien").
		Where("a.visible = ?", 1)

	if err := q.Find(&awls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Anamnesis, 0, len(awls))
	for i := range awls {
		ent := toEntity(&awls[i].AnamnesisModel)
		ent.NamaPasien = awls[i].NamaPasien
		res = append(res, ent)
	}
	return res, nil
}

// GetAllB mengambil semua data anamnesis dari database rme-b yang masih visible.
// NOTE: Nama database mengandung tanda '-' sehingga harus di-quote dengan backticks.
func (r *RepositoryMySQL) GetAllB() ([]*entity.AnamnesisB, error) {
	ctx := context.Background()
	type modelWithName struct {
		model.AnamnesisBModel
		NamaPasien string `gorm:"column:nama_pasien"`
	}
	var rows []modelWithName

	q := r.db.WithContext(ctx).
		Table("`rme-b`.anamnesis a").
		Select("a.*, p.name AS nama_pasien").
		Joins("LEFT JOIN `rme-b`.patients p ON p.id = a.id_pasien").
		Where("a.visible = ?", 1)

	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.AnamnesisB, 0, len(rows))
	for i := range rows {
		ent := toEntityB(&rows[i].AnamnesisBModel)
		ent.NamaPasien = rows[i].NamaPasien
		res = append(res, ent)
	}
	return res, nil
}
