package hw

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name         string
		args         args
		wantDistance float64
	}{
		{
			name:         "#1",
			args:         args{x1: 1, y1: 1, x2: 4, y2: 5},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := Distance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
