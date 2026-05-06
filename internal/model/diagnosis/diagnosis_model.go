package diagnosis

import "time"

type DiagnosisModel struct {
	IDDiagnosis      int       `gorm:"primaryKey;autoIncrement;column:id_diagnosis"`
	IDPasien         int       `gorm:"column:id_pasien"`
	IDDokter         int       `gorm:"column:id_dokter"`
	Tanggal          time.Time `gorm:"column:tanggal"`
	KodeIcdUtama     string    `gorm:"column:kode_icd_utama"`
	KodeIcdSekunder  string    `gorm:"column:kode_icd_sekunder"` // JSON-encoded array
	DiagnosisBanding string    `gorm:"column:diagnosis_banding"` // JSON-encoded array
	Status           string    `gorm:"column:status"`
	DasarDiagnosis   string    `gorm:"column:dasar_diagnosis"` // JSON-encoded array
	Catatan          string    `gorm:"column:catatan"`
	Visible          int       `gorm:"column:visible"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
}

func (DiagnosisModel) TableName() string {
	return "diagnosis"
}

// DiagnosisBModel maps to the rme-b `diagnosis` table structure.
type DiagnosisBModel struct {
	IDDiagnosis      int       `gorm:"primaryKey;autoIncrement;column:id_diagnosis"`
	IDPasien         int       `gorm:"column:id_pasien"`
	IDDokter         int       `gorm:"column:id_dokter"`
	Tanggal          time.Time `gorm:"column:tanggal"`
	KodeIcdUtama     string    `gorm:"column:kode_icd_utama"`
	KodeIcdSekunder  string    `gorm:"column:kode_icd_sekunder"`
	DiagnosisBanding string    `gorm:"column:diagnosis_banding"`
	Status           string    `gorm:"column:status"`
	DasarDiagnosis   string    `gorm:"column:dasar_diagnosis"`
	Catatan          string    `gorm:"column:catatan"`
	Visible          int       `gorm:"column:visible"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
}

func (DiagnosisBModel) TableName() string {
	return "diagnosis"
}

// DiagnosisWarehouseModel maps to the `data_warehouse`.diagnosis table.
// It stores data from both rsA and rsB with a `source` marker.
type DiagnosisWarehouseModel struct {
	Source      string  `gorm:"column:source"`
	IDDiagnosis int     `gorm:"column:id_diagnosis"`
	IDPasien    int     `gorm:"column:id_pasien"`
	IDDokter    int     `gorm:"column:id_dokter"`
	NamaPasien  *string `gorm:"column:nama_pasien"`

	Tanggal          time.Time `gorm:"column:tanggal"`
	KodeIcdUtama     string    `gorm:"column:kode_icd_utama"`
	KodeIcdSekunder  string    `gorm:"column:kode_icd_sekunder"`
	DiagnosisBanding string    `gorm:"column:diagnosis_banding"`
	Status           string    `gorm:"column:status"`
	DasarDiagnosis   string    `gorm:"column:dasar_diagnosis"`
	Catatan          string    `gorm:"column:catatan"`
	Visible          int       `gorm:"column:visible"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
}

func (DiagnosisWarehouseModel) TableName() string {
	return "`data_warehouse`.diagnosis"
}
