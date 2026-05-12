package pemeriksaan_fungsi_organ

import "time"

type PemeriksaanFungsiOrgan struct {
	ID            int
	IDPasien      int
	IDDokter      int
	Tanggal       time.Time
	DateMake      time.Time
	DateUpdate    time.Time
	GangguanBAB   bool
	GangguanBAK   bool
	MualMuntah    bool
	Demam         bool
	Perdarahan    bool
	PenurunanBB   bool
	GangguanGerak bool
	Nyeri         bool
	SesakNapas    bool
	Pusing        bool
	Catatan       string
	Visible       int
	NamaPasien    string
}

type PemeriksaanFungsiOrganWarehouse struct {
	Source                   string
	IDPemeriksaanFungsiOrgan int
	IDPasien                 int
	NamaPasien               *string
	IDDokter                 int
	Tanggal                  time.Time
	DateMake                 time.Time
	DateUpdate               time.Time
	GangguanBAB              bool
	GangguanBAK              bool
	MualMuntah               bool
	Demam                    bool
	Perdarahan               bool
	PenurunanBB              bool
	GangguanGerak            bool
	Nyeri                    bool
	SesakNapas               bool
	Pusing                   bool
	Catatan                  string
	Visible                  int
}
