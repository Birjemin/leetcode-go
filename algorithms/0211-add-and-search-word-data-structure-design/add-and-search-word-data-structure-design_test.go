package add_and_search_word_data_structure_design

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	obj.AddWord("ba")

	ast.Equal(false, obj.Search("pad"))
	ast.Equal(true, obj.Search("bad"))
	ast.Equal(true, obj.Search(".ad"))
	ast.Equal(true, obj.Search("b.."))
	ast.Equal(true, obj.Search("b."))
	ast.Equal(true, obj.Search("ba"))
	ast.Equal(false, obj.Search("b"))
}
