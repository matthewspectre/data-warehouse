package pemeriksaanvital

import (
	"context"
	"errors"
	"time"

	entity "rme/internal/entity/pemeriksaan_vital"
	model "rme/internal/model/pemeriksaan_vital"
	repo "rme/internal/repository/pemeriksaan_vital"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.PemeriksaanVital) *model.PemeriksaanVitalModel {
	if e == nil {
		return nil
	}
	return &model.PemeriksaanVitalModel{
		ID:             e.ID,
		IDPasien:       e.IDPasien,
		IDDokter:       e.IDDokter,
		DateMake:       e.DateMake,
		DateUpdate:     e.DateUpdate,
		TekananDarah:   e.TekananDarah,
		DenyutNadi:     e.DenyutNadi,
		SuhuTubuh:      e.SuhuTubuh,
		FrekuensiNapas: e.FrekuensiNapas,
		BeratBadan:     e.BeratBadan,
		TinggiBadan:    e.TinggiBadan,
		Visible:        e.Visible,
	}
}

func toEntity(m *model.PemeriksaanVitalModel) *entity.PemeriksaanVital {
	if m == nil {
		return nil
	}
	return &entity.PemeriksaanVital{
		ID:             m.ID,
		IDPasien:       m.IDPasien,
		IDDokter:       m.IDDokter,
		DateMake:       m.DateMake,
		DateUpdate:     m.DateUpdate,
		TekananDarah:   m.TekananDarah,
		DenyutNadi:     m.DenyutNadi,
		SuhuTubuh:      m.SuhuTubuh,
		FrekuensiNapas: m.FrekuensiNapas,
		BeratBadan:     m.BeratBadan,
		TinggiBadan:    m.TinggiBadan,
		Visible:        m.Visible,
	}
}

func toEntityB(m *model.PemeriksaanVitalBModel) *entity.PemeriksaanVitalB {
	if m == nil {
		return nil
	}
	return &entity.PemeriksaanVitalB{
		ID:               m.ID,
		IDPasien:         m.IDPasien,
		IDDokter:         m.IDDokter,
		DateMake:         m.DateMake,
		DateUpdate:       m.DateUpdate,
		TekananDarah:     m.TekananDarah,
		DenyutNadi:       m.DenyutNadi,
		SuhuTubuh:        m.SuhuTubuh,
		FrekuensiNapas:   m.FrekuensiNapas,
		BeratBadan:       m.BeratBadan,
		TinggiBadan:      m.TinggiBadan,
		SaturasiOksigen:  m.SaturasiOksigen,
		TingkatKesadaran: m.TingkatKesadaran,
		IndeksMasaTubuh:  m.IndeksMasaTubuh,
		Visible:          m.Visible,
	}
}

func toEntityWarehouse(m *model.PemeriksaanVitalWarehouseModel) *entity.PemeriksaanVitalWarehouse {
	if m == nil {
		return nil
	}
	return &entity.PemeriksaanVitalWarehouse{
		Source:             m.Source,
		IDPemeriksaanVital: m.IDPemeriksaanVital,
		IDPasien:           m.IDPasien,
		NamaPasien:         m.NamaPasien,
		DateMake:           m.DateMake,
		DateUpdate:         m.DateUpdate,
		TekananDarah:       m.TekananDarah,
		DenyutNadi:         m.DenyutNadi,
		SuhuTubuh:          m.SuhuTubuh,
		FrekuensiNapas:     m.FrekuensiNapas,
		BeratBadan:         m.BeratBadan,
		TinggiBadan:        m.TinggiBadan,
		SaturasiOksigen:    m.SaturasiOksigen,
		TingkatKesadaran:   m.TingkatKesadaran,
		IndeksMasaTubuh:    m.IndeksMasaTubuh,
		Visible:            m.Visible,
	}
}

