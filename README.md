# Trie

---

### Overview
Designed for keyword retrieval and searching

### How to get
```
go get github.com/vearne/trie
```

### Usage
Building a trie tree
```
NewTrie() *Trie
(t *Trie) Add(keyword string, meta interface{})
```
`meta` 
Additional information associated with the keyword
   
##### 1) try to find one keyword
```
(t *Trie) Query(text string) (result *HitResult, ok bool)
```
##### 2) find all keyword
```
(t *Trie) QueryAll(text string) []HitResult
```

HitResut
```
type HitResult struct {
	Keyword string
	Meta    interface{}
	Pos     int
}
```

### Example
```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/vearne/trie"
)

func main() {
    // init tree
	tree := trie.NewTrie()
	tree.Add("our story", 1)
	tree.Add("our", 2)
	tree.Add("hello", 3)
	tree.Add("too short", 4)
	tree.Add("world", "hit word")

	text := "We get to decide what our story is. Nobody else gets to tell you what your story is"
	// try to find one keyword
	// Shortest match, so return "our",  but "our story"
	result, ok := tree.Query(text)
	if ok {
		fmt.Println("--Query--")
		fmt.Println(result.Keyword, result.Meta, result.Pos)
	}
	// find all keyword
	itemList := tree.QueryAll(text)
	buf, _ := json.Marshal(itemList)
	fmt.Println("--QueryAll--")
	fmt.Println(string(buf))
}
```
### Result
```
--Query--
our 2 22
--QueryAll--
[{"Keyword":"our","Meta":2,"Pos":22},{"Keyword":"our story","Meta":1,"Pos":22},{"Keyword":"our","Meta":2,"Pos":71},{"Keyword":"our story","Meta":1,"Pos":71}]
```


### Reference
1. [trie](https://en.wikipedia.org/wiki/Trie)





