package pemeriksaan_fungsi_organ

type PemeriksaanFungsiOrganCreateRequest struct {
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	Tanggal       string `json:"tanggal"` // RFC3339 optional
	GangguanBAB   bool   `json:"gangguanBAB"`
	GangguanBAK   bool   `json:"gangguanBAK"`
	MualMuntah    bool   `json:"mualMuntah"`
	Demam         bool   `json:"demam"`
	Perdarahan    bool   `json:"perdarahan"`
	PenurunanBB   bool   `json:"penurunanBB"`
	GangguanGerak bool   `json:"gangguanGerak"`
	Nyeri         bool   `json:"nyeri"`
	SesakNapas    bool   `json:"sesakNapas"`
	Pusing        bool   `json:"pusing"`
	Catatan       string `json:"catatan"`
}

type PemeriksaanFungsiOrganResponse struct {
	ID            int    `json:"id"`
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	Tanggal       string `json:"tanggal"`
	GangguanBAB   bool   `json:"gangguanBAB"`
	GangguanBAK   bool   `json:"gangguanBAK"`
	MualMuntah    bool   `json:"mualMuntah"`
	Demam         bool   `json:"demam"`
	Perdarahan    bool   `json:"perdarahan"`
	PenurunanBB   bool   `json:"penurunanBB"`
	GangguanGerak bool   `json:"gangguanGerak"`
	Nyeri         bool   `json:"nyeri"`
	SesakNapas    bool   `json:"sesakNapas"`
	Pusing        bool   `json:"pusing"`
	Catatan       string `json:"catatan"`
}

type PemeriksaanFungsiOrganBResponse struct {
	Source        string `json:"source"`
	ID            int    `json:"id"`
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	Tanggal       string `json:"tanggal"`
	GangguanBAB   bool   `json:"gangguanBAB"`
	GangguanBAK   bool   `json:"gangguanBAK"`
	MualMuntah    bool   `json:"mualMuntah"`
	Demam         bool   `json:"demam"`
	Perdarahan    bool   `json:"perdarahan"`
	PenurunanBB   bool   `json:"penurunanBB"`
	GangguanGerak bool   `json:"gangguanGerak"`
	Nyeri         bool   `json:"nyeri"`
	SesakNapas    bool   `json:"sesakNapas"`
	Pusing        bool   `json:"pusing"`
	Catatan       string `json:"catatan"`
}

type PemeriksaanFungsiOrganWarehouseResponse struct {
	Source                   string  `json:"source"`
	IDPemeriksaanFungsiOrgan int     `json:"id_pemeriksaan_fungsi_organ"`
	IDPasien                 int     `json:"id_pasien"`
	NamaPasien               *string `json:"nama_pasien"`
	IDDokter                 int     `json:"id_dokter"`
	Tanggal                  string  `json:"tanggal"`
	DateMake                 string  `json:"date_make"`
	DateUpdate               string  `json:"date_update"`
	GangguanBAB              bool    `json:"gangguanBAB"`
	GangguanBAK              bool    `json:"gangguanBAK"`
	MualMuntah               bool    `json:"mualMuntah"`
	Demam                    bool    `json:"demam"`
	Perdarahan               bool    `json:"perdarahan"`
	PenurunanBB              bool    `json:"penurunanBB"`
	GangguanGerak            bool    `json:"gangguanGerak"`
	Nyeri                    bool    `json:"nyeri"`
	SesakNapas               bool    `json:"sesakNapas"`
	Pusing                   bool    `json:"pusing"`
	Catatan                  string  `json:"catatan"`
	Visible                  int     `json:"visible"`
}

type GetWarehouseResponse struct {
	Data []PemeriksaanFungsiOrganWarehouseResponse `json:"data"`
}
