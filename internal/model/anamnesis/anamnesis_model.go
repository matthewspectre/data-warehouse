package anamnesis

import "time"

type AnamnesisModel struct {
	ID                    int       `gorm:"primaryKey;autoIncrement;column:id_anamnesis"`
	IDPasien              int       `gorm:"column:id_pasien"`
	IDDokter              int       `gorm:"column:id_dokter"`
	Text                  string    `gorm:"column:text"`
	DateMake              time.Time `gorm:"column:date_make"`
	DateUpdate            time.Time `gorm:"column:date_update"`
	IDDataKlinik          int       `gorm:"column:id_data_klinik"`
	RiwayatPengobatan     string    `gorm:"column:riwayat_pengobatan"`
	RiwayatKeluarga       string    `gorm:"column:riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `gorm:"column:riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `gorm:"column:riwayat_penyakit_lain"`
	StatusKehamilan       string    `gorm:"column:status_kehamilan"`
	KeluhanTambahan       string    `gorm:"column:keluhan_tambahan"`

	Visible int `gorm:"column:visible"`
}

func (AnamnesisModel) TableName() string {
	return "anamnesis"
}

// AnamnesisBModel maps to the rme-b `anamnesis` table structure.
type AnamnesisBModel struct {
	ID                    int       `gorm:"primaryKey;autoIncrement;column:id_anamnesis"`
	IDPasien              int       `gorm:"column:id_pasien"`
	IDDokter              int       `gorm:"column:id_dokter"`
	Text                  string    `gorm:"column:text"`
	DateMake              time.Time `gorm:"column:date_make"`
	DateUpdate            time.Time `gorm:"column:date_update"`
	IDDataKlinik          int       `gorm:"column:id_data_klinik"`
	RiwayatPengobatan     string    `gorm:"column:riwayat_pengobatan"`
	RiwayatKeluarga       string    `gorm:"column:riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `gorm:"column:riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `gorm:"column:riwayat_penyakit_lain"`
	StatusKehamilan       string    `gorm:"column:status_kehamilan"`
	KeluhanTambahan       string    `gorm:"column:keluhan_tambahan"`
	KeluhanUtama          string    `gorm:"column:keluhan_utama"`
	RiwayatAlergi         string    `gorm:"column:riwayat_alergi"`
	Visible               int       `gorm:"column:visible"`
}

func (AnamnesisBModel) TableName() string {
	return "anamnesis"
}

// AnamnesisWarehouseModel maps to the `data_warehouse`.anamnesis table.
// It stores data from both rsA and rsB with a `source` marker.
type AnamnesisWarehouseModel struct {
	Source                string    `gorm:"column:source"`
	IDAnamnesis           int       `gorm:"column:id_anamnesis"`
	IDPasien              int       `gorm:"column:id_pasien"`
	NamaPasien            *string   `gorm:"column:nama_pasien"`
	Text                  string    `gorm:"column:text"`
	DateMake              time.Time `gorm:"column:date_make"`
	DateUpdate            time.Time `gorm:"column:date_update"`
	IDDataKlinik          int       `gorm:"column:id_data_klinik"`
	RiwayatPengobatan     string    `gorm:"column:riwayat_pengobatan"`
	RiwayatKeluarga       string    `gorm:"column:riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `gorm:"column:riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `gorm:"column:riwayat_penyakit_lain"`
	RiwayatAlergi         *string   `gorm:"column:riwayat_alergi"`
	StatusKehamilan       string    `gorm:"column:status_kehamilan"`
	KeluhanUtama          *string   `gorm:"column:keluhan_utama"`
	KeluhanTambahan       string    `gorm:"column:keluhan_tambahan"`
	Visible               int       `gorm:"column:visible"`
}

func (AnamnesisWarehouseModel) TableName() string {
	// NOTE: Use backticks because it will be used as-is by GORM.
	return "`data_warehouse`.anamnesis"
}
