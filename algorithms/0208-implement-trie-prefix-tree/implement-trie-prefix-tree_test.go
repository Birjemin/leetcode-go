package implement_trie_prefix_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	trie := Constructor()

	trie.Insert("apple")
	ast.Equal(true, trie.Search("apple"))   // returns true
	ast.Equal(false, trie.Search("app"))    // returns false
	ast.Equal(true, trie.StartsWith("app")) // returns true
	trie.Insert("app")
	ast.Equal(true, trie.Search("app")) // returns true
}
