package organya

// Resource defines a single sound resource.
type Resource struct {
	Start    uint32 `json:"start"`
	Duration uint8  `json:"duration"`
	Note     uint8  `json:"note"`
	Volume   uint8  `json:"volume"`
	Pan      uint8  `json:"pan"`
}
