package request

type GetChemistryReq struct {
	TypeChemical string `json:"typeChemical" form:"typeChemical"`
	GroupName    string `json:"groupName" form:"groupName"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type DeleteChemistryReq struct {
	TypeChemical string `json:"typeChemical" form:"typeChemical"`
	GroupName    string `json:"groupName" form:"groupName"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type InsertChemistryReq struct {
	TypeChemical string `json:"typeChemical,omitempty"`
	GroupName    string `json:"groupName,omitempty"`
	TypeSpectrum string `json:"typeSpectrum,omitempty"`
	Chemical     string `json:"chemical,omitempty"`
	VideoUrl     string `json:"videoUrl,omitempty"`
}

type UpdateChemistryReq struct {
	TypeChemical string `json:"typeChemical,omitempty"`
	GroupName    string `json:"groupName,omitempty"`
	TypeSpectrum string `json:"typeSpectrum,omitempty"`
	Chemical     string `json:"chemical,omitempty"`
	VideoUrl     string `json:"videoUrl,omitempty"`
}

type GetRefDocument struct {
	Type string `json:"type"`
}

type ImportRefDocument struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type UpdateRefDocument struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type DeleteRefDocument struct {
	Type string `json:"type"`
}
