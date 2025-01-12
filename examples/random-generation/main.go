package main

import (
	"log/slog"
	"os"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/synth-veena/pkg/generator"
	"github.com/synth-veena/pkg/instruments"
	"github.com/synth-veena/pkg/raagas"
	"github.com/synth-veena/pkg/taalas"
	"github.com/synth-veena/pkg/types"
)

var (
	raagamName     string
	taalamName     string
	kaalam         int
	tempo          int
	cycles         int
	instrumentName string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&raagamName, "raagam", "r", "shivaranjani", "Name of the raagam")
	rootCmd.PersistentFlags().StringVarP(&taalamName, "taalam", "t", "rupaka", "Name of the taalam")
	rootCmd.PersistentFlags().IntVarP(&kaalam, "kaalam", "k", types.Kaala2nd, "Kaalam (speed) of the composition")
	rootCmd.PersistentFlags().IntVarP(&tempo, "tempo", "b", 90, "Tempo in beats per minute")
	rootCmd.PersistentFlags().IntVarP(&cycles, "cycles", "c", 10, "Number of taalam cycles")
	rootCmd.PersistentFlags().StringVarP(&instrumentName, "instrument", "i", "veena", "Name of the instrument")

	viper.BindPFlag("raagam", rootCmd.PersistentFlags().Lookup("raagam"))
	viper.BindPFlag("taalam", rootCmd.PersistentFlags().Lookup("taalam"))
	viper.BindPFlag("kaalam", rootCmd.PersistentFlags().Lookup("kaalam"))
	viper.BindPFlag("tempo", rootCmd.PersistentFlags().Lookup("tempo"))
	viper.BindPFlag("cycles", rootCmd.PersistentFlags().Lookup("cycles"))
	viper.BindPFlag("instrument", rootCmd.PersistentFlags().Lookup("instrument"))
}

func initConfig() {
	viper.SetEnvPrefix("SYNTH")
	viper.AutomaticEnv()
}

var helpText = `
Synth-Veena is a Carnatic music generator that creates compositions in various raagams and taalams.

Available Raagams:
  - shivaranjani     - Pentatonic raagam, peaceful mood
  - kalyani          - Major scale equivalent, bright and devotional
  - sankarabharanam  - Similar to Western major scale
  - chakravaaka      - Equivalent to Western major scale
  - mohana           - Pentatonic scale, peaceful and devotional
  - hindola          - Pentatonic scale with unique intervals
  - brindavana_saarangi - Romantic and devotional mood
  - bhairavi         - Morning raagam, devotional
  - hamsadhvani      - Light and quick compositions
  - thodi            - Complex morning raagam
  - sindhu_bhairavi  - Allows all 12 semitones
  - mayamalavagowla  - Parent scale of Carnatic music
  - aananda_bhairavi - Devotional morning raagam

Available Taalams:
  - aadi     - 8 beats (4+2+2)
  - rupaka   - 6 beats (2+4) 
  - triputa  - 7 beats (3+2+2)

Available Instruments:
  - veena           - Traditional Indian string instrument
  - violin          - Western string instrument
  - flute           - Traditional Indian bamboo flute
  - harmonica       - Western free reed instrument
  - bagpiper        - Traditional Scottish wind instrument
  - electric_guitar - Modern electric string instrument

The generated output will be two files:
  - A .wav audio file with the generated music
  - A .txt file containing the swara notation

Example:
  synth-veena -r bhairavi -t aadi -k 2 -b 90 -c 10 -i veena
`

func init() {
	rootCmd.Long = helpText
	rootCmd.MarkPersistentFlagRequired("raagam")
	rootCmd.MarkPersistentFlagRequired("taalam")
	rootCmd.MarkPersistentFlagRequired("kaalam")
	rootCmd.MarkPersistentFlagRequired("tempo")
	rootCmd.MarkPersistentFlagRequired("cycles")
}

var rootCmd = &cobra.Command{
	Use:   "synth-veena",
	Short: "Generate Carnatic music compositions",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			raagam     *types.Raagam
			ok         bool
			taalam     *types.Taalam
			instrument *instruments.Instrument
		)
		if raagam, ok = raagas.RaagaLookup[raagamName]; !ok {
			return fmt.Errorf("unsupported raagam: %s", raagamName)
		}

		if taalam, ok = taalas.TaalaLookup[taalamName]; !ok {
			return fmt.Errorf("unsupported taalam: %s", taalamName)
		}

		if instrument, ok = instruments.InstrumentLookup[instrumentName]; !ok {
			return fmt.Errorf("unsupported instrument: %s", instrumentName)
		}

		return generator.Generate(
			generator.RandomGenerate,
			*raagam,
			*taalam,
			kaalam,
			tempo,
			cycles,
			*instrument,
		)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("unable to generate: " + err.Error())
		os.Exit(1)
	}
}
