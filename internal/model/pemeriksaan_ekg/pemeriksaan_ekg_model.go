package pemeriksaan_ekg

import "time"

type PemeriksaanEkgModel struct {
	ID                 int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien           int       `gorm:"column:id_pasien"`
	IDDokter           int       `gorm:"column:id_dokter"`
	DetakJantung       *int      `gorm:"column:detak_jantung"`
	Irama              string    `gorm:"column:irama"`
	PRInterval         *float64  `gorm:"column:pr_interval"`
	QRSDuration        *float64  `gorm:"column:qrs_duration"`
	QTTcInterval       *float64  `gorm:"column:qt_qtc_interval"`
	AxisJantung        string    `gorm:"column:axis_jantung"`
	STElevationDepress string    `gorm:"column:st_elevation_depression"`
	TWaveAbnormality   string    `gorm:"column:t_wave_abnormality"`
	InterpretasiDokter string    `gorm:"column:interpretasi_dokter"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	Visible            int       `gorm:"column:visible"`
}

func (PemeriksaanEkgModel) TableName() string {
	return "pemeriksaan_ekg"
}

// PemeriksaanEkgBModel maps to the rme-b `pemeriksaan_ekg` table structure.
type PemeriksaanEkgBModel struct {
	ID                 int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien           int       `gorm:"column:id_pasien"`
	IDDokter           int       `gorm:"column:id_dokter"`
	DetakJantung       *int      `gorm:"column:detak_jantung"`
	Irama              string    `gorm:"column:irama"`
	PRInterval         *float64  `gorm:"column:pr_interval"`
	QRSDuration        *float64  `gorm:"column:qrs_duration"`
	QTTcInterval       *float64  `gorm:"column:qt_qtc_interval"`
	PWave              string    `gorm:"column:p_wave"`
	AVBlock            string    `gorm:"column:av_block"`
	InterpretasiDokter string    `gorm:"column:interpretasi_dokter"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	Visible            int       `gorm:"column:visible"`
}

func (PemeriksaanEkgBModel) TableName() string {
	return "pemeriksaan_ekg"
}

// PemeriksaanEkgWarehouseModel maps to the `data_warehouse`.pemeriksaan_ekg table.
// It stores data from both rsA and rsB with a `source` marker.
type PemeriksaanEkgWarehouseModel struct {
	Source           string  `gorm:"column:source"`
	IDPemeriksaanEkg int     `gorm:"column:id_pemeriksaan_ekg"`
	IDPasien         int     `gorm:"column:id_pasien"`
	NamaPasien       *string `gorm:"column:nama_pasien"`

	DetakJantung *int     `gorm:"column:detak_jantung"`
	Irama        string   `gorm:"column:irama"`
	PRInterval   *float64 `gorm:"column:pr_interval"`
	QRSDuration  *float64 `gorm:"column:qrs_duration"`
	QTTcInterval *float64 `gorm:"column:qt_qtc_interval"`

	// Kolom yang ada di rsA
	AxisJantung        string `gorm:"column:axis_jantung"`
	STElevationDepress string `gorm:"column:st_elevation_depression"`
	TWaveAbnormality   string `gorm:"column:t_wave_abnormality"`

	// Kolom yang ada di rsB
	PWave   string `gorm:"column:p_wave"`
	AVBlock string `gorm:"column:av_block"`

	InterpretasiDokter string `gorm:"column:interpretasi_dokter"`
	DateMake           string `gorm:"column:date_make"`
	DateUpdate         string `gorm:"column:date_update"`
	Visible            int    `gorm:"column:visible"`
}

func (PemeriksaanEkgWarehouseModel) TableName() string {
	return "`data_warehouse`.pemeriksaan_ekg"
}
