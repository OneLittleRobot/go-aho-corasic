package aho

type TrieNode struct {
	childNodes map[rune]*TrieNode
	char       rune
	isEnd      bool
	parent     *TrieNode
	fall       *TrieNode
}

func (trieNode *TrieNode) getChild(char rune) *TrieNode {
	return trieNode.childNodes[char]
}

func (trieNode *TrieNode) addString(str string) *TrieNode {
	chars := []rune(str)
	return trieNode.add(chars)
}

func (trieNode *TrieNode) add(chars []rune) *TrieNode {
	char, chars := chars[0], chars[1:]

	currentNode := trieNode.getChild(char)

	if currentNode == nil {
		currentNode = &TrieNode{char: char, parent: trieNode}
		currentNode.childNodes = make(map[rune]*TrieNode)
		trieNode.childNodes[char] = currentNode
	}

	if len(chars) > 0 {
		currentNode = currentNode.add(chars)
	} else {
		currentNode.isEnd = true
	}
	return trieNode
}