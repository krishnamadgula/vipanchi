package types

import (
	"bytes"
	"text/template"
)

type SwaraLipi struct {
	T          *Taalam
	R          *Raagam
	Kaala      int
	lineLength int
	Lines      []*Line
}

const docTmpl = `
Raagam: {{ .R.Name }}
Taalam: {{ .T.Name }}
Kaala: {{ .Kaala }}

{{ range $i, $line := .Lines }}
Line: 
{{- range $j, $akshara := $line.Aksharas }}
    {{ range $k, $matra := $akshara.Matras }}{{ if $k }}, {{ end }}{{ $matra }}{{ end }}
{{- end }}
{{ end }}
`

func NewSwaraLipi(r *Raagam, t *Taalam, kaala int) *SwaraLipi {
	length := 0
	for _, l := range t.Sequence {
		length += l
	}

	return &SwaraLipi{
		R:          r,
		T:          t,
		lineLength: length,
		Kaala:      kaala,
	}
}

func (s *SwaraLipi) AppendMatra(matra string) {
	if len(s.Lines) == 0 {
		s.Lines = append(s.Lines, &Line{})
	}

	lastLine := s.Lines[len(s.Lines)-1]
	if len(lastLine.Aksharas) == s.lineLength {
		s.Lines = append(s.Lines, &Line{})
	}

	lastLine = s.Lines[len(s.Lines)-1]
	if len(lastLine.Aksharas) == 0 {
		lastLine.appendAkshara(&Akshara{})
	}

	lastAkshara := lastLine.Aksharas[len(lastLine.Aksharas)-1]

	if len(lastAkshara.Matras) == s.Kaala {
		lastLine.appendAkshara(&Akshara{})
	}

	lastAkshara = lastLine.Aksharas[len(lastLine.Aksharas)-1]
	lastAkshara.appendMatra(matra)

}

func (s *SwaraLipi) Notes() []float32 {
	var swaras []string

	// Loop through each line in the SwaraLipi
	for _, line := range s.Lines {
		// Loop through each Akshara in the line
		for _, akshara := range line.Aksharas {
			// Append all Matras (swaras) from the Akshara to the result
			swaras = append(swaras, akshara.Matras...)
		}
	}

	frequencies := make([]float32, 0, len(swaras))
	for _, swara := range swaras {
		frequencies = append(frequencies, Swaras[swara])
	}

	return frequencies
}

func (s *SwaraLipi) Doc() string {
	// handle err
	tmpl, _ := template.New("musicSheet").Parse(docTmpl)
	out := bytes.NewBufferString("")
	_ = tmpl.Execute(out, *s)
	return out.String()
}

// func (s *SwaraLipi) appendLine() {
// 	s.Lines = append(
// 		s.Lines,
// 		&Line{},
// 	)
// }

// func (s *SwaraLipi) newAkshara() {
// 	lastLine := s.Lines[len(s.Lines)-1]
// 	if len(lastLine.Aksharas) == s.lineLength {
// 		s.newLine()
// 	}

// 	lastLine.Aksharas = append(
// 		lastLine.Aksharas,
// 		Akshara{Matras: make([]string, 0, s.kaala)},
// 	)
// }

type Line struct {
	// Aksharas constitute a line
	Aksharas []*Akshara
}

func (l *Line) appendAkshara(a *Akshara) {
	l.Aksharas = append(l.Aksharas, a)
}

type Akshara struct {
	// matras are the individual swara that are ung in each Akshara.
	Matras []string
}

func (a *Akshara) appendMatra(matra string) {
	a.Matras = append(a.Matras, matra)
}
