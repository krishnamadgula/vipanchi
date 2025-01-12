package instruments

var InstrumentLookup = map[string]*Instrument{
	"veena":           &Veena,
	"violin":          &Violin,
	"flute":           &Flute,
	"harmonica":       &Harmonica,
	"bagpiper":        &Bagpiper,
	"electric_guitar": &ElectricGuitar,
	"shenai":          &Shenai,
}
