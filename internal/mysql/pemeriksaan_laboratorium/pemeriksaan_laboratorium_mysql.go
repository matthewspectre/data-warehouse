package pemeriksaan_laboratorium

import (
	"context"
	"errors"

	model "rme/internal/model/pemeriksaan_laboratorium"
	repo "rme/internal/repository/pemeriksaan_laboratorium"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanLaboratoriumWithNames, error) {
	var rows []repo.PemeriksaanLaboratoriumWithNames
	q := r.db.Table("pemeriksaan_laboratorium as pl").
		Select("pl.id, pl.id_pasien, p.name as nama_pasien, pl.id_dokter, d.nama_dokter as nama_dokter, pl.hb, pl.ht, pl.leukosit, pl.trombosit, pl.gula_puasa, pl.gula_sewaktu, pl.hba1c AS hb_a1c, pl.kolesterol_total AS kolesterol, pl.hdl, pl.ldl, pl.trigliserida, pl.sgot, pl.sgpt, pl.ureum, pl.kreatinin, pl.asam_urat, pl.natrium, pl.kalium, pl.klorida, pl.visible, pl.waktu").
		Joins("JOIN patients p ON pl.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pl.id_dokter = d.id").
		Where("pl.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pl.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pl.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanLaboratoriumWithNames, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func (r *RepositoryMySQL) GetAllB(idPasien *int, idDokter *int) ([]*repo.PemeriksaanLaboratoriumWithNamesB, error) {
	var rows []repo.PemeriksaanLaboratoriumWithNamesB
	q := r.db.Table("`rme-b`.pemeriksaan_laboratorium as pl").
		Select("pl.id, pl.id_pasien, p.name as nama_pasien, pl.id_dokter, d.nama_dokter as nama_dokter, pl.hb, pl.ht, pl.leukosit, pl.trombosit, pl.gula_puasa, pl.gula_sewaktu, pl.hba1c AS hb_a1c, pl.kolesterol_total AS kolesterol, pl.hdl, pl.ldl, pl.trigliserida, pl.sgot, pl.sgpt, pl.ureum, pl.kreatinin, pl.asam_urat, pl.natrium, pl.kalium, pl.klorida, pl.visible, pl.waktu, pl.laju_endap_darah, pl.albumin").
		Joins("JOIN `rme-b`.patients p ON pl.id_pasien = p.id").
		Joins("LEFT JOIN `rme-b`.dokter d ON pl.id_dokter = d.id").
		Where("pl.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pl.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pl.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanLaboratoriumWithNamesB, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func toWarehouseModelA(rw *repo.PemeriksaanLaboratoriumWithNames) *model.PemeriksaanLaboratoriumWarehouseModel {
	if rw == nil {
		return nil
	}
	var nama *string
	if rw.NamaPasien != "" {
		n := rw.NamaPasien
		nama = &n
	}
	return &model.PemeriksaanLaboratoriumWarehouseModel{
		Source:                    "rsA",
		IDPemeriksaanLaboratorium: rw.ID,
		IDPasien:                  rw.IDPasien,
		NamaPasien:                nama,
		Hb:                        rw.Hb,
		Ht:                        rw.Ht,
		Leukosit:                  rw.Leukosit,
		Trombosit:                 rw.Trombosit,
		GulaPuasa:                 rw.GulaPuasa,
		GulaSewaktu:               rw.GulaSewaktu,
		HbA1c:                     rw.HbA1c,
		Kolesterol:                rw.Kolesterol,
		HDL:                       rw.HDL,
		LDL:                       rw.LDL,
		Trigliserida:              rw.Trigliserida,
		SGOT:                      rw.SGOT,
		SGPT:                      rw.SGPT,
		Ureum:                     rw.Ureum,
		Kreatinin:                 rw.Kreatinin,
		AsamUrat:                  rw.AsamUrat,
		Natrium:                   rw.Natrium,
		Kalium:                    rw.Kalium,
		Klorida:                   rw.Klorida,
		Waktu:                     rw.Waktu,
		LajuEndapDarah:            nil,
		Albumin:                   nil,
		Visible:                   rw.Visible,
	}
}

func toWarehouseModelB(rw *repo.PemeriksaanLaboratoriumWithNamesB) *model.PemeriksaanLaboratoriumWarehouseModel {
	if rw == nil {
		return nil
	}
	var nama *string
	if rw.NamaPasien != "" {
		n := rw.NamaPasien
		nama = &n
	}
	return &model.PemeriksaanLaboratoriumWarehouseModel{
		Source:                    "rsB",
		IDPemeriksaanLaboratorium: rw.ID,
		IDPasien:                  rw.IDPasien,
		NamaPasien:                nama,
		Hb:                        rw.Hb,
		Ht:                        rw.Ht,
		Leukosit:                  rw.Leukosit,
		Trombosit:                 rw.Trombosit,
		GulaPuasa:                 rw.GulaPuasa,
		GulaSewaktu:               rw.GulaSewaktu,
		HbA1c:                     rw.HbA1c,
		Kolesterol:                rw.Kolesterol,
		HDL:                       rw.HDL,
		LDL:                       rw.LDL,
		Trigliserida:              rw.Trigliserida,
		SGOT:                      rw.SGOT,
		SGPT:                      rw.SGPT,
		Ureum:                     rw.Ureum,
		Kreatinin:                 rw.Kreatinin,
		AsamUrat:                  rw.AsamUrat,
		Natrium:                   rw.Natrium,
		Kalium:                    rw.Kalium,
		Klorida:                   rw.Klorida,
		Waktu:                     rw.Waktu,
		LajuEndapDarah:            rw.LajuEndapDarah,
		Albumin:                   rw.Albumin,
		Visible:                   rw.Visible,
	}
}

func (r *RepositoryMySQL) upsertWarehouse(rows []*model.PemeriksaanLaboratoriumWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	ctx := context.Background()

	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_pemeriksaan_laboratorium"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"hb",
			"ht",
			"leukosit",
			"trombosit",
			"gula_puasa",
			"gula_sewaktu",
			"hba1c",
			"kolesterol_total",
			"hdl",
			"ldl",
			"trigliserida",
			"sgot",
			"sgpt",
			"ureum",
			"kreatinin",
			"asam_urat",
			"natrium",
			"kalium",
			"klorida",
			"waktu",
			"laju_endap_darah",
			"albumin",
			"visible",
		}),
	}

	err := r.db.WithContext(ctx).Clauses(conflict).CreateInBatches(rows, 500).Error
	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil
}

