package payengine

import (
	"reflect"
	"testing"
	"time"
)

func TestWindow(t *testing.T) {
	type args struct {
		input  TimeRange
		window TimeRange
	}
	tests := []struct {
		name string
		args args
		wantRange TimeRange
	}{
		{
			"window inside",
			args{
				TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
				TimeRange{Start: time.Date(2021, 1, 1, 10, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			},
			TimeRange{Start: time.Date(2021, 1, 1, 10, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
		{
			"window outside",
			args{
				TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
				TimeRange{Start: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 18, 0, 0, 0, time.UTC)},
			},
			TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
		},
		{
			"window start",
			args{
				TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
				TimeRange{Start: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC)},
		},
		{
			"window end",
			args{
				TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
				TimeRange{Start: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 20, 0, 0, 0, time.UTC)},
			},
			TimeRange{Start: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
		},
		{
			"window none",
			args{
				TimeRange{Start: time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
				TimeRange{Start: time.Date(2021, 1, 1, 18, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 20, 0, 0, 0, time.UTC)},
			},
			TimeRange{Start: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC), End: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if window(&tt.args.input, tt.args.window); !reflect.DeepEqual(tt.args.input, tt.wantRange) {
				t.Errorf("got %+v, want %+v", tt.args.input, tt.wantRange)
			}
		})
	}
}