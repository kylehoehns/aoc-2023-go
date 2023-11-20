package strings

import "testing"

func TestHasAllUniqueRunes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "All the same letter",
			args: args{"AAAAA"},
			want: false,
		},
		{
			name: "Upper and lowercase should be unique",
			args: args{"Aa"},
			want: true,
		},
		{
			name: "Allows non-letter chars",
			args: args{"-)(!*&"},
			want: true,
		},
		{
			name: "Normal non-unique input",
			args: args{"Hello"},
			want: false,
		},
		{
			name: "Normal unique input",
			args: args{"Alphabet"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAllUniqueRunes(tt.args.input); got != tt.want {
				t.Errorf("HasAllUniqueRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
