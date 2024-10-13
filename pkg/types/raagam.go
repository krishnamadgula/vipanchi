package types

import "errors"

type Raagam struct {
	Name        string
	swaras      *Swara
	swaraLookup map[float32]*Swara
	// TODO need to figure how to define bhaavam - variation in swaras
}

// TODO not handling vakra ragas / arohanam and avrohanam variations also
type Swara struct {
	avroh     *Swara
	aaroh     *Swara
	frequency float32
}

func newSwara(swaraSthana float32) *Swara {
	return &Swara{frequency: swaraSthana}
}

func NewSimpleRaagam(name string, swaraSthanas []float32) *Raagam {
	lookup := map[float32]*Swara{}
	start := newSwara(swaraSthanas[0])
	lookup[swaraSthanas[0]] = start
	cur := start
	for i := 1; i < len(swaraSthanas); i++ {
		note := newSwara(swaraSthanas[i])
		cur.aaroh = note
		note.avroh = cur
		cur = note
		lookup[swaraSthanas[i]] = note
	}

	return &Raagam{swaras: start, Name: name, swaraLookup: lookup}
}

func (r *Raagam) Aaroh(curSwara float32) (float32, error) {
	// TODO assuming curSwara always in raga
	curNote := r.swaraLookup[curSwara]
	if curNote.aaroh == nil {
		return 0, errors.New("higher than current octave / sthayi")
	}

	return curNote.aaroh.frequency, nil

}

func (r *Raagam) Avroh(curSwara float32) (float32, error) {
	// TODO assuming curSwara always in raga
	curNote := r.swaraLookup[curSwara]
	if curNote.avroh == nil {
		return 0, errors.New("lower than current octave / sthayi")
	}

	return curNote.avroh.frequency, nil
}

func TransformedSwaras(scale float32, swara float32) float32 {
	multiplier := scale / CScale
	return swara * multiplier
}
