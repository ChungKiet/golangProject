package enum

// init 2 type enum of type material and type spectrum
type TypeMaterialValue string
type TypeMaterialEnv struct {
	REFERENCE      TypeMaterialValue
	SOLUTION_GUIDE TypeMaterialValue
}

var TypeMaterial = &TypeMaterialEnv{
	"REFERENCE",
	"SOLUTION_GUIDE",
}

type SpectrumValue string
type SpectrumEnv struct {
	IS  SpectrumValue
	MS  SpectrumValue
	MRS SpectrumValue
}

var Spectrum = &SpectrumEnv{
	"IS",
	"MS",
	"MRS",
}
