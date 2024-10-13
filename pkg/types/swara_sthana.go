package types

/*
S (C4) – 261.63 Hz
R₁ (D♭4) – 277.18 Hz
G₃ (E4) – 329.63 Hz
M₁ (F4) – 349.23 Hz
P (G4) – 392.00 Hz
D₁ (A♭4) – 415.30 Hz
N₃ (B4) – 493.88 Hz
*/
const (
	CScale = 261.63
	AScale = 440
)

const (
	SLower = float32(261.63)
	R1     = float32(277.18)
	R2     = float32(293.66)

	G2 = float32(311.13)
	G3 = float32(329.63)

	M1 = float32(349.23)
	M2 = float32(369.99)

	P  = float32(392)
	D1 = float32(415.30)
	D2 = float32(440)

	N2     = float32(466.16)
	N3     = float32(493.88)
	SUpper = float32(523.25)
)
