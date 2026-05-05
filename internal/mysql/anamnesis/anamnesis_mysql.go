package anamnesis

// MySQL implementation of anamnesis repository.

import (
	"context"
	"errors"

	entity "rme/internal/entity/anamnesis"
	model "rme/internal/model/anamnesis"
	repo "rme/internal/repository/anamnesis"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func toEntityWarehouse(m *model.AnamnesisWarehouseModel) *entity.AnamnesisWarehouse {
	if m == nil {
		return nil
	}
	return &entity.AnamnesisWarehouse{
		Source:                m.Source,
		IDAnamnesis:           m.IDAnamnesis,
		IDPasien:              m.IDPasien,
		NamaPasien:            m.NamaPasien,
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

// GetAllWarehouse mengambil semua data hasil ETL dari `data_warehouse`.anamnesis.
// Default: hanya yang visible=1.
func (r *RepositoryMySQL) GetAllWarehouse(nik *string) ([]*entity.AnamnesisWarehouse, error) {
	ctx := context.Background()
	var rows []model.AnamnesisWarehouseModel

	var idPasienRSA *int
	var idPasienRSB *int
	if nik != nil && *nik != "" {
		type patientIDRow struct {
			ID int `gorm:"column:id"`
		}

		var rowA patientIDRow
		txA := r.db.WithContext(ctx).
			Table("patients").
			Select("id").
			Where("nik = ?", *nik).
			Take(&rowA)
		if txA.Error != nil && !errors.Is(txA.Error, gorm.ErrRecordNotFound) {
			return nil, txA.Error
		}
		if txA.Error == nil {
			v := rowA.ID
			idPasienRSA = &v
		}

		var rowB patientIDRow
		txB := r.db.WithContext(ctx).
			Table("`rme-b`.patients").
			Select("id").
			Where("nik = ?", *nik).
			Take(&rowB)
		if txB.Error != nil && !errors.Is(txB.Error, gorm.ErrRecordNotFound) {
			return nil, txB.Error
		}
		if txB.Error == nil {
			v := rowB.ID
			idPasienRSB = &v
		}

		if idPasienRSA == nil && idPasienRSB == nil {
			return []*entity.AnamnesisWarehouse{}, nil
		}
	}

	q := r.db.WithContext(ctx).
		Model(&model.AnamnesisWarehouseModel{}).
		Where("visible = ?", 1).
		Order("date_make DESC")

	if idPasienRSA != nil || idPasienRSB != nil {
		var cond *gorm.DB
		if idPasienRSA != nil {
			cond = r.db.Where("source = ? AND id_pasien = ?", "rsA", *idPasienRSA)
		}
		if idPasienRSB != nil {
			if cond == nil {
				cond = r.db.Where("source = ? AND id_pasien = ?", "rsB", *idPasienRSB)
			} else {
				cond = cond.Or("source = ? AND id_pasien = ?", "rsB", *idPasienRSB)
			}
		}
		q = q.Where(cond)
	}

	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.AnamnesisWarehouse, 0, len(rows))
	for i := range rows {
		ent := toEntityWarehouse(&rows[i])
		if ent != nil {
			res = append(res, ent)
		}
	}
	return res, nil
}

func toWarehouseModelA(e *entity.Anamnesis) *model.AnamnesisWarehouseModel {
	if e == nil {
		return nil
	}
	var nama *string
	if e.NamaPasien != "" {
		n := e.NamaPasien
		nama = &n
	}
	return &model.AnamnesisWarehouseModel{
		Source:                "rsA",
		IDAnamnesis:           e.ID,
		IDPasien:              e.IDPasien,
		NamaPasien:            nama,
		Text:                  e.Text,
		DateMake:              e.DateMake,
		DateUpdate:            e.DateUpdate,
		IDDataKlinik:          e.IDDataKlinik,
		RiwayatPengobatan:     e.RiwayatPengobatan,
		RiwayatKeluarga:       e.RiwayatKeluarga,
		RiwayatPenyakitDahulu: e.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   e.RiwayatPenyakitLain,
		RiwayatAlergi:         nil,
		StatusKehamilan:       e.StatusKehamilan,
		KeluhanUtama:          nil,
		KeluhanTambahan:       e.KeluhanTambahan,
		Visible:               e.Visible,
	}
}

func toWarehouseModelB(e *entity.AnamnesisB) *model.AnamnesisWarehouseModel {
	if e == nil {
		return nil
	}
	var nama *string
	if e.NamaPasien != "" {
		n := e.NamaPasien
		nama = &n
	}
	var alergi *string
	if e.RiwayatAlergi != "" {
		a := e.RiwayatAlergi
		alergi = &a
	}
	var utama *string
	if e.KeluhanUtama != "" {
		u := e.KeluhanUtama
		utama = &u
	}
	return &model.AnamnesisWarehouseModel{
		Source:                "rsB",
		IDAnamnesis:           e.ID,
		IDPasien:              e.IDPasien,
		NamaPasien:            nama,
		Text:                  e.Text,
		DateMake:              e.DateMake,
		DateUpdate:            e.DateUpdate,
		IDDataKlinik:          e.IDDataKlinik,
		RiwayatPengobatan:     e.RiwayatPengobatan,
		RiwayatKeluarga:       e.RiwayatKeluarga,
		RiwayatPenyakitDahulu: e.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   e.RiwayatPenyakitLain,
		RiwayatAlergi:         alergi,
		StatusKehamilan:       e.StatusKehamilan,
		KeluhanUtama:          utama,
		KeluhanTambahan:       e.KeluhanTambahan,
		Visible:               e.Visible,
	}
}

func (r *RepositoryMySQL) upsertWarehouse(rows []*model.AnamnesisWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	ctx := context.Background()

	// Idempotent-ish behavior depends on a UNIQUE KEY on (source, id_anamnesis)
	// in `data_warehouse`.anamnesis.
	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_anamnesis"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"text",
			"date_make",
			"date_update",
			"id_data_klinik",
			"riwayat_pengobatan",
			"riwayat_keluarga",
			"riwayat_penyakit_dahulu",
			"riwayat_penyakit_lain",
			"riwayat_alergi",
			"status_kehamilan",
			"keluhan_utama",
			"keluhan_tambahan",
			"visible",
		}),
	}

	// Batch insert for ETL.
	err := r.db.WithContext(ctx).Clauses(conflict).CreateInBatches(rows, 500).Error
	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil
}

func (r *RepositoryMySQL) UpsertWarehouseA(rows []*entity.Anamnesis) (int64, error) {
	mdls := make([]*model.AnamnesisWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelA(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) UpsertWarehouseB(rows []*entity.AnamnesisB) (int64, error) {
	mdls := make([]*model.AnamnesisWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelB(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}
