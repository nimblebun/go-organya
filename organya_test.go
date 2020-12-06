package organya_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"pkg.nimblebun.works/go-organya"
)

func TestOrganya(t *testing.T) {
	t.Run("should return correct Organya info", func(t *testing.T) {
		correctData, err := ioutil.ReadFile("./__tests__/Balcony.txt")
		if err != nil {
			panic(err)
		}

		// This file was generated using the original COrg library using the same
		// algorithm. We're checking our own output against this file to see if
		// everything's correct.
		org, err := organya.Open("./__tests__/Balcony.org")
		if err != nil {
			t.Errorf("Opening organya file failed with error %v", err)
		}

		output := ""

		output += fmt.Sprintf("Wait value: %d\n", org.WaitValue)
		output += fmt.Sprintf("Loop start: %d\n", org.LoopStart)
		output += fmt.Sprintf("Loop end: %d\n", org.LoopEnd)

		for i, track := range org.Tracks {
			output += fmt.Sprintf("\nTrack %d:\n\n", i)

			output += fmt.Sprintf("Instrument: %d\n", track.Instrument)
			output += fmt.Sprintf("Resource count: %d\n", track.ResourceCount)
			output += fmt.Sprintf("Loop start resource: %d\n", track.LoopStartResource)

			for j, res := range track.Resources {
				output += fmt.Sprintf("\nResource %d:\n", j)

				output += fmt.Sprintf("    Start: %d\n", res.Start)
				output += fmt.Sprintf("    Duration: %d\n", res.Duration)
				output += fmt.Sprintf("    Note: %d\n", res.Note)
				output += fmt.Sprintf("    Volume: %d\n", res.Volume)
				output += fmt.Sprintf("    Pan: %d\n", res.Pan)
			}

			output += fmt.Sprintf("\n============================\n")
		}

		if output != string(correctData) {
			t.Errorf("Incorrect data found for the provied Organya file.")
		}
	})
}
