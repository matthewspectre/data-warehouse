package anamnesis

// Request/response DTOs for anamnesis handlers.

import "time"

// AnamnesisResponse merepresentasikan data anamnesis yang dikembalikan ke client.
type AnamnesisResponse struct {
	Source                string    `json:"source"`
	ID                    int       `json:"id_anamnesis"`
	IDPasien              int       `json:"id_pasien"`
	IDDokter              int       `json:"id_dokter"`
	NamaPasien            string    `json:"nama_pasien"`
	Text                  string    `json:"text"`
	DateMake              time.Time `json:"date_make"`
	DateUpdate            time.Time `json:"date_update"`
	IDDataKlinik          int       `json:"id_data_klinik"`
	RiwayatPengobatan     string    `json:"riwayat_pengobatan"`
	RiwayatKeluarga       string    `json:"riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `json:"riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `json:"riwayat_penyakit_lain"`
	StatusKehamilan       string    `json:"status_kehamilan"`
	KeluhanTambahan       string    `json:"keluhan_tambahan"`
	Visible               int       `json:"visible"`
}

// AnamnesisBResponse merepresentasikan data anamnesis dari sumber RME-B.
type AnamnesisBResponse struct {
	Source                string    `json:"source"`
	ID                    int       `json:"id_anamnesis"`
	IDPasien              int       `json:"id_pasien"`
	IDDokter              int       `json:"id_dokter"`
	NamaPasien            string    `json:"nama_pasien"`
	Text                  string    `json:"text"`
	DateMake              time.Time `json:"date_make"`
	DateUpdate            time.Time `json:"date_update"`
	IDDataKlinik          int       `json:"id_data_klinik"`
	RiwayatPengobatan     string    `json:"riwayat_pengobatan"`
	RiwayatKeluarga       string    `json:"riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `json:"riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `json:"riwayat_penyakit_lain"`
	RiwayatAlergi         string    `json:"riwayat_alergi"`
	StatusKehamilan       string    `json:"status_kehamilan"`
	KeluhanUtama          string    `json:"keluhan_utama"`
	KeluhanTambahan       string    `json:"keluhan_tambahan"`
	Visible               int       `json:"visible"`
}

// GetAllResponse menyatukan response anamnesis dari rsA dan rsB.
type GetAllResponse struct {
	RSA []AnamnesisResponse  `json:"rs_a"`
	RSB []AnamnesisBResponse `json:"rs_b"`
}

// AnamnesisWarehouseResponse merepresentasikan baris dari data_warehouse.anamnesis.
type AnamnesisWarehouseResponse struct {
	Source                string    `json:"source"`
	IDAnamnesis           int       `json:"id_anamnesis"`
	IDPasien              int       `json:"id_pasien"`
	NamaPasien            *string   `json:"nama_pasien"`
	Text                  string    `json:"text"`
	DateMake              time.Time `json:"date_make"`
	DateUpdate            time.Time `json:"date_update"`
	IDDataKlinik          int       `json:"id_data_klinik"`
	RiwayatPengobatan     string    `json:"riwayat_pengobatan"`
	RiwayatKeluarga       string    `json:"riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `json:"riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `json:"riwayat_penyakit_lain"`
	RiwayatAlergi         *string   `json:"riwayat_alergi"`
	StatusKehamilan       string    `json:"status_kehamilan"`
	KeluhanUtama          *string   `json:"keluhan_utama"`
	KeluhanTambahan       string    `json:"keluhan_tambahan"`
	Visible               int       `json:"visible"`
}

type GetWarehouseResponse struct {
	Data []AnamnesisWarehouseResponse `json:"data"`
}
