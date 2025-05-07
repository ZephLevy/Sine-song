package songgenerator

import (
	"fmt"
	"log"
	"math"
)

type note struct {
	duration float64
	name     string
}

func GetSong(sampleRateInt int) []int16 {
	sampleRate := float64(sampleRateInt)
	var frequencies []float64
	notes := []note{
		{duration: 1, name: "D5"},
		{duration: 1, name: "B4"},
		{duration: 1, name: "G4"},
		{duration: 1, name: "D4"},
		{duration: 1 / 3.0, name: "E4"},
		{duration: 1 / 3.0, name: "F#4"},
		{duration: 1 / 3.0, name: "G4"},
		{duration: 2 / 3.0, name: "E4"},
		{duration: 1 / 3.0, name: "G4"},
		{duration: 2, name: "D4"},
	}
	fmt.Println(notes)
	for _, note := range notes {
		frequency, err := getFrequency(note.name)
		if err != nil {
			// Assume a typo or smth equivalent
			// Still, I'm printing to let the user know
			log.Println("Error getting frequency:", err)
			continue
		}
		for i := range int(note.duration * sampleRate) {
			_ = i
			frequencies = append(frequencies, frequency)
		}
	}

	var samples []int16
	for i, frequency := range frequencies {
		sample := int16(math.Sin(2.0*math.Pi*frequency*float64(i)/sampleRate) * 32767)
		samples = append(samples, sample)
	}
	return samples
}
