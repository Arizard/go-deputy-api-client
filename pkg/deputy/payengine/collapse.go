package payengine

import "time"

type TimeRange struct {
	Start time.Time
	End time.Time
}

func window(input *TimeRange, windowBounds TimeRange) {
	if windowBounds.Start.After(input.End) {
		input.Start = input.End
		return
	}

	if windowBounds.End.Before(input.Start) {
		input.End = input.Start
		return
	}

	if input.Start.Before(windowBounds.Start) {
		input.Start = windowBounds.Start
	}

	if input.End.After(windowBounds.End) {
		input.End = windowBounds.End
	}
}

func windowOverlaps(overlaps []Overlap, windowBounds TimeRange) {
	for _, overlap := range overlaps {
		window(&overlap.Bounds, windowBounds)
	}
}

func Collapse(overlaps []Overlap, window TimeRange) (result []Overlap) {
	SortOverlaps(overlaps)

	windowOverlaps(overlaps, window)

	return result
}
