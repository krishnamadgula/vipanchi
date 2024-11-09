package instruments

type Instrument struct {
	// Harmonics represents the relative amplitudes of overtones
	// Each entry is (harmonic number, relative amplitude)
	Harmonics [][2]float64
	// Attack is the time in seconds for the note to reach full amplitude
	Attack float64
	// Decay is the time in seconds for the note to decay to sustain level
	Decay float64
	// Sustain is the amplitude level during the main note duration (0-1)
	Sustain float64
	// Release is the time in seconds for the note to fade to zero
	Release float64
}

func GetEnvelopeMultiplier(t, duration float64, inst Instrument) float64 {
	attackTime := inst.Attack
	decayTime := inst.Decay
	releaseTime := inst.Release
	sustainLevel := inst.Sustain

	if t < attackTime {
		return t / attackTime
	} else if t < attackTime+decayTime {
		decayProgress := (t - attackTime) / decayTime
		return 1.0 - (1.0-sustainLevel)*decayProgress
	} else if t < duration-releaseTime {
		return sustainLevel
	} else {
		releaseProgress := (t - (duration - releaseTime)) / releaseTime
		return sustainLevel * (1.0 - releaseProgress)
	}
}
