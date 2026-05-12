package pemeriksaan_fungsi_organ

import "time"

type PemeriksaanFungsiOrganModel struct {
	ID            int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien      int       `gorm:"column:id_pasien"`
	IDDokter      int       `gorm:"column:id_dokter"`
	Tanggal       time.Time `gorm:"column:tanggal"`
	GangguanBAB   bool      `gorm:"column:gangguan_bab"`
	GangguanBAK   bool      `gorm:"column:gangguan_bak"`
	MualMuntah    bool      `gorm:"column:mual_muntah"`
	Demam         bool      `gorm:"column:demam"`
	Perdarahan    bool      `gorm:"column:perdarahan"`
	PenurunanBB   bool      `gorm:"column:penurunan_bb"`
	GangguanGerak bool      `gorm:"column:gangguan_gerak"`
	Nyeri         bool      `gorm:"column:nyeri"`
	SesakNapas    bool      `gorm:"column:sesak_napas"`
	Pusing        bool      `gorm:"column:pusing"`
	Catatan       string    `gorm:"column:catatan"`
	Visible       int       `gorm:"column:visible"`
	DateMake      time.Time `gorm:"column:date_make"`
	DateUpdate    time.Time `gorm:"column:date_update"`
}

func (PemeriksaanFungsiOrganModel) TableName() string {
	return "pemeriksaan_fungsi_organ"
}

// PemeriksaanFungsiOrganBModel maps to the rme-b pemeriksaan_fungsi_organ table.
type PemeriksaanFungsiOrganBModel struct {
	ID            int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien      int       `gorm:"column:id_pasien"`
	IDDokter      int       `gorm:"column:id_dokter"`
	Tanggal       time.Time `gorm:"column:tanggal"`
	GangguanBAB   bool      `gorm:"column:gangguan_bab"`
	GangguanBAK   bool      `gorm:"column:gangguan_bak"`
	MualMuntah    bool      `gorm:"column:mual_muntah"`
	Demam         bool      `gorm:"column:demam"`
	Perdarahan    bool      `gorm:"column:perdarahan"`
	PenurunanBB   bool      `gorm:"column:penurunan_bb"`
	GangguanGerak bool      `gorm:"column:gangguan_gerak"`
	Nyeri         bool      `gorm:"column:nyeri"`
	SesakNapas    bool      `gorm:"column:sesak_napas"`
	Pusing        bool      `gorm:"column:pusing"`
	Catatan       string    `gorm:"column:catatan"`
	Visible       int       `gorm:"column:visible"`
	DateMake      time.Time `gorm:"column:date_make"`
	DateUpdate    time.Time `gorm:"column:date_update"`
}

func (PemeriksaanFungsiOrganBModel) TableName() string {
	return "pemeriksaan_fungsi_organ"
}

// PemeriksaanFungsiOrganWarehouseModel maps to the `data_warehouse`.pemeriksaan_fungsi_organ table.
// It stores data from both rsA and rsB with a `source` marker.
type PemeriksaanFungsiOrganWarehouseModel struct {
	Source                   string    `gorm:"column:source"`
	IDPemeriksaanFungsiOrgan int       `gorm:"column:id_pemeriksaan_fungsi_organ"`
	IDPasien                 int       `gorm:"column:id_pasien"`
	NamaPasien               *string   `gorm:"column:nama_pasien"`
	IDDokter                 int       `gorm:"column:id_dokter"`
	Tanggal                  time.Time `gorm:"column:tanggal"`
	GangguanBAB              bool      `gorm:"column:gangguan_bab"`
	GangguanBAK              bool      `gorm:"column:gangguan_bak"`
	MualMuntah               bool      `gorm:"column:mual_muntah"`
	Demam                    bool      `gorm:"column:demam"`
	Perdarahan               bool      `gorm:"column:perdarahan"`
	PenurunanBB              bool      `gorm:"column:penurunan_bb"`
	GangguanGerak            bool      `gorm:"column:gangguan_gerak"`
	Nyeri                    bool      `gorm:"column:nyeri"`
	SesakNapas               bool      `gorm:"column:sesak_napas"`
	Pusing                   bool      `gorm:"column:pusing"`
	Catatan                  string    `gorm:"column:catatan"`
	Visible                  int       `gorm:"column:visible"`
	DateMake                 time.Time `gorm:"column:date_make"`
	DateUpdate               time.Time `gorm:"column:date_update"`
}

func (PemeriksaanFungsiOrganWarehouseModel) TableName() string {
	return "`data_warehouse`.pemeriksaan_fungsi_organ"
}
