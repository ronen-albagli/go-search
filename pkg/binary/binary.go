package binary

import "go-search/pkg/timer"

// Define a trie node structure.
type trieNode struct {
	children    map[byte]*trieNode
	isEndOfWord bool
}

// Initialize a new trie node.
func newTrieNode() *trieNode {
	return &trieNode{
		children:    make(map[byte]*trieNode),
		isEndOfWord: false,
	}
}

// Add a word to the trie.
func insert(root *trieNode, word string) {
	node := root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

// Search for words in the text using the binary search trie.
func searchWords(root *trieNode, text string) []string {
	foundWords := make([]string, 0)

	for i := 0; i < len(text); i++ {
		node := root
		for j := i; j < len(text); j++ {
			char := text[j]
			if node.children[char] == nil {
				break
			}
			node = node.children[char]
			if node.isEndOfWord {
				foundWords = append(foundWords, text[i:j+1])
			}
		}
	}

	return foundWords
}

// BinarySearchSubstring searches for multiple words in a given text using a binary search trie and returns a slice of found words.
func Search(text string, words []string) []string {
	// Initialize the root of the trie.
	t := &timer.Timer{}
	t.Start()
	root := newTrieNode()

	// Define the words to search for.
	// Insert the words into the trie.
	for _, word := range words {
		insert(root, word)
	}

	foundWords := searchWords(root, text)

	t.Finish()

	return foundWords
}
