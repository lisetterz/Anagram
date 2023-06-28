package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		word  string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "single word anagram",
			args: args{
				"cars",
				"scar",
			},
			want: true,
		},
		{
			name: "phrase with space and uppercase anagram",
			args: args{
				"I am Lord Voldemort",
				"Tom Malvoro Riddle",
			},
			want: true,
		},
		{
			name: "Negative - not anagram",
			args: args{
				"foo",
				"something",
			},
			want: false,
		},
		{
			name: "Negative - more repetitions on letters",
			args: args{
				"bear",
				"beaar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := formatString(tt.args.word)
			word2 := formatString(tt.args.word2)
			if got := isAnagram(word, word2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
