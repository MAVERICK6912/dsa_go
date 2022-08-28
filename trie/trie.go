package trie

type Trie struct {
	root *node
}

type node struct {
	links [26]*node
	end   bool
}

// New instantiates the trie on which it is called upon.
func (t *Trie) New() *Trie {
	t.root = &node{
		end: false,
	}
	return t
}

// Insert inserts given string/strings into the trie on which it is called upon.
// time complexity for Insert is O(n*k), where n is the number of words and k is length of a word.
func (t *Trie) Insert(vals ...string) {
	for _, v := range vals {
		tempNode := t.root
		tempNode.insertWord(v)
	}
}

// insertWord inserts characters in a given word to the trie.
// time complexity for insertWord is O(k), where k is length of word.
func (n *node) insertWord(word string) {
	// insert each character in word into the trie
	for _, ch := range word {
		if !n.containsKey(ch) {
			n.makeLink(ch, node{end: false})
		}
		// move to next reference trie
		n = n.getNextLink(ch)
	}
	n.setEnd()
}

// Search returns true if word is found in trie otherwise returns false.
// time complextiy: O(n), where n is length of word to search.
func (t Trie) Search(word string) bool {
	tempNode := t.root
	for _, ch := range word {
		if !tempNode.containsKey(ch) {
			return false
		}
		tempNode = tempNode.getNextLink(ch)
	}
	return tempNode.isEnd()
}

// Search returns true if prefix is found in trie otherwise returns false.
// time complextiy: O(n), where n is length of prefix to search.
func (t Trie) StartsWith(prefix string) bool {
	tempNode := t.root
	for _, ch := range prefix {
		if !tempNode.containsKey(ch) {
			return false
		}
		tempNode = tempNode.getNextLink(ch)
	}
	return true
}

// getNextLink returns nextLink given the character.
func (n node) getNextLink(ch rune) *node {
	return n.links[ch-'a']
}

// setEnd sets end flag to true for the node it is called upon.
func (n *node) setEnd() {
	n.end = true
}

// isEnd returns true if current node is end of a word false otherwise.
func (n node) isEnd() bool {
	return n.end
}

// makeLink sets a new node at the position of passed character.
func (n *node) makeLink(ch rune, newNode node) {
	n.links[ch-'a'] = &newNode
}

// containsKey checks if the given character has already been inserted into the trie.
func (n node) containsKey(ch rune) bool {
	return n.links[ch-'a'] != nil
}
