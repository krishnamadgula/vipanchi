package types

const (
	CScaleFreq = 261.63
	CScale     = "C"
	AScaleFreq = 440
	AScale     = "A"
)

var Scales = map[string]float32{
	CScale: CScaleFreq,
	AScale: AScaleFreq,
}


