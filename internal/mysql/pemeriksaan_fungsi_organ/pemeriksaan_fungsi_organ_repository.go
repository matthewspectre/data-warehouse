package pemeriksaan_fungsi_organ

import (
	"errors"
	"fmt"
	entity "rme/internal/entity/pemeriksaan_fungsi_organ"
	model "rme/internal/model/pemeriksaan_fungsi_organ"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.PemeriksaanFungsiOrgan) error {
	mm := model.PemeriksaanFungsiOrganModel{
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		Tanggal:       d.Tanggal,
		GangguanBAB:   d.GangguanBAB,
		GangguanBAK:   d.GangguanBAK,
		MualMuntah:    d.MualMuntah,
		Demam:         d.Demam,
		Perdarahan:    d.Perdarahan,
		PenurunanBB:   d.PenurunanBB,
		GangguanGerak: d.GangguanGerak,
		Nyeri:         d.Nyeri,
		SesakNapas:    d.SesakNapas,
		Pusing:        d.Pusing,
		Catatan:       d.Catatan,
		Visible:       1,
		DateMake:      d.Tanggal,
		DateUpdate:    d.Tanggal,
	}
	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.ID = mm.ID
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.PemeriksaanFungsiOrgan, error) {
	var mm model.PemeriksaanFungsiOrganModel
	if err := m.DB.Where("id = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	e := &entity.PemeriksaanFungsiOrgan{
		ID:            mm.ID,
		IDPasien:      mm.IDPasien,
		IDDokter:      mm.IDDokter,
		Tanggal:       mm.Tanggal,
		DateMake:      mm.DateMake,
		DateUpdate:    mm.DateUpdate,
		GangguanBAB:   mm.GangguanBAB,
		GangguanBAK:   mm.GangguanBAK,
		MualMuntah:    mm.MualMuntah,
		Demam:         mm.Demam,
		Perdarahan:    mm.Perdarahan,
		PenurunanBB:   mm.PenurunanBB,
		GangguanGerak: mm.GangguanGerak,
		Nyeri:         mm.Nyeri,
		SesakNapas:    mm.SesakNapas,
		Pusing:        mm.Pusing,
		Catatan:       mm.Catatan,
		Visible:       mm.Visible,
	}
	return e, nil
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error) {
	var list []model.PemeriksaanFungsiOrganModel
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
	out := make([]*entity.PemeriksaanFungsiOrgan, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanFungsiOrgan{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			Tanggal:       mm.Tanggal,
			DateMake:      mm.DateMake,
			DateUpdate:    mm.DateUpdate,
			GangguanBAB:   mm.GangguanBAB,
			GangguanBAK:   mm.GangguanBAK,
			MualMuntah:    mm.MualMuntah,
			Demam:         mm.Demam,
			Perdarahan:    mm.Perdarahan,
			PenurunanBB:   mm.PenurunanBB,
			GangguanGerak: mm.GangguanGerak,
			Nyeri:         mm.Nyeri,
			SesakNapas:    mm.SesakNapas,
			Pusing:        mm.Pusing,
			Catatan:       mm.Catatan,
			Visible:       mm.Visible,
		})
	}
	return out, nil
}

