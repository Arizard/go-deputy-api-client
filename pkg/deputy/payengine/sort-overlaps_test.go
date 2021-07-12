package payengine

import (
	"reflect"
	"testing"
)

func TestSortOverlaps(t *testing.T) {
	type args struct {
		overlaps []Overlap
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []Overlap
	}{
		{
			"single sort",
			args{
				[]Overlap{
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 20,
						},
					},
				},
			},
			[]Overlap{
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 20,
					},
				},
			},
		},
		{
			"bigger sort",
			args{
				[]Overlap{
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 20,
						},
					},
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 10,
						},
					},
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 1000,
						},
					},
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 400,
						},
					},
					{
						Bounds: TimeRange{},
						PayRule:   &PayRule{
							Rank: 5,
						},
					},
				},
			},
			[]Overlap{
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 1000,
					},
				},
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 400,
					},
				},
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 20,
					},
				},
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 10,
					},
				},
				{
					Bounds: TimeRange{},
					PayRule:   &PayRule{
						Rank: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SortOverlaps(tt.args.overlaps); !reflect.DeepEqual(tt.args.overlaps, tt.wantSorted) {
				gotSorted := tt.args.overlaps
				t.Errorf("SortOverlaps() = %v, want %v", gotSorted, tt.wantSorted)
				for _, o := range gotSorted {
					t.Logf("got %d", o.PayRule.Rank)
				}
				for _, o := range tt.wantSorted {
					t.Logf("want %d", o.PayRule.Rank)
				}
			}
		})
	}
}