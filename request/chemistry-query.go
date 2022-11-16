package request

type GetChemistryReq struct {
	TypeMaterial string `json:"typeMaterial" form:"typeMaterial"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type DeleteChemistryReq struct {
	TypeMaterial string `json:"typeMaterial" form:"typeMaterial"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type InsertChemistryReq struct {
	TypeMaterial string `json:"typeMaterial"`
	TypeSpectrum string `json:"typeSpectrum"`
	Chemical     string `json:"chemical"`
	HTMLText     string `json:"htmlText,omitempty"`
	VideoUrl     string `json:"videoUrl,omitempty"`
}

type GetChemistryListBySpectrum struct {
	TypeMaterial string `json:"typeMaterial"`
	TypeSpectrum string `json:"typeSpectrum"`
}
