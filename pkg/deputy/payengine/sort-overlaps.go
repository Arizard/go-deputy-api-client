package payengine

import "sort"

func SortOverlaps(overlaps []Overlap) {
	rank := func(i, j int) bool {
		o1 := overlaps[i]
		o2 := overlaps[j]
		return o1.PayRule.Rank > o2.PayRule.Rank
	}

	sort.Slice(overlaps, rank)
}
