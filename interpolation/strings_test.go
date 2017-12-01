package interpolation

import (
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestInterpolationFuncLower(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Uppercase becomes lowercase",
			text:        `${lower("FOO")}`,
			expectation: "foo",
		},
		{
			description: "Lowercase stays lowercase",
			text:        `${lower("foo")}`,
			expectation: "foo",
		},
	}

	lowerTestFunc := testInterpolationFunc(keyFuncs{"lower": interpolationFuncLower})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lowerTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncUpper(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Uppercase stays uppercase",
			text:        `${upper("FOO")}`,
			expectation: "FOO",
		},
		{
			description: "Lowercase becomes uppercase",
			text:        `${upper("foo")}`,
			expectation: "FOO",
		},
	}

	lowerTestFunc := testInterpolationFunc(keyFuncs{"upper": interpolationFuncUpper})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lowerTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncJoin(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Joins multiple values",
			text:        `${join(",", i)}`,
			expectation: "Foo,Bar",
			vars: map[string]ast.Variable{
				"i": ast.Variable{
					Type: ast.TypeList,
					Value: []ast.Variable{
						{
							Type:  ast.TypeString,
							Value: "Foo",
						},
						{
							Type:  ast.TypeString,
							Value: "Bar",
						},
					},
				},
			},
		},
		{
			description: "Bad variable length fails",
			text:        `${join(",")}`,
			expectation: "",
			evalError:   true,
		},
		{
			description: "Bad parse",
			text:        `${join(",", [4]}`,
			expectation: "",
			parseError:  true,
		},
		{
			description: "Bad array item",
			text:        `${join(",", i)}`,
			expectation: "Foo,Bar",
			vars: map[string]ast.Variable{
				"i": ast.Variable{
					Type: ast.TypeList,
					Value: []ast.Variable{
						{
							Type:  ast.TypeInt,
							Value: 4,
						},
					},
				},
			},
			evalError: true,
		},
	}

	joinTestFunc := testInterpolationFunc(keyFuncs{"join": interpolationFuncJoin})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			joinTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncReplace(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Replace N occurences",
			text:        `${replace("foo bar bar", " bar", "", -1)}`,
			expectation: "foo",
		},
		{
			description: "Replace 0 occurences",
			text:        `${replace("foo bar bar", " bar", "", 0)}`,
			expectation: "foo bar bar",
		},
		{
			description: "Replace 1 occurences",
			text:        `${replace("foo bar bar", " bar", "", 1)}`,
			expectation: "foo bar",
		},
	}

	replaceTestFunc := testInterpolationFunc(keyFuncs{"replace": interpolationFuncReplace})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			replaceTestFunc(t, tc)
		})
	}
}
