package Testing

// We can create an array of structures.
//This will contain the description, the input and the expected output.
var testCases = []struct {
	description string
	input []int
	expected int
} {
	{
		description: "Contains all negative",
		input: []int{-1,-2,-3,-4},
		expected: 0,
	},
	{
		description: "Contains all positive",
		input: []int{1,2,3,4,5},
		expected: 15,
	},
	{
		description: "Contains all zero",
		input: []int{0,0,0,0},
		expected: 0,
	},
	{
		description: "Contains some positive and some negative",
		input: []int{1,2,3,-4},
		expected: 6,
	},
}