func (m *mysqlRepo) GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error) {
	var list []model.PemeriksaanFungsiOrganBModel
	q := m.DB.Table("`rme-b`.pemeriksaan_fungsi_organ").Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.PemeriksaanFungsiOrgan, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanFungsiOrgan{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			Tanggal:       mm.Tanggal,
			DateMake:      mm.DateMake,
			DateUpdate:    mm.DateUpdate,
			GangguanBAB:   mm.GangguanBAB,
			GangguanBAK:   mm.GangguanBAK,
			MualMuntah:    mm.MualMuntah,
			Demam:         mm.Demam,
			Perdarahan:    mm.Perdarahan,
			PenurunanBB:   mm.PenurunanBB,
			GangguanGerak: mm.GangguanGerak,
			Nyeri:         mm.Nyeri,
			SesakNapas:    mm.SesakNapas,
			Pusing:        mm.Pusing,
			Catatan:       mm.Catatan,
			Visible:       mm.Visible,
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

func toWarehouseModel(source string, e *entity.PemeriksaanFungsiOrgan) *model.PemeriksaanFungsiOrganWarehouseModel {
	if e == nil {
		return nil
	}
	return &model.PemeriksaanFungsiOrganWarehouseModel{
		Source:                   source,
		IDPemeriksaanFungsiOrgan: e.ID,
		IDPasien:                 e.IDPasien,
		NamaPasien:               nil,
		IDDokter:                 e.IDDokter,
		Tanggal:                  e.Tanggal,
		GangguanBAB:              e.GangguanBAB,
		GangguanBAK:              e.GangguanBAK,
		MualMuntah:               e.MualMuntah,
		Demam:                    e.Demam,
		Perdarahan:               e.Perdarahan,
		PenurunanBB:              e.PenurunanBB,
		GangguanGerak:            e.GangguanGerak,
		Nyeri:                    e.Nyeri,
		SesakNapas:               e.SesakNapas,
		Pusing:                   e.Pusing,
		Catatan:                  e.Catatan,
		Visible:                  e.Visible,
		DateMake:                 e.DateMake,
		DateUpdate:               e.DateUpdate,
	}
}

func (m *mysqlRepo) upsertWarehouse(rows []*model.PemeriksaanFungsiOrganWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_pemeriksaan_fungsi_organ"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"id_dokter",
			"tanggal",
			"gangguan_bab",
			"gangguan_bak",
			"mual_muntah",
			"demam",
			"perdarahan",
			"penurunan_bb",
			"gangguan_gerak",
			"nyeri",
			"sesak_napas",
			"pusing",
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

func (m *mysqlRepo) UpsertWarehouseA(rows []*entity.PemeriksaanFungsiOrgan) (int64, error) {
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

	mdls := make([]*model.PemeriksaanFungsiOrganWarehouseModel, 0, len(rows))
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

func (m *mysqlRepo) UpsertWarehouseB(rows []*entity.PemeriksaanFungsiOrgan) (int64, error) {
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

	mdls := make([]*model.PemeriksaanFungsiOrganWarehouseModel, 0, len(rows))
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

func (m *mysqlRepo) GetAllWarehouse(nik *string) ([]*entity.PemeriksaanFungsiOrganWarehouse, error) {
	var rows []model.PemeriksaanFungsiOrganWarehouseModel

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
			return []*entity.PemeriksaanFungsiOrganWarehouse{}, nil
		}
	}

	q := m.DB.Model(&model.PemeriksaanFungsiOrganWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_pemeriksaan_fungsi_organ DESC")

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

	out := make([]*entity.PemeriksaanFungsiOrganWarehouse, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		out = append(out, &entity.PemeriksaanFungsiOrganWarehouse{
			Source:                   r.Source,
			IDPemeriksaanFungsiOrgan: r.IDPemeriksaanFungsiOrgan,
			IDPasien:                 r.IDPasien,
			NamaPasien:               r.NamaPasien,
			IDDokter:                 r.IDDokter,
			Tanggal:                  r.Tanggal,
			DateMake:                 r.DateMake,
			DateUpdate:               r.DateUpdate,
			GangguanBAB:              r.GangguanBAB,
			GangguanBAK:              r.GangguanBAK,
			MualMuntah:               r.MualMuntah,
			Demam:                    r.Demam,
			Perdarahan:               r.Perdarahan,
			PenurunanBB:              r.PenurunanBB,
			GangguanGerak:            r.GangguanGerak,
			Nyeri:                    r.Nyeri,
			SesakNapas:               r.SesakNapas,
			Pusing:                   r.Pusing,
			Catatan:                  r.Catatan,
			Visible:                  r.Visible,
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
		case "tanggal":
			switch val := v.(type) {
			case string:
				if t, err := time.Parse(time.RFC3339, val); err == nil {
					dbUpdates["tanggal"] = t
				} else {
					dbUpdates["tanggal"] = val
				}
			case time.Time:
				dbUpdates["tanggal"] = val
			default:
				dbUpdates["tanggal"] = v
			}
		default:
			dbUpdates[k] = v
		}
	}
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.PemeriksaanFungsiOrganModel{}).Where("id = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.PemeriksaanFungsiOrganModel{}).
		Where("id = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
