package pemeriksaan_penunjang_bedah

import (
	"errors"
	"fmt"
	"time"

	entity "rme/internal/entity/pemeriksaan_penunjang_bedah"
	model "rme/internal/model/pemeriksaan_penunjang_bedah"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.PemeriksaanPenunjangBedah) error {
	mm := model.PemeriksaanPenunjangBedahModel{
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		ButuhUSG:      d.ButuhUSG,
		ButuhRontgen:  d.ButuhRontgen,
		ButuhCTScan:   d.ButuhCTScan,
		ButuhBiopsi:   d.ButuhBiopsi,
		StatusOperasi: d.StatusOperasi,
		JenisTindakan: d.JenisTindakan,
		Prioritas:     d.Prioritas,
		CatatanBedah:  d.CatatanBedah,
		JadwalBedah:   d.JadwalBedah,
		Visible:       1,
		DateMake:      time.Now(),
		DateUpdate:    time.Now(),
	}
	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.ID = mm.ID
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error) {
	var mm model.PemeriksaanPenunjangBedahModel
	if err := m.DB.Where("id = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	e := &entity.PemeriksaanPenunjangBedah{
		ID:            mm.ID,
		IDPasien:      mm.IDPasien,
		IDDokter:      mm.IDDokter,
		ButuhUSG:      mm.ButuhUSG,
		ButuhRontgen:  mm.ButuhRontgen,
		ButuhCTScan:   mm.ButuhCTScan,
		ButuhBiopsi:   mm.ButuhBiopsi,
		StatusOperasi: mm.StatusOperasi,
		JenisTindakan: mm.JenisTindakan,
		Prioritas:     mm.Prioritas,
		CatatanBedah:  mm.CatatanBedah,
		JadwalBedah:   mm.JadwalBedah,
		Visible:       mm.Visible,
		DateMake:      mm.DateMake,
		DateUpdate:    mm.DateUpdate,
	}
	return e, nil
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	var list []model.PemeriksaanPenunjangBedahModel
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
	out := make([]*entity.PemeriksaanPenunjangBedah, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanPenunjangBedah{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			ButuhUSG:      mm.ButuhUSG,
			ButuhRontgen:  mm.ButuhRontgen,
			ButuhCTScan:   mm.ButuhCTScan,
			ButuhBiopsi:   mm.ButuhBiopsi,
			StatusOperasi: mm.StatusOperasi,
			JenisTindakan: mm.JenisTindakan,
			Prioritas:     mm.Prioritas,
			CatatanBedah:  mm.CatatanBedah,
			JadwalBedah:   mm.JadwalBedah,
			Visible:       mm.Visible,
			DateMake:      mm.DateMake,
			DateUpdate:    mm.DateUpdate,
		})
	}
	return out, nil
}

func (m *mysqlRepo) GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	var list []model.PemeriksaanPenunjangBedahBModel
	q := m.DB.Table("`rme-b`.pemeriksaan_penunjang_bedah").Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.PemeriksaanPenunjangBedah, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanPenunjangBedah{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			ButuhUSG:      mm.ButuhUSG,
			ButuhRontgen:  mm.ButuhRontgen,
			ButuhCTScan:   mm.ButuhCTScan,
			ButuhBiopsi:   mm.ButuhBiopsi,
			StatusOperasi: "",
			JenisTindakan: mm.JenisTindakan,
			Prioritas:     mm.Prioritas,
			CatatanBedah:  mm.CatatanBedah,
			JadwalBedah:   mm.JadwalBedah,
			Visible:       mm.Visible,
			DateMake:      mm.DateMake,
			DateUpdate:    mm.DateUpdate,
		})
	}
	return out, nil
}

func (m *mysqlRepo) getPatientNamesRSA(ids []int) (map[int]string, error) {
	out := map[int]string{}
	if len(ids) == 0 {
		return out, nil
	}
	type row struct {
		ID   int    `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}
	var rows []row
	if err := m.DB.Table("patients").Select("id, name").Where("id IN (?)", ids).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, r := range rows {
		out[r.ID] = r.Name
	}
	return out, nil
}

func (m *mysqlRepo) getPatientNamesRSB(ids []int) (map[int]string, error) {
	out := map[int]string{}
	if len(ids) == 0 {
		return out, nil
	}
	type row struct {
		ID   int    `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}
	var rows []row
	if err := m.DB.Table("`rme-b`.patients").Select("id, name").Where("id IN (?)", ids).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, r := range rows {
		out[r.ID] = r.Name
	}
	return out, nil
}

