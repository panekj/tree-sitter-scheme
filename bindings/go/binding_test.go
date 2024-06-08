package tree_sitter_scheme_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-scheme"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_scheme.Language())
	if language == nil {
		t.Errorf("Error loading Scheme grammar")
	}
}
