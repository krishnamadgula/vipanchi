package types

const (
	LaghuTisra     = 3
	LaghuChatusra  = 4
	LaghuKhanda    = 5
	LaghuMisra     = 7
	LaghuSankeerna = 9
	Dhrutam        = 2
	AnuDhrutam     = 1

	Kaala1st = 1
	Kaala2nd = 2
	Kaala3rd = 4
)

type Taalam struct {
	// TODO: revisit, having tempo at taalam level.
	Name string
	// Sequence consists of Laghu, Dhrutam and Anudhrutam
	Sequence []int
}
