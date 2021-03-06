package interpolation

import "testing"

func TestInterpolationFuncMax(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Should find the max value",
			text:        `${max(0, 2)}`,
			expectation: "2",
		},
		{
			description: "One input should fail",
			text:        `${max(4)}`,
			expectation: "",
			evalError:   true,
		},
		{
			description: "Long list finds max",
			text:        `${max(100, 30, 450, 1, 25)}`,
			expectation: "450",
		},
	}

	maxTestFunc := testInterpolationFunc(keyFuncs{"max": interpolationFuncMax})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			maxTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncMin(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Should find the min value",
			text:        `${min(140, 250)}`,
			expectation: "140",
		},
		{
			description: "One input should fail",
			text:        `${min(4)}`,
			expectation: "",
			evalError:   true,
		},
		{
			description: "Long list finds min",
			text:        `${min(100, 30, 450, 1, 25)}`,
			expectation: "1",
		},
	}

	minTestFunc := testInterpolationFunc(keyFuncs{"min": interpolationFuncMin})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			minTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncAvg(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Find average of 20 and 80",
			text:        `${avg(20, 80)}`,
			expectation: "50",
		},
		{
			description: "Single input finds average",
			text:        `${avg(10)}`,
			expectation: "10",
		},
		{
			description: "No inputs returns error",
			text:        `${avg()}`,
			evalError:   true,
		},
		{
			description: "0 value returns average",
			text:        `${avg(0)}`,
			expectation: "0",
		},
	}

	avgTestFunc := testInterpolationFunc(keyFuncs{"avg": interpolationFuncAvg})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			avgTestFunc(t, tc)
		})
	}
}