func toWarehouseModel(source string, e *entity.PemeriksaanPenunjangBedah) *model.PemeriksaanPenunjangBedahWarehouseModel {
	if e == nil {
		return nil
	}
	return &model.PemeriksaanPenunjangBedahWarehouseModel{
		Source:                      source,
		IDPemeriksaanPenunjangBedah: e.ID,
		IDPasien:                    e.IDPasien,
		NamaPasien:                  nil,
		IDDokter:                    e.IDDokter,
		ButuhUSG:                    e.ButuhUSG,
		ButuhRontgen:                e.ButuhRontgen,
		ButuhCTScan:                 e.ButuhCTScan,
		ButuhBiopsi:                 e.ButuhBiopsi,
		StatusOperasi:               e.StatusOperasi,
		JenisTindakan:               e.JenisTindakan,
		Prioritas:                   e.Prioritas,
		CatatanBedah:                e.CatatanBedah,
		JadwalBedah:                 e.JadwalBedah,
		Visible:                     e.Visible,
		DateMake:                    e.DateMake,
		DateUpdate:                  e.DateUpdate,
	}
}

func (m *mysqlRepo) upsertWarehouse(rows []*model.PemeriksaanPenunjangBedahWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_pemeriksaan_penunjang_bedah"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"id_dokter",
			"butuh_usg",
			"butuh_rontgen",
			"butuh_ctscan",
			"butuh_biopsi",
			"status_operasi",
			"jenis_tindakan",
			"prioritas",
			"catatan_bedah",
			"jadwal_bedah",
			"visible",
			"date_make",
			"date_update",
		}),
	}

	err := m.DB.Clauses(conflict).CreateInBatches(rows, 500).Error
	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil
}

func (m *mysqlRepo) UpsertWarehouseA(rows []*entity.PemeriksaanPenunjangBedah) (int64, error) {
	ids := make([]int, 0, len(rows))
	seen := map[int]struct{}{}
	for _, d := range rows {
		if d == nil {
			continue
		}
		if _, ok := seen[d.IDPasien]; ok {
			continue
		}
		seen[d.IDPasien] = struct{}{}
		ids = append(ids, d.IDPasien)
	}
	patientNames, err := m.getPatientNamesRSA(ids)
	if err != nil {
		return 0, err
	}

	mdls := make([]*model.PemeriksaanPenunjangBedahWarehouseModel, 0, len(rows))
	for _, d := range rows {
		mdl := toWarehouseModel("rsA", d)
		if mdl != nil {
			if name, ok := patientNames[mdl.IDPasien]; ok && name != "" {
				n := name
				mdl.NamaPasien = &n
			}
			mdls = append(mdls, mdl)
		}
	}
	return m.upsertWarehouse(mdls)
}

func (m *mysqlRepo) UpsertWarehouseB(rows []*entity.PemeriksaanPenunjangBedah) (int64, error) {
	ids := make([]int, 0, len(rows))
	seen := map[int]struct{}{}
	for _, d := range rows {
		if d == nil {
			continue
		}
		if _, ok := seen[d.IDPasien]; ok {
			continue
		}
		seen[d.IDPasien] = struct{}{}
		ids = append(ids, d.IDPasien)
	}
	patientNames, err := m.getPatientNamesRSB(ids)
	if err != nil {
		return 0, err
	}

	mdls := make([]*model.PemeriksaanPenunjangBedahWarehouseModel, 0, len(rows))
	for _, d := range rows {
		mdl := toWarehouseModel("rsB", d)
		if mdl != nil {
			if name, ok := patientNames[mdl.IDPasien]; ok && name != "" {
				n := name
				mdl.NamaPasien = &n
			}
			mdls = append(mdls, mdl)
		}
	}
	return m.upsertWarehouse(mdls)
}

