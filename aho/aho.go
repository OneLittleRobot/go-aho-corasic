package aho

import (
	"unicode/utf8"
)

func check(fall *TrieNode, trieNode *TrieNode, char rune) bool {
	checkNode := fall.getChild(char)
	if checkNode != nil && fall != trieNode {
		return false
	}
	return true
}

func build(trieNode *TrieNode) *TrieNode {
	var queue []*TrieNode
	trieNode.fall = trieNode
	queue = append(queue, trieNode)
	isEmpty := false
	for isEmpty == false {
		node := queue[0]
		queue = queue[1:]
		for _, currentNode := range node.childNodes {
			queue = append(queue, currentNode)
		}
		if node != trieNode {
			fall := node.parent.fall
			for check(fall, trieNode, node.char) == false {
				fall = fall.fall
			}
			node.fall = fall.getChild(node.char)
			if node.fall == nil {
				node.fall = trieNode
			}
			if node.fall == node {
				node.fall = trieNode
			}
		}
		if (len(queue) == 0) {
			isEmpty = true
		}
	}
	return trieNode
}

func findPhrase(trieNode *TrieNode) string {
	var tmp []rune
	currentNode := trieNode
	isEnd := trieNode.parent != nil
	for isEnd == true {
		tmp = append(tmp, currentNode.char)
		currentNode = currentNode.parent
		isEnd = currentNode.parent != nil
	}
	return string(reverse(tmp))
}

type Result struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

func search(str string, trieNode *TrieNode) []Result {
	var results []Result
	chars := []rune(str)
	currentState := trieNode
	for i, char := range chars {
		node := currentState
		for node.char != 0 && node.getChild(char) == nil {
			node = node.fall
		}
		node = node.getChild(char)
		if node == nil {
			node = trieNode
		}
		currentNode := node
		for currentNode.char != 0 {
			if currentNode.isEnd == true {
				phrase := findPhrase(currentNode)
				index := 1 + i - utf8.RuneCountInString(phrase)
				results = append(results, Result{index, phrase})
			}
			currentNode = currentNode.fall
		}
		currentState = node
	}
	return results
}

type aho struct {
	trieNode *TrieNode
}

func NewSearch(phrases []string) aho {
	a := aho{}
	a.trieNode = &TrieNode{}
	a.trieNode.childNodes = make(map[rune]*TrieNode)
	for _, phrase := range phrases {
		a.trieNode.addString(phrase)
	}
	return a
}

func (a aho) Add(phrase string) {
	a.trieNode.addString(phrase)
}

func (a aho) Build() {
	a.trieNode = build(a.trieNode)
}

func (a aho) Exec(phrase string) []Result {
	return search(phrase, a.trieNode)
}
