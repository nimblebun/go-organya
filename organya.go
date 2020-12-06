package organya

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"os"
)

// Organya defines the file data of an Organya file.
type Organya struct {
	WaitValue uint16 `json:"waitValue"`
	LoopStart uint32 `json:"loopStart"`
	LoopEnd   uint32 `json:"loopEnd"`

	Tracks [OrgTrackCount]Track `json:"tracks"`
}

// Open will open the Organya file on the provided file path and return its
// data.
func Open(filename string) (*Organya, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var organya Organya

	f.Seek(6, io.SeekCurrent)

	waitValue := make([]byte, 2)
	f.Read(waitValue)
	organya.WaitValue = binary.LittleEndian.Uint16(waitValue)

	f.Seek(2, io.SeekCurrent)

	loopStart := make([]byte, 4)
	loopEnd := make([]byte, 4)
	f.Read(loopStart)
	f.Read(loopEnd)
	organya.LoopStart = binary.LittleEndian.Uint32(loopStart)
	organya.LoopEnd = binary.LittleEndian.Uint32(loopEnd)

	instrument := make([]byte, 1)
	resourceCount := make([]byte, 2)

	for i := 0; i < OrgTrackCount; i++ {
		f.Seek(2, io.SeekCurrent)
		f.Read(instrument)
		organya.Tracks[i].Instrument = instrument[0]

		f.Seek(1, io.SeekCurrent)
		f.Read(resourceCount)
		organya.Tracks[i].ResourceCount = binary.LittleEndian.Uint16(resourceCount)
	}

	for i := 0; i < OrgTrackCount; i++ {
		assigned := false

		organya.Tracks[i].Resources = make([]Resource, organya.Tracks[i].ResourceCount)

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			start := make([]byte, 4)
			f.Read(start)
			organya.Tracks[i].Resources[j].Start = binary.LittleEndian.Uint32(start)

			if !assigned && organya.Tracks[i].Resources[j].Start >= organya.LoopStart {
				organya.Tracks[i].LoopStartResource = j
				assigned = true
			}
		}

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			note := make([]byte, 1)
			f.Read(note)
			organya.Tracks[i].Resources[j].Note = note[0]
		}

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			if organya.Tracks[i].Resources[j].Note == OrgNoChange {
				f.Seek(1, io.SeekCurrent)
				organya.Tracks[i].Resources[j].Duration = organya.Tracks[i].Resources[j-1].Duration
			} else {
				duration := make([]byte, 1)
				f.Read(duration)
				organya.Tracks[i].Resources[j].Duration = duration[0]
			}
		}

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			if organya.Tracks[i].Resources[j].Volume == OrgNoChange {
				organya.Tracks[i].Resources[j].Volume = organya.Tracks[i].Resources[j-1].Volume
			} else {
				volume := make([]byte, 1)
				f.Read(volume)
				organya.Tracks[i].Resources[j].Volume = volume[0]
			}
		}

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			pan := make([]byte, 1)
			f.Read(pan)

			if pan[0] == OrgNoChange {
				organya.Tracks[i].Resources[j].Pan = organya.Tracks[i].Resources[j-1].Pan
			} else {
				organya.Tracks[i].Resources[j].Pan = pan[0]
			}
		}

		for j := uint16(0); j < organya.Tracks[i].ResourceCount; j++ {
			if organya.Tracks[i].Resources[j].Note == OrgNoChange {
				organya.Tracks[i].Resources[j].Note = organya.Tracks[i].Resources[j-1].Note
			}
		}
	}

	return &organya, nil
}

// NewSession will initialize a playback session on the provided Organya object.
func (org *Organya) NewSession() *Session {
	var session *Session

	session.Org = org
	session.CurrentClick = 0

	for i := 0; i < OrgTrackCount; i++ {
		session.Angles[i] = 0
		session.ResourceUpTo[i] = 0
	}

	return session
}

// JSON will convert the current Organya object into JSON.
func (org *Organya) JSON() ([]byte, error) {
	return json.Marshal(org)
}
