package models

type KodeWilayah struct {
	Provinsi  []string `json:"provinsi"`
	Kabupaten []string `json:"kabupaten"`
	Kecamatan []string `json:"kecamatan"`
}