type joinedRow struct {
	ID             int       `gorm:"column:id_pemeriksaan_vital"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
	TekananDarah   string    `gorm:"column:tekanan_darah"`
	DenyutNadi     int       `gorm:"column:denyut_nadi"`
	SuhuTubuh      float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas int       `gorm:"column:frekuensi_napas"`
	BeratBadan     float64   `gorm:"column:berat_badan"`
	TinggiBadan    float64   `gorm:"column:tinggi_badan"`
	Visible        int       `gorm:"column:visible"`
	NamaPasien     string    `gorm:"column:nama_pasien"`
	NamaDokter     string    `gorm:"column:nama_dokter"`
}

type joinedRowB struct {
	ID               int       `gorm:"column:id_pemeriksaan_vital"`
	IDPasien         int       `gorm:"column:id_pasien"`
	IDDokter         int       `gorm:"column:id_dokter"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
	TekananDarah     string    `gorm:"column:tekanan_darah"`
	DenyutNadi       int       `gorm:"column:denyut_nadi"`
	SuhuTubuh        float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas   int       `gorm:"column:frekuensi_napas"`
	BeratBadan       float64   `gorm:"column:berat_badan"`
	TinggiBadan      float64   `gorm:"column:tinggi_badan"`
	SaturasiOksigen  int       `gorm:"column:saturasi_oksigen"`
	TingkatKesadaran string    `gorm:"column:tingkat_kesadaran"`
	IndeksMasaTubuh  float64   `gorm:"column:indeks_masa_tubuh"`
	Visible          int       `gorm:"column:visible"`
	NamaPasien       string    `gorm:"column:nama_pasien"`
	NamaDokter       string    `gorm:"column:nama_dokter"`
}

func rowToEntity(rj *joinedRow) *entity.PemeriksaanVital {
	if rj == nil {
		return nil
	}
	return &entity.PemeriksaanVital{
		ID:             rj.ID,
		IDPasien:       rj.IDPasien,
		IDDokter:       rj.IDDokter,
		DateMake:       rj.DateMake,
		DateUpdate:     rj.DateUpdate,
		TekananDarah:   rj.TekananDarah,
		DenyutNadi:     rj.DenyutNadi,
		SuhuTubuh:      rj.SuhuTubuh,
		FrekuensiNapas: rj.FrekuensiNapas,
		BeratBadan:     rj.BeratBadan,
		TinggiBadan:    rj.TinggiBadan,
		Visible:        rj.Visible,
		NamaPasien:     rj.NamaPasien,
		NamaDokter:     rj.NamaDokter,
	}
}

func rowToEntityB(rj *joinedRowB) *entity.PemeriksaanVitalB {
	if rj == nil {
		return nil
	}
	return &entity.PemeriksaanVitalB{
		ID:               rj.ID,
		IDPasien:         rj.IDPasien,
		IDDokter:         rj.IDDokter,
		DateMake:         rj.DateMake,
		DateUpdate:       rj.DateUpdate,
		TekananDarah:     rj.TekananDarah,
		DenyutNadi:       rj.DenyutNadi,
		SuhuTubuh:        rj.SuhuTubuh,
		FrekuensiNapas:   rj.FrekuensiNapas,
		BeratBadan:       rj.BeratBadan,
		TinggiBadan:      rj.TinggiBadan,
		SaturasiOksigen:  rj.SaturasiOksigen,
		TingkatKesadaran: rj.TingkatKesadaran,
		IndeksMasaTubuh:  rj.IndeksMasaTubuh,
		Visible:          rj.Visible,
		NamaPasien:       rj.NamaPasien,
		NamaDokter:       rj.NamaDokter,
	}
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error) {
	ctx := context.Background()
	args := make([]interface{}, 0)
	query := `SELECT pv.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
			  FROM pemeriksaan_vital pv
			  LEFT JOIN patients p ON p.id = pv.id_pasien
			  LEFT JOIN dokter d ON d.id = pv.id_dokter
			  WHERE pv.visible = 1`
	if idPasien != nil {
		query += " AND pv.id_pasien = ?"
		args = append(args, *idPasien)
	}
	if idDokter != nil {
		query += " AND pv.id_dokter = ?"
		args = append(args, *idDokter)
	}
	query += " ORDER BY pv.date_make DESC"

	var rows []joinedRow
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.PemeriksaanVital, 0, len(rows))
	for i := range rows {
		res = append(res, rowToEntity(&rows[i]))
	}
	return res, nil
}

func (r *RepositoryMySQL) GetAllB(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVitalB, error) {
	ctx := context.Background()
	args := make([]interface{}, 0)
	query := `SELECT pv.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
			  FROM ` + "`rme-b`" + `.pemeriksaan_vital pv
			  LEFT JOIN ` + "`rme-b`" + `.patients p ON p.id = pv.id_pasien
			  LEFT JOIN ` + "`rme-b`" + `.dokter d ON d.id = pv.id_dokter
			  WHERE pv.visible = 1`
	if idPasien != nil {
		query += " AND pv.id_pasien = ?"
		args = append(args, *idPasien)
	}
	if idDokter != nil {
		query += " AND pv.id_dokter = ?"
		args = append(args, *idDokter)
	}
	query += " ORDER BY pv.date_make DESC"

	var rows []joinedRowB
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.PemeriksaanVitalB, 0, len(rows))
	for i := range rows {
		res = append(res, rowToEntityB(&rows[i]))
	}
	return res, nil
}