func (r *RepositoryMySQL) UpsertWarehouseA(rows []*repo.PemeriksaanLaboratoriumWithNames) (int64, error) {
	mdls := make([]*model.PemeriksaanLaboratoriumWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelA(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) UpsertWarehouseB(rows []*repo.PemeriksaanLaboratoriumWithNamesB) (int64, error) {
	mdls := make([]*model.PemeriksaanLaboratoriumWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelB(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) GetAllWarehouse(nik *string) ([]*repo.PemeriksaanLaboratoriumWarehouse, error) {
	ctx := context.Background()
	var rows []model.PemeriksaanLaboratoriumWarehouseModel

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
			return []*repo.PemeriksaanLaboratoriumWarehouse{}, nil
		}
	}

	q := r.db.WithContext(ctx).
		Model(&model.PemeriksaanLaboratoriumWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_pemeriksaan_laboratorium DESC")

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

	res := make([]*repo.PemeriksaanLaboratoriumWarehouse, 0, len(rows))
	for i := range rows {
		rw := &rows[i]
		res = append(res, &repo.PemeriksaanLaboratoriumWarehouse{
			Source:                    rw.Source,
			IDPemeriksaanLaboratorium: rw.IDPemeriksaanLaboratorium,
			IDPasien:                  rw.IDPasien,
			NamaPasien:                rw.NamaPasien,
			Hb:                        rw.Hb,
			Ht:                        rw.Ht,
			Leukosit:                  rw.Leukosit,
			Trombosit:                 rw.Trombosit,
			GulaPuasa:                 rw.GulaPuasa,
			GulaSewaktu:               rw.GulaSewaktu,
			HbA1c:                     rw.HbA1c,
			Kolesterol:                rw.Kolesterol,
			HDL:                       rw.HDL,
			LDL:                       rw.LDL,
			Trigliserida:              rw.Trigliserida,
			SGOT:                      rw.SGOT,
			SGPT:                      rw.SGPT,
			Ureum:                     rw.Ureum,
			Kreatinin:                 rw.Kreatinin,
			AsamUrat:                  rw.AsamUrat,
			Natrium:                   rw.Natrium,
			Kalium:                    rw.Kalium,
			Klorida:                   rw.Klorida,
			Waktu:                     rw.Waktu,
			LajuEndapDarah:            rw.LajuEndapDarah,
			Albumin:                   rw.Albumin,
			Visible:                   rw.Visible,
		})
	}
	return res, nil
}
