package lesson_7

import (
	"reflect"
	"testing"
)

func TestSortStr(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test 1",
			args: args{s: []string{"a", "c", "b"}},
			want: []string{"a", "b", "c"},
		},
		{name: "Test 2",
			args: args{s: []string{"0", "c", "11"}},
			want: []string{"0", "11", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
