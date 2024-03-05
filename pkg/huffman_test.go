// Description: Huffman encoder implementation
package pkg

import (
	"reflect"
	"testing"
)

func Test_genNodes(t *testing.T) {
	type args struct {
		stats map[string]int
	}
	tests := []struct {
		name string
		args args
		want []Node
	}{
		{
			name: "Test node generation in sorted order",
			args: args{
				stats: map[string]int{
					"c": 5,
					"d": 2,
				},
			},
			want: []Node{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenNodes(tt.args.stats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
