package template

import (
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestTemplate_Render(t *testing.T) {
	t.Run("Test render with variable passed", func(t *testing.T) {
		vars := map[string]ast.Variable{
			"foo": ast.Variable{
				Type: ast.TypeMap,
				Value: map[string]ast.Variable{
					"bar": ast.Variable{
						Type:  ast.TypeString,
						Value: "Bar",
					},
				},
			},
		}
		tmpl, err := NewTemplate(vars)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		result, err := tmpl.Render(`${has(foo, "bar")}`)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if result.Value.(string) != "true" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "true", result.Value.(string))
		}
	})

	t.Run("Test render with nil variables", func(t *testing.T) {
		tmpl, err := NewTemplate(nil)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		result, err := tmpl.Render(`${lower("FOO")}`)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if result.Value.(string) != "foo" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "foo", result.Value.(string))
		}
	})
}