func toWarehouseModelA(e *entity.PemeriksaanVital) *model.PemeriksaanVitalWarehouseModel {
	if e == nil {
		return nil
	}
	var nama *string
	if e.NamaPasien != "" {
		n := e.NamaPasien
		nama = &n
	}
	return &model.PemeriksaanVitalWarehouseModel{
		Source:             "rsA",
		IDPemeriksaanVital: e.ID,
		IDPasien:           e.IDPasien,
		NamaPasien:         nama,
		DateMake:           e.DateMake,
		DateUpdate:         e.DateUpdate,
		TekananDarah:       e.TekananDarah,
		DenyutNadi:         e.DenyutNadi,
		SuhuTubuh:          e.SuhuTubuh,
		FrekuensiNapas:     e.FrekuensiNapas,
		BeratBadan:         e.BeratBadan,
		TinggiBadan:        e.TinggiBadan,
		SaturasiOksigen:    nil,
		TingkatKesadaran:   nil,
		IndeksMasaTubuh:    nil,
		Visible:            e.Visible,
	}
}

func toWarehouseModelB(e *entity.PemeriksaanVitalB) *model.PemeriksaanVitalWarehouseModel {
	if e == nil {
		return nil
	}
	var nama *string
	if e.NamaPasien != "" {
		n := e.NamaPasien
		nama = &n
	}
	sat := e.SaturasiOksigen
	kes := e.TingkatKesadaran
	imt := e.IndeksMasaTubuh
	return &model.PemeriksaanVitalWarehouseModel{
		Source:             "rsB",
		IDPemeriksaanVital: e.ID,
		IDPasien:           e.IDPasien,
		NamaPasien:         nama,
		DateMake:           e.DateMake,
		DateUpdate:         e.DateUpdate,
		TekananDarah:       e.TekananDarah,
		DenyutNadi:         e.DenyutNadi,
		SuhuTubuh:          e.SuhuTubuh,
		FrekuensiNapas:     e.FrekuensiNapas,
		BeratBadan:         e.BeratBadan,
		TinggiBadan:        e.TinggiBadan,
		SaturasiOksigen:    &sat,
		TingkatKesadaran:   &kes,
		IndeksMasaTubuh:    &imt,
		Visible:            e.Visible,
	}
}

func (r *RepositoryMySQL) upsertWarehouse(rows []*model.PemeriksaanVitalWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	ctx := context.Background()

	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_pemeriksaan_vital"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"date_make",
			"date_update",
			"tekanan_darah",
			"denyut_nadi",
			"suhu_tubuh",
			"frekuensi_napas",
			"berat_badan",
			"tinggi_badan",
			"saturasi_oksigen",
			"tingkat_kesadaran",
			"indeks_masa_tubuh",
			"visible",
		}),
	}

	err := r.db.WithContext(ctx).Clauses(conflict).CreateInBatches(rows, 500).Error
	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil
}

func (r *RepositoryMySQL) UpsertWarehouseA(rows []*entity.PemeriksaanVital) (int64, error) {
	mdls := make([]*model.PemeriksaanVitalWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelA(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) UpsertWarehouseB(rows []*entity.PemeriksaanVitalB) (int64, error) {
	mdls := make([]*model.PemeriksaanVitalWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelB(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

// GetAllWarehouse mengambil semua data hasil ETL dari `data_warehouse`.pemeriksaan_vital.
// Default: hanya yang visible=1.
// Jika nik diberikan, maka akan dicari dulu id_pasien di DB rme dan rme-b,
// lalu data warehouse difilter sesuai source masing-masing.
func (r *RepositoryMySQL) GetAllWarehouse(nik *string) ([]*entity.PemeriksaanVitalWarehouse, error) {
	ctx := context.Background()
	var rows []model.PemeriksaanVitalWarehouseModel

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
			return []*entity.PemeriksaanVitalWarehouse{}, nil
		}
	}

	q := r.db.WithContext(ctx).
		Model(&model.PemeriksaanVitalWarehouseModel{}).
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
	res := make([]*entity.PemeriksaanVitalWarehouse, 0, len(rows))
	for i := range rows {
		ent := toEntityWarehouse(&rows[i])
		if ent != nil {
			res = append(res, ent)
		}
	}
	return res, nil
}
