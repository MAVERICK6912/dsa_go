package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trieInsertTestData = []struct {
	input    []string
	expected Trie
}{{
	input: []string{"abc"},
	expected: Trie{
		root: &node{
			links: [26]*node{
				{
					links: [26]*node{
						nil,
						{
							links: [26]*node{
								nil,
								nil,
								{
									end: true,
								},
							},
							end: false,
						},
					},
					end: false,
				},
			},
		},
	},
}}

var trieSearchTestData = []struct {
	input    []string
	target   []string
	expected []bool
}{
	{
		input:    []string{"hawkgirl", "flash"},
		target:   []string{"hawkgirl", "robin"},
		expected: []bool{true, false},
	},
	{
		input:    []string{"batman", "robin", "catwoman", "nightwing"},
		target:   []string{"robin", "batgirl", "nightwing"},
		expected: []bool{true, false, true},
	},
}

var trieStartsWithTestData = []struct {
	input    []string
	target   []string
	expected []bool
}{
	{
		input:    []string{"hawkgirl", "flash"},
		target:   []string{"haw", "d"},
		expected: []bool{true, false},
	},
	{
		input:    []string{"batman", "robin", "catwoman", "nightwing"},
		target:   []string{"ro", "batg", "night"},
		expected: []bool{true, false, true},
	},
}

func TestInsert(t *testing.T) {
	for _, val := range trieInsertTestData {
		var trie Trie
		trie.New()
		trie.Insert(val.input...)
		assert.Equal(t, trie, val.expected)
	}
}

func TestSearch(t *testing.T) {
	for _, data := range trieSearchTestData {
		var trie Trie
		trie.New()
		trie.Insert(data.input...)
		for ind, target := range data.target {
			actual := trie.Search(target)
			assert.Equal(t, data.expected[ind], actual)
		}
	}
}

func TestStartsWith(t *testing.T) {
	for _, data := range trieSearchTestData {
		var trie Trie
		trie.New()
		trie.Insert(data.input...)
		for ind, target := range data.target {
			actual := trie.StartsWith(target)
			assert.Equal(t, data.expected[ind], actual)
		}
	}
}
