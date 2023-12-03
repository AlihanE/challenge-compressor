package huffman

import "testing"

func TestSorter_Parse(t *testing.T) {
	type args struct {
		input string
	}

	sorter := NewSorter()
	tests := []struct {
		name string
		s    *Sorter
		args args
	}{
		{
			name: "1",
			s:    sorter,
			args: args{
				input: "99,2;101,2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Parse(tt.args.input)
			if tt.s.nodes[0].Weight() != 2 {
				t.Error("Invalid weight. Weight:", tt.s.nodes[0].Weight())
			}
		})
	}
}
