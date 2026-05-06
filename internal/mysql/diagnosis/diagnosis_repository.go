package diagnosis

import (
	"encoding/json"
	"errors"
	"fmt"

	entity "rme/internal/entity/diagnosis"
	model "rme/internal/model/diagnosis"
	icdModel "rme/internal/model/icd"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.Diagnosis) error {
	mm := model.DiagnosisModel{
		IDPasien:     d.IDPasien,
		IDDokter:     d.IDDokter,
		Tanggal:      d.Tanggal,
		KodeIcdUtama: d.KodeIcdUtama,
		Status:       d.Status,
		Catatan:      d.Catatan,
		Visible:      1,
		DateMake:     d.Tanggal,
		DateUpdate:   d.Tanggal,
	}

	if b, err := json.Marshal(d.KodeIcdSekunder); err == nil {
		mm.KodeIcdSekunder = string(b)
	}
	if b, err := json.Marshal(d.DiagnosisBanding); err == nil {
		mm.DiagnosisBanding = string(b)
	}
	if b, err := json.Marshal(d.DasarDiagnosis); err == nil {
		mm.DasarDiagnosis = string(b)
	}

	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.IDDiagnosis = mm.IDDiagnosis
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.Diagnosis, error) {
	var mm model.DiagnosisModel
	if err := m.DB.Where("id_diagnosis = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m.mapModelToEntity(&mm)
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.Diagnosis, error) {
	var list []model.DiagnosisModel
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
	out := make([]*entity.Diagnosis, 0, len(list))
	for i := range list {
		e, err := m.mapModelToEntity(&list[i])
		if err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}

func (m *mysqlRepo) GetAllB(idDokter *int, idPasien *int) ([]*entity.Diagnosis, error) {
	var list []model.DiagnosisBModel
	q := m.DB.Table("`rme-b`.diagnosis").Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Diagnosis, 0, len(list))
	for i := range list {
		e, err := m.mapModelBToEntity(&list[i])
		if err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}

func (m *mysqlRepo) mapModelToEntity(mm *model.DiagnosisModel) (*entity.Diagnosis, error) {
	var sec []string
	var band []string
	var dasar []string
	if mm.KodeIcdSekunder != "" {
		_ = json.Unmarshal([]byte(mm.KodeIcdSekunder), &sec)
	}
	if mm.DiagnosisBanding != "" {
		_ = json.Unmarshal([]byte(mm.DiagnosisBanding), &band)
	}
	if mm.DasarDiagnosis != "" {
		_ = json.Unmarshal([]byte(mm.DasarDiagnosis), &dasar)
	}

	e := &entity.Diagnosis{
		IDDiagnosis:      mm.IDDiagnosis,
		IDPasien:         mm.IDPasien,
		IDDokter:         mm.IDDokter,
		Tanggal:          mm.Tanggal,
		KodeIcdUtama:     mm.KodeIcdUtama,
		KodeIcdSekunder:  sec,
		DiagnosisBanding: band,
		Status:           mm.Status,
		DasarDiagnosis:   dasar,
		Catatan:          mm.Catatan,
		Visible:          mm.Visible,
	}

	// Fetch nama for utama and sekunder by joining icd10 table
	// utama
	if e.KodeIcdUtama != "" {
		var icd icdModel.IcdModel
		if err := m.DB.Where("kode = ?", e.KodeIcdUtama).First(&icd).Error; err == nil {
			// attach name by replacing kode with "kode|nama" convention? We will keep kode fields and resolve names in usecase/handler
			// Alternatively, store mapping in entity by convention — leave names to higher layer.
			_ = icd
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return e, nil
}

func (m *mysqlRepo) mapModelBToEntity(mm *model.DiagnosisBModel) (*entity.Diagnosis, error) {
	var sec []string
	var band []string
	var dasar []string
	if mm.KodeIcdSekunder != "" {
		_ = json.Unmarshal([]byte(mm.KodeIcdSekunder), &sec)
	}
	if mm.DiagnosisBanding != "" {
		_ = json.Unmarshal([]byte(mm.DiagnosisBanding), &band)
	}
	if mm.DasarDiagnosis != "" {
		_ = json.Unmarshal([]byte(mm.DasarDiagnosis), &dasar)
	}

	e := &entity.Diagnosis{
		IDDiagnosis:      mm.IDDiagnosis,
		IDPasien:         mm.IDPasien,
		IDDokter:         mm.IDDokter,
		Tanggal:          mm.Tanggal,
		KodeIcdUtama:     mm.KodeIcdUtama,
		KodeIcdSekunder:  sec,
		DiagnosisBanding: band,
		Status:           mm.Status,
		DasarDiagnosis:   dasar,
		Catatan:          mm.Catatan,
		Visible:          mm.Visible,
	}
	return e, nil
}

func toWarehouseModel(source string, d *entity.Diagnosis) (*model.DiagnosisWarehouseModel, error) {
	if d == nil {
		return nil, nil
	}

	sek := "[]"
	if b, err := json.Marshal(d.KodeIcdSekunder); err == nil {
		sek = string(b)
	} else {
		return nil, err
	}
	band := "[]"
	if b, err := json.Marshal(d.DiagnosisBanding); err == nil {
		band = string(b)
	} else {
		return nil, err
	}
	dasar := "[]"
	if b, err := json.Marshal(d.DasarDiagnosis); err == nil {
		dasar = string(b)
	} else {
		return nil, err
	}

	return &model.DiagnosisWarehouseModel{
		Source:           source,
		IDDiagnosis:      d.IDDiagnosis,
		IDPasien:         d.IDPasien,
		IDDokter:         d.IDDokter,
		NamaPasien:       nil,
		Tanggal:          d.Tanggal,
		KodeIcdUtama:     d.KodeIcdUtama,
		KodeIcdSekunder:  sek,
		DiagnosisBanding: band,
		Status:           d.Status,
		DasarDiagnosis:   dasar,
		Catatan:          d.Catatan,
		Visible:          d.Visible,
		DateMake:         d.Tanggal,
		DateUpdate:       d.Tanggal,
	}, nil
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

func (m *mysqlRepo) upsertWarehouse(rows []*model.DiagnosisWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}

	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_diagnosis"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"id_dokter",
			"nama_pasien",
			"tanggal",
			"kode_icd_utama",
			"kode_icd_sekunder",
			"diagnosis_banding",
			"status",
			"dasar_diagnosis",
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

func (m *mysqlRepo) ETLUpsertWarehouseA(rows []*entity.Diagnosis) (int64, error) {
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

	mdls := make([]*model.DiagnosisWarehouseModel, 0, len(rows))
	for _, d := range rows {
		mdl, err := toWarehouseModel("rsA", d)
		if err != nil {
			return 0, err
		}
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

func (m *mysqlRepo) ETLUpsertWarehouseB(rows []*entity.Diagnosis) (int64, error) {
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

	mdls := make([]*model.DiagnosisWarehouseModel, 0, len(rows))
	for _, d := range rows {
		mdl, err := toWarehouseModel("rsB", d)
		if err != nil {
			return 0, err
		}
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

func (m *mysqlRepo) GetAllWarehouse(nik *string) ([]*entity.DiagnosisWarehouse, error) {
	var rows []model.DiagnosisWarehouseModel

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
			return []*entity.DiagnosisWarehouse{}, nil
		}
	}

	q := m.DB.Model(&model.DiagnosisWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_diagnosis DESC")

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

	out := make([]*entity.DiagnosisWarehouse, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		var sec []string
		var band []string
		var dasar []string
		if r.KodeIcdSekunder != "" {
			_ = json.Unmarshal([]byte(r.KodeIcdSekunder), &sec)
		}
		if r.DiagnosisBanding != "" {
			_ = json.Unmarshal([]byte(r.DiagnosisBanding), &band)
		}
		if r.DasarDiagnosis != "" {
			_ = json.Unmarshal([]byte(r.DasarDiagnosis), &dasar)
		}

		out = append(out, &entity.DiagnosisWarehouse{
			Source:     r.Source,
			NamaPasien: r.NamaPasien,
			Diagnosis: entity.Diagnosis{
				IDDiagnosis:      r.IDDiagnosis,
				IDPasien:         r.IDPasien,
				IDDokter:         r.IDDokter,
				Tanggal:          r.Tanggal,
				KodeIcdUtama:     r.KodeIcdUtama,
				KodeIcdSekunder:  sec,
				DiagnosisBanding: band,
				Status:           r.Status,
				DasarDiagnosis:   dasar,
				Catatan:          r.Catatan,
				Visible:          r.Visible,
			},
		})
	}

	return out, nil
}

// helper to bulk fetch ICD names — used by usecase/handler
func (m *mysqlRepo) GetIcdNamesForCodes(codes []string) (map[string]string, error) {
	out := map[string]string{}
	if len(codes) == 0 {
		return out, nil
	}
	var icds []icdModel.IcdModel
	if err := m.DB.Where("kode IN (?)", codes).Find(&icds).Error; err != nil {
		return nil, err
	}
	for _, v := range icds {
		out[v.Kode] = v.Nama
	}
	return out, nil
}

// ensure mysqlRepo implements Repository
var _ = fmt.Sprintf

func (m *mysqlRepo) Update(id int, updates map[string]interface{}) error {
	dbUpdates := map[string]interface{}{}

	// map known update keys to columns; marshal slices to JSON strings
	for k, v := range updates {
		switch k {
		case "kode_icd_utama":
			if s, ok := v.(string); ok {
				dbUpdates["kode_icd_utama"] = s
			}
		case "kode_icd_sekunder":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["kode_icd_sekunder"] = string(b)
				}
			}
		case "diagnosis_banding":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["diagnosis_banding"] = string(b)
				}
			}
		case "dasar_diagnosis":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["dasar_diagnosis"] = string(b)
				}
			}
		case "status":
			if s, ok := v.(string); ok {
				dbUpdates["status"] = s
			}
		case "catatan":
			if s, ok := v.(string); ok {
				dbUpdates["catatan"] = s
			}
		case "tanggal":
			if t, ok := v.(string); ok {
				dbUpdates["tanggal"] = t
			}
		}
	}

	// always update date_update
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.DiagnosisModel{}).Where("id_diagnosis = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.DiagnosisModel{}).
		Where("id_diagnosis = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