func (m *mysqlRepo) GetAllWarehouse(nik *string) ([]*entity.PemeriksaanPenunjangBedahWarehouse, error) {
	var rows []model.PemeriksaanPenunjangBedahWarehouseModel

	var idPasienRSA *int
	var idPasienRSB *int
	if nik != nil && *nik != "" {
		type patientIDRow struct {
			ID int `gorm:"column:id"`
		}

		var rowA patientIDRow
		txA := m.DB.Table("patients").Select("id").Where("nik = ?", *nik).Take(&rowA)
		if txA.Error != nil && !errors.Is(txA.Error, gorm.ErrRecordNotFound) {
			return nil, txA.Error
		}
		if txA.Error == nil {
			v := rowA.ID
			idPasienRSA = &v
		}

		var rowB patientIDRow
		txB := m.DB.Table("`rme-b`.patients").Select("id").Where("nik = ?", *nik).Take(&rowB)
		if txB.Error != nil && !errors.Is(txB.Error, gorm.ErrRecordNotFound) {
			return nil, txB.Error
		}
		if txB.Error == nil {
			v := rowB.ID
			idPasienRSB = &v
		}

		if idPasienRSA == nil && idPasienRSB == nil {
			return []*entity.PemeriksaanPenunjangBedahWarehouse{}, nil
		}
	}

	q := m.DB.Model(&model.PemeriksaanPenunjangBedahWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_pemeriksaan_penunjang_bedah DESC")

	if idPasienRSA != nil || idPasienRSB != nil {
		var cond *gorm.DB
		if idPasienRSA != nil {
			cond = m.DB.Where("source = ? AND id_pasien = ?", "rsA", *idPasienRSA)
		}
		if idPasienRSB != nil {
			if cond == nil {
				cond = m.DB.Where("source = ? AND id_pasien = ?", "rsB", *idPasienRSB)
			} else {
				cond = cond.Or("source = ? AND id_pasien = ?", "rsB", *idPasienRSB)
			}
		}
		q = q.Where(cond)
	}

	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}

	out := make([]*entity.PemeriksaanPenunjangBedahWarehouse, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		out = append(out, &entity.PemeriksaanPenunjangBedahWarehouse{
			Source:                      r.Source,
			IDPemeriksaanPenunjangBedah: r.IDPemeriksaanPenunjangBedah,
			IDPasien:                    r.IDPasien,
			NamaPasien:                  r.NamaPasien,
			IDDokter:                    r.IDDokter,
			ButuhUSG:                    r.ButuhUSG,
			ButuhRontgen:                r.ButuhRontgen,
			ButuhCTScan:                 r.ButuhCTScan,
			ButuhBiopsi:                 r.ButuhBiopsi,
			StatusOperasi:               r.StatusOperasi,
			JenisTindakan:               r.JenisTindakan,
			Prioritas:                   r.Prioritas,
			CatatanBedah:                r.CatatanBedah,
			JadwalBedah:                 r.JadwalBedah,
			Visible:                     r.Visible,
			DateMake:                    r.DateMake,
			DateUpdate:                  r.DateUpdate,
		})
	}
	return out, nil
}

// ensure mysqlRepo implements Repository
var _ = fmt.Sprintf

func (m *mysqlRepo) Update(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	dbUpdates := map[string]interface{}{}
	for k, v := range updates {
		switch k {
		case "jadwalBedah", "jadwal_bedah":
			switch val := v.(type) {
			case string:
				if t, err := time.Parse(time.RFC3339, val); err == nil {
					dbUpdates["jadwal_bedah"] = t
				} else {
					dbUpdates["jadwal_bedah"] = val
				}
			case time.Time:
				dbUpdates["jadwal_bedah"] = val
			default:
				dbUpdates["jadwal_bedah"] = v
			}
		case "statusOperasi", "status_operasi":
			dbUpdates["status_operasi"] = v
		case "jenisTindakan", "jenis_tindakan":
			dbUpdates["jenis_tindakan"] = v
		case "prioritas":
			dbUpdates["prioritas"] = v
		case "catatanBedah", "catatan_bedah":
			dbUpdates["catatan_bedah"] = v
		case "butuhUSG", "butuh_usg":
			dbUpdates["butuh_usg"] = v
		case "butuhRontgen", "butuh_rontgen":
			dbUpdates["butuh_rontgen"] = v
		case "butuhCTScan", "butuh_ctscan":
			dbUpdates["butuh_ctscan"] = v
		case "butuhBiopsi", "butuh_biopsi":
			dbUpdates["butuh_biopsi"] = v
		case "id_dokter":
			dbUpdates["id_dokter"] = v
		default:
			dbUpdates[k] = v
		}
	}
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).Where("id = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).
		Where("id = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) HideByPatient(idPasien int) error {
	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).
		Where("id_pasien = ? AND visible = 1", idPasien).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
