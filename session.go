package organya

// Session is the playback session of an Organya object.
type Session struct {
	Angles       [OrgTrackCount]float64
	ResourceUpTo [OrgTrackCount]uint16
	CurrentClick uint32

	Org *Organya
}

func shouldUpdate(s *Session, i int) bool {
	if s.ResourceUpTo[i] >= s.Org.Tracks[i].ResourceCount-1 {
		return false
	}

	if s.CurrentClick < s.Org.Tracks[i].Resources[s.ResourceUpTo[i]+1].Start {
		return false
	}

	return true
}

// Click will advance in the track.
func (s *Session) Click() {
	s.CurrentClick++

	if s.CurrentClick >= s.Org.LoopEnd {
		s.CurrentClick = s.Org.LoopStart

		for i := 0; i < OrgTrackCount; i++ {
			s.ResourceUpTo[i] = s.Org.Tracks[i].LoopStartResource
		}
	}

	for i := 0; i < OrgTrackCount; i++ {
		if shouldUpdate(s, i) {
			s.ResourceUpTo[i]++
		}
	}
}

// GetResource will return the current resource from the provided track.
func (s *Session) GetResource(track int) *Resource {
	return &(s.Org.Tracks[track].Resources[s.ResourceUpTo[track]])
}

// TrackSounding specifies whether the provided track is active.
func (s *Session) TrackSounding(track int) bool {
	currentResource := s.GetResource(track)

	start := currentResource.Start
	end := currentResource.Start + uint32(currentResource.Duration) - 1

	if s.ResourceUpTo[track] >= s.Org.Tracks[track].ResourceCount {
		return false
	}

	if s.CurrentClick < start || s.CurrentClick > end {
		return false
	}

	return true
}
