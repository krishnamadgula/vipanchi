package types

import "errors"

type Raagam struct {
	Name string
	// TODO need to figure how to define bhaavam - variation in swaras
	Aarohana  SwaraTransitions
	Avarohana SwaraTransitions
}

type Swara struct {
	SwaraSthana float32
	// Gamakam is a function such as an oscilation / linear interpolation from
	// one frequency to another.
	// TODO more work around Gamaka
	// Gamakam func(SwaraSthana float32) []float32
}

// SwaraTransition is the structural representation that tells us the number
// of steps to traverse from one swara to another.
// FIXME: This doesn't work for Vakra raagams, loops are not handled!
type SwaraTransitions map[string]map[int]string

func NewRaagam(name string, aarohanam, avarohanam []string) *Raagam {
	r := Raagam{Name: name}
	r.Aarohana = constructTransition(aarohanam)
	r.Avarohana = constructTransition(avarohanam)
	return &r
}

func NewSimpleRaagam(name string, swaraSthanas []string) *Raagam {
	avarohanam := make([]string, len(swaraSthanas))

	for i := range swaraSthanas {
		avarohanam[i] = swaraSthanas[len(swaraSthanas)-i-1]
	}
	return NewRaagam(name, swaraSthanas, avarohanam)
}

func constructTransition(swaras []string) SwaraTransitions {
	m := SwaraTransitions{}
	for i := 0; i < len(swaras)-1; i++ {
		if _, ok := m[swaras[i]]; !ok {
			m[swaras[i]] = make(map[int]string)
		}

		for j := i + 1; j < len(swaras); j++ {
			m[swaras[i]][j-i] = swaras[j]
		}
	}

	return m
}

func (r *Raagam) Aaroh(curSwara string, step int) (string, error) {
	curNote, ok := r.Aarohana[curSwara][step]
	if !ok {
		return "", errors.New("higher than current octave / sthayi")
	}

	return curNote, nil
}

func (r *Raagam) Avroh(curSwara string, step int) (string, error) {
	curNote, ok := r.Avarohana[curSwara][step]
	if !ok {
		return "", errors.New("lower than current octave / sthayi")
	}

	return curNote, nil
}

func (r *Raagam) AarohOne(curSwara string) (string, error) {
	// TODO assuming curSwara always in raga
	return r.Aaroh(curSwara, 1)

}

func (r *Raagam) AvrohOne(curSwara string) (string, error) {
	return r.Avroh(curSwara, 1)
}
