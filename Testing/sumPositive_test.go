package Testing

import "testing"

func TestScore(t *testing.T) {
	for _,tc := range testCases {
		t.Run(tc.description,func(t *testing.T){
			actual := sumPositive(tc.input)
			if actual != tc.expected {
				t.Errorf("Score(%d) = %#v, want: %#v",
					tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			sumPositive(tc.input)
		}
	}
}

