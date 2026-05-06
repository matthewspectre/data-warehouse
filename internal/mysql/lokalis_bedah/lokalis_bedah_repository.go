package lokalis_bedah

import (
	"errors"

	entity "rme/internal/entity/lokalis_bedah"
	model "rme/internal/model/lokalis_bedah"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (m *mysqlRepo) GetAllB(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	var list []model.LokalisBedahBModel
	q := m.DB.Table("`rme-b`.status_lokalis_bedah").Where("visible = 1")
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
			Warna:          "",
			NyeriTekan:     mm.NyeriTekan,
			Konsistensi:    "",
			Mobilitas:      "",
			TandaRadang:    mm.TandaRadang,
			Fluktuasi:      mm.Fluktuasi,
			Catatan:        mm.Catatan,
			Visible:        mm.Visible,
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

func toWarehouseModel(source string, e *entity.LokalisBedah) *model.LokalisBedahWarehouseModel {
	if e == nil {
		return nil
	}
	return &model.LokalisBedahWarehouseModel{
		Source:               source,
		IDStatusLokalisBedah: e.ID,
		IDPasien:             e.IDPasien,
		NamaPasien:           nil,
		IDDokter:             e.IDDokter,
		Tanggal:              e.Tanggal,
		LokasiKelainan:       e.LokasiKelainan,
		JenisKelainan:        e.JenisKelainan,
		Ukuran:               e.Ukuran,
		Warna:                e.Warna,
		NyeriTekan:           e.NyeriTekan,
		Konsistensi:          e.Konsistensi,
		Mobilitas:            e.Mobilitas,
		TandaRadang:          e.TandaRadang,
		Fluktuasi:            e.Fluktuasi,
		Catatan:              e.Catatan,
		Visible:              e.Visible,
		DateMake:             e.Tanggal,
		DateUpdate:           e.Tanggal,
	}
}

func (m *mysqlRepo) upsertWarehouse(rows []*model.LokalisBedahWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_status_lokalis_bedah"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"id_dokter",
			"tanggal",
			"lokasi_kelainan",
			"jenis_kelainan",
			"ukuran",
			"warna",
			"nyeri_tekan",
			"konsistensi",
			"mobilitas",
			"tanda_radang",
			"fluktuasi",
			"catatan",
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

func (m *mysqlRepo) UpsertWarehouseA(rows []*entity.LokalisBedah) (int64, error) {
	ids := make([]int, 0, len(rows))
	seen := map[int]struct{}{}
	for _, e := range rows {
		if e == nil {
			continue
		}
		if _, ok := seen[e.IDPasien]; ok {
			continue
		}
		seen[e.IDPasien] = struct{}{}
		ids = append(ids, e.IDPasien)
	}
	patientNames, err := m.getPatientNamesRSA(ids)
	if err != nil {
		return 0, err
	}

	mdls := make([]*model.LokalisBedahWarehouseModel, 0, len(rows))
	for _, e := range rows {
		mdl := toWarehouseModel("rsA", e)
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

func (m *mysqlRepo) UpsertWarehouseB(rows []*entity.LokalisBedah) (int64, error) {
	ids := make([]int, 0, len(rows))
	seen := map[int]struct{}{}
	for _, e := range rows {
		if e == nil {
			continue
		}
		if _, ok := seen[e.IDPasien]; ok {
			continue
		}
		seen[e.IDPasien] = struct{}{}
		ids = append(ids, e.IDPasien)
	}
	patientNames, err := m.getPatientNamesRSB(ids)
	if err != nil {
		return 0, err
	}

	mdls := make([]*model.LokalisBedahWarehouseModel, 0, len(rows))
	for _, e := range rows {
		mdl := toWarehouseModel("rsB", e)
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

func (m *mysqlRepo) GetAllWarehouse(nik *string) ([]*entity.LokalisBedahWarehouse, error) {
	var rows []model.LokalisBedahWarehouseModel

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
			return []*entity.LokalisBedahWarehouse{}, nil
		}
	}

	q := m.DB.Model(&model.LokalisBedahWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_status_lokalis_bedah DESC")

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

	out := make([]*entity.LokalisBedahWarehouse, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		out = append(out, &entity.LokalisBedahWarehouse{
			Source:     r.Source,
			NamaPasien: r.NamaPasien,
			LokalisBedah: entity.LokalisBedah{
				ID:             r.IDStatusLokalisBedah,
				IDPasien:       r.IDPasien,
				IDDokter:       r.IDDokter,
				Tanggal:        r.Tanggal,
				LokasiKelainan: r.LokasiKelainan,
				JenisKelainan:  r.JenisKelainan,
				Ukuran:         r.Ukuran,
				Warna:          r.Warna,
				NyeriTekan:     r.NyeriTekan,
				Konsistensi:    r.Konsistensi,
				Mobilitas:      r.Mobilitas,
				TandaRadang:    r.TandaRadang,
				Fluktuasi:      r.Fluktuasi,
				Catatan:        r.Catatan,
				Visible:        r.Visible,
			},
		})
	}
	return out, nil
}
