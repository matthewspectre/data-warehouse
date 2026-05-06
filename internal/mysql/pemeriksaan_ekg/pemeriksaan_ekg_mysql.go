package pemeriksaan_ekg

import (
	"context"
	"errors"

	model "rme/internal/model/pemeriksaan_ekg"
	repo "rme/internal/repository/pemeriksaan_ekg"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error) {
	ctx := context.Background()
	var rows []repo.PemeriksaanEkgWithNames
	q := r.db.WithContext(ctx).Table("pemeriksaan_ekg as pe").
		Select("pe.id, pe.id_pasien, p.name as nama_pasien, pe.id_dokter, d.nama_dokter as nama_dokter, pe.detak_jantung, pe.irama, pe.pr_interval, pe.qrs_duration, pe.qt_qtc_interval AS qt_tc_interval, pe.axis_jantung, pe.st_elevation_depression AS st_elevation_depress, pe.t_wave_abnormality, pe.interpretasi_dokter, pe.visible, pe.date_make").
		Joins("JOIN patients p ON pe.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pe.id_dokter = d.id").
		Where("pe.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pe.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pe.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanEkgWithNames, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func (r *RepositoryMySQL) GetAllB(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNamesB, error) {
	ctx := context.Background()
	var rows []repo.PemeriksaanEkgWithNamesB
	q := r.db.WithContext(ctx).Table("`rme-b`.pemeriksaan_ekg as pe").
		Select("pe.id, pe.id_pasien, p.name as nama_pasien, pe.id_dokter, d.nama_dokter as nama_dokter, pe.detak_jantung, pe.irama, pe.pr_interval, pe.qrs_duration, pe.qt_qtc_interval AS qt_tc_interval, pe.p_wave, pe.av_block, pe.interpretasi_dokter, pe.visible, pe.date_make, pe.date_update").
		Joins("JOIN `rme-b`.patients p ON pe.id_pasien = p.id").
		Joins("LEFT JOIN `rme-b`.dokter d ON pe.id_dokter = d.id").
		Where("pe.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pe.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pe.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanEkgWithNamesB, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func toWarehouseModelA(rw *repo.PemeriksaanEkgWithNames) *model.PemeriksaanEkgWarehouseModel {
	if rw == nil {
		return nil
	}
	var nama *string
	if rw.NamaPasien != "" {
		n := rw.NamaPasien
		nama = &n
	}
	return &model.PemeriksaanEkgWarehouseModel{
		Source:             "rsA",
		IDPemeriksaanEkg:   rw.ID,
		IDPasien:           rw.IDPasien,
		NamaPasien:         nama,
		DetakJantung:       rw.DetakJantung,
		Irama:              rw.Irama,
		PRInterval:         rw.PRInterval,
		QRSDuration:        rw.QRSDuration,
		QTTcInterval:       rw.QTTcInterval,
		AxisJantung:        rw.AxisJantung,
		STElevationDepress: rw.STElevationDepress,
		TWaveAbnormality:   rw.TWaveAbnormality,
		PWave:              "",
		AVBlock:            "",
		InterpretasiDokter: rw.InterpretasiDokter,
		DateMake:           rw.DateMake,
		DateUpdate:         "",
		Visible:            rw.Visible,
	}
}

func toWarehouseModelB(rw *repo.PemeriksaanEkgWithNamesB) *model.PemeriksaanEkgWarehouseModel {
	if rw == nil {
		return nil
	}
	var nama *string
	if rw.NamaPasien != "" {
		n := rw.NamaPasien
		nama = &n
	}
	return &model.PemeriksaanEkgWarehouseModel{
		Source:             "rsB",
		IDPemeriksaanEkg:   rw.ID,
		IDPasien:           rw.IDPasien,
		NamaPasien:         nama,
		DetakJantung:       rw.DetakJantung,
		Irama:              rw.Irama,
		PRInterval:         rw.PRInterval,
		QRSDuration:        rw.QRSDuration,
		QTTcInterval:       rw.QTTcInterval,
		AxisJantung:        "",
		STElevationDepress: "",
		TWaveAbnormality:   "",
		PWave:              rw.PWave,
		AVBlock:            rw.AVBlock,
		InterpretasiDokter: rw.InterpretasiDokter,
		DateMake:           rw.DateMake,
		DateUpdate:         rw.DateUpdate,
		Visible:            rw.Visible,
	}
}

func (r *RepositoryMySQL) upsertWarehouse(rows []*model.PemeriksaanEkgWarehouseModel) (int64, error) {
	if len(rows) == 0 {
		return 0, nil
	}
	ctx := context.Background()

	conflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "id_pemeriksaan_ekg"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"id_pasien",
			"nama_pasien",
			"detak_jantung",
			"irama",
			"pr_interval",
			"qrs_duration",
			"qt_qtc_interval",
			"axis_jantung",
			"st_elevation_depression",
			"t_wave_abnormality",
			"p_wave",
			"av_block",
			"interpretasi_dokter",
			"date_make",
			"date_update",
			"visible",
		}),
	}

	err := r.db.WithContext(ctx).Clauses(conflict).CreateInBatches(rows, 500).Error
	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil
}

func (r *RepositoryMySQL) UpsertWarehouseA(rows []*repo.PemeriksaanEkgWithNames) (int64, error) {
	mdls := make([]*model.PemeriksaanEkgWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelA(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) UpsertWarehouseB(rows []*repo.PemeriksaanEkgWithNamesB) (int64, error) {
	mdls := make([]*model.PemeriksaanEkgWarehouseModel, 0, len(rows))
	for _, e := range rows {
		m := toWarehouseModelB(e)
		if m != nil {
			mdls = append(mdls, m)
		}
	}
	return r.upsertWarehouse(mdls)
}

func (r *RepositoryMySQL) GetAllWarehouse(nik *string) ([]*repo.PemeriksaanEkgWarehouse, error) {
	ctx := context.Background()
	var rows []model.PemeriksaanEkgWarehouseModel

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
			return []*repo.PemeriksaanEkgWarehouse{}, nil
		}
	}

	q := r.db.WithContext(ctx).
		Model(&model.PemeriksaanEkgWarehouseModel{}).
		Where("visible = ?", 1).
		Order("id_pemeriksaan_ekg DESC")

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

	res := make([]*repo.PemeriksaanEkgWarehouse, 0, len(rows))
	for i := range rows {
		rw := &rows[i]
		res = append(res, &repo.PemeriksaanEkgWarehouse{
			Source:             rw.Source,
			IDPemeriksaanEkg:   rw.IDPemeriksaanEkg,
			IDPasien:           rw.IDPasien,
			NamaPasien:         rw.NamaPasien,
			Visible:            rw.Visible,
			DetakJantung:       rw.DetakJantung,
			Irama:              rw.Irama,
			PRInterval:         rw.PRInterval,
			QRSDuration:        rw.QRSDuration,
			QTTcInterval:       rw.QTTcInterval,
			AxisJantung:        rw.AxisJantung,
			STElevationDepress: rw.STElevationDepress,
			TWaveAbnormality:   rw.TWaveAbnormality,
			PWave:              rw.PWave,
			AVBlock:            rw.AVBlock,
			InterpretasiDokter: rw.InterpretasiDokter,
			DateMake:           rw.DateMake,
			DateUpdate:         rw.DateUpdate,
		})
	}
	return res, nil
}
