package pemeriksaanvital

import "time"

// Domain entity untuk pemeriksaan vital.
type PemeriksaanVital struct {
	ID             int
	IDPasien       int
	IDDokter       int
	DateMake       time.Time
	DateUpdate     time.Time
	TekananDarah   string
	DenyutNadi     int
	SuhuTubuh      float64
	FrekuensiNapas int
	BeratBadan     float64
	TinggiBadan    float64
	Visible        int
	NamaPasien     string
	NamaDokter     string
}

type PemeriksaanVitalB struct {
	ID               int
	IDPasien         int
	IDDokter         int
	DateMake         time.Time
	DateUpdate       time.Time
	TekananDarah     string
	DenyutNadi       int
	SuhuTubuh        float64
	FrekuensiNapas   int
	BeratBadan       float64
	TinggiBadan      float64
	SaturasiOksigen  int
	TingkatKesadaran string
	IndeksMasaTubuh  float64
	Visible          int
	NamaPasien       string
	NamaDokter       string
}

type PemeriksaanVitalWarehouse struct {
	Source             string
	IDPemeriksaanVital int
	IDPasien           int
	NamaPasien         *string
	DateMake           time.Time
	DateUpdate         time.Time
	TekananDarah       string
	DenyutNadi         int
	SuhuTubuh          float64
	FrekuensiNapas     int
	BeratBadan         float64
	TinggiBadan        float64
	SaturasiOksigen    *int
	TingkatKesadaran   *string
	IndeksMasaTubuh    *float64
	Visible            int
}
