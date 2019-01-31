package trie 

import (
	//"encoding/json"
	//"fmt"
	"encoding/json"
	"fmt"
	"testing"
)

var tree *Trie

func init() {
	tree = NewTrie()
	tree.Add("our story", 1)
	tree.Add("our", 2)
	tree.Add("hello", 3)
	tree.Add("too short", 4)
	tree.Add("world", "hit word")
}

const (
	Text1 = "We get to decide what our story is. Nobody else gets to tell you what your story is"
	Text2 = "Life is like an ice cream, enjoy it before it melts"
	Text3 = "Hello world"
	Text4 = "It seems like life is too short to hold grudges."
)

func TestQuery1(t *testing.T) {
	result, ok := tree.Query(Text1)
	target := "our"
	if ok {
		t.Logf("success, expect:%v, got:%v", target, result.Keyword)
	} else {
		t.Errorf("error, expect:%v, got:%v", target, result.Keyword)
	}
}

func TestQuery2(t *testing.T) {
	result, ok := tree.Query(Text4)
	target := "too short"
	if ok {
		t.Logf("success, expect:%v, got:%v", target, result.Keyword)
	} else {
		t.Errorf("error, expect:%v, got:%v", target, result.Keyword)
	}
}

func TestQuery3(t *testing.T) {
	result, ok := tree.Query(Text3)
	target := "world"
	if ok && result.Pos == 6 {
		t.Logf("success, expect:%v, got:%v", target, result.Keyword)
	} else {
		t.Errorf("error, expect:%v, got:%v", target, result.Keyword)
	}
}

func TestQuery4(t *testing.T) {
	_, ok := tree.Query(Text2)
	if !ok {
		t.Logf("success, expect:%v, got:%v", false, ok)
	} else {
		t.Errorf("error, expect:%v, got:%v", false, ok)
	}
}

func TestQuery5(t *testing.T) {
	itemList := tree.QueryAll(Text1)
	buf, _ := json.Marshal(itemList)
	fmt.Println(string(buf))
	count := len(itemList)
	if count == 4 {
		t.Logf("success, expect:%v, got:%v", 4, count)
	} else {
		t.Errorf("error, expect:%v, got:%v", 4, count)
	}
}
