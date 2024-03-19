package pkg

import (
	"reflect"
	"testing"
)

func TestSortedNodes_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *SortedNodes
		want Node
	}{
		{
			name: "Test node generation in sorted order",
			s: &SortedNodes{
				{
					Freq:  2,
					Val:   "d",
					Left:  nil,
					Right: nil,
				},
				{
					Freq:  5,
					Val:   "c",
					Left:  nil,
					Right: nil,
				},
			},
			want: Node{
				Freq:  2,
				Val:   "d",
				Left:  nil,
				Right: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedNodes.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
