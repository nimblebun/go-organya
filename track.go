package organya

// Track defines a note row of a given instrument.
type Track struct {
	Instrument        uint8  `json:"instrument"`
	ResourceCount     uint16 `json:"resourceCount"`
	LoopStartResource uint16 `json:"loopStartResource"`

	Resources []Resource `json:"resources"`
}
