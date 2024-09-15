# trie datastructure

- this datastructure is optomized to find words in words (use by autocomplete)

## references

- <https://youtu.be/H-6-8_p88r0>

## document status

- complete: NO
- tested: NO

---

## Implementing a Trie Data Structure in Go

A trie, also known as a prefix tree, is a tree-based data structure that is commonly used to store and retrieve strings efficiently, especially when dealing with a large set of strings. Each node in a trie represents a character of a string, and paths down the tree represent prefixes of words.

## Table of Contents

- [trie datastructure](#trie-datastructure)
  - [references](#references)
  - [document status](#document-status)
  - [Implementing a Trie Data Structure in Go](#implementing-a-trie-data-structure-in-go)
  - [Table of Contents](#table-of-contents)
  - [Introduction to Tries](#introduction-to-tries)
  - [Applications of Tries](#applications-of-tries)
  - [Implementing a Trie in Go](#implementing-a-trie-in-go)
    - [Trie Node Structure](#trie-node-structure)
    - [Trie Structure](#trie-structure)
    - [Insert Method](#insert-method)
    - [Search Method](#search-method)
    - [StartsWith Method](#startswith-method)
  - [Complete Implementation](#complete-implementation)
  - [Usage Example](#usage-example)
  - [Conclusion](#conclusion)

## Introduction to Tries

A trie is a special type of tree used to store associative data structures. Unlike a binary search tree, no node in the tree stores the key associated with that node; instead, its position in the tree defines the key with which it is associated. All the descendants of a node have a common prefix of the string associated with that node.

## Applications of Tries

- **Autocomplete Systems**: Predicting the rest of a word based on the prefix typed.
- **Spell Checkers**: Finding words with a common prefix or correcting misspelled words.
- **IP Routing**: Longest prefix matching in routing tables.
- **Database Engines**: Efficient storage and retrieval of strings.

## Implementing a Trie in Go

Let's implement a simple trie in Go that supports insertion of words, searching for a full word, and checking if any words start with a given prefix.

### Trie Node Structure

Each node in the trie will represent a character and have a map of its children nodes. We'll also keep a boolean flag to indicate if a node marks the end of a word.

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}
```

### Trie Structure

The trie itself will have a root node, which is an empty `TrieNode`.

```go
type Trie struct {
    root *TrieNode
}
```

### Insert Method

The `Insert` method adds a word to the trie by iterating through each character and creating nodes as needed.

```go
func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}
```

### Search Method

The `Search` method checks if a word exists in the trie.

```go
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}
```

### StartsWith Method

The `StartsWith` method checks if there's any word in the trie that starts with the given prefix.

```go
func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return true
}
```

## Complete Implementation

Here's the complete code with all the components together:

```go
package main

import (
    "fmt"
)

// TrieNode represents each node in the trie.
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

// Trie represents the trie data structure.
type Trie struct {
    root *TrieNode
}

// NewTrie initializes and returns a new trie.
func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{children: make(map[rune]*TrieNode)},
    }
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

// Search checks if a word exists in the trie.
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}

// StartsWith checks if any word in the trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return true
}

func main() {
    trie := NewTrie()
    trie.Insert("hello")
    trie.Insert("helium")

    fmt.Println(trie.Search("hello"))    // Output: true
    fmt.Println(trie.Search("helix"))    // Output: false
    fmt.Println(trie.StartsWith("hel"))  // Output: true
    fmt.Println(trie.StartsWith("hey"))  // Output: false
}
```

## Usage Example

In the `main` function, we:

1. Initialize a new trie.
2. Insert words "hello" and "helium".
3. Search for words and prefixes.

**Output:**

```sh
true
false
true
false
```

## Conclusion

Tries are powerful data structures for efficient retrieval of strings based on prefixes. The implementation in Go is straightforward, leveraging maps and runes to handle Unicode characters properly. This basic trie can be extended with additional features like delete operations or storing values at nodes.