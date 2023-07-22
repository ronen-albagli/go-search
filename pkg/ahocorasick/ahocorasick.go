package ahocorasick

import (
	timer "go-search/pkg/timer"
)

// Define a trie node structure.
type trieNode struct {
	children    map[rune]*trieNode
	isEndOfWord bool
	fail        *trieNode
	output      []string
}

// Initialize a new trie node.
func newTrieNode() *trieNode {
	return &trieNode{
		children:    make(map[rune]*trieNode),
		isEndOfWord: false,
		fail:        nil,
		output:      nil,
	}
}

// Add a pattern to the trie.
func insert(root *trieNode, pattern string) {
	node := root
	for _, char := range pattern {
		if node.children[char] == nil {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
	node.output = append(node.output, pattern)
}

// Build failure links using Breadth-First Search (BFS).
func buildFailureLinks(root *trieNode) {
	queue := make([]*trieNode, 0)

	// Initialize the root's children to point to the root itself.
	for _, child := range root.children {
		child.fail = root
		queue = append(queue, child)
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for char, child := range currentNode.children {
			queue = append(queue, child)
			failureNode := currentNode.fail
			for failureNode != nil && failureNode.children[char] == nil {
				failureNode = failureNode.fail
			}
			if failureNode == nil {
				child.fail = root
			} else {
				child.fail = failureNode.children[char]
				child.output = append(child.output, child.fail.output...)
			}
		}
	}
}

// Search for patterns in the text using the Aho-Corasick algorithm.
func ahoCorasickSearch(root *trieNode, text string) []string {
	results := make([]string, 0)
	node := root

	for _, char := range text {
		for node != nil && node.children[char] == nil {
			node = node.fail
		}
		if node == nil {
			node = root
			continue
		}
		node = node.children[char]
		if node.isEndOfWord {
			results = append(results, node.output...)
		}
	}

	return results
}

func Search(text string, patterns []string) []string {
	t := &timer.Timer{}

	t.Start()
	// Initialize the root of the trie.
	root := newTrieNode()

	// Define the patterns to search for.
	// Insert the patterns into the trie.
	for _, pattern := range patterns {
		insert(root, pattern)
	}

	// Build the failure links in the trie.
	buildFailureLinks(root)

	results := ahoCorasickSearch(root, text)

	t.Finish()
	return results

}
