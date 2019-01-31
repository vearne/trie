package trie

type Node struct {
	val      rune
	meta     interface{}
	hit      bool
	children map[rune]*Node
}

type Trie struct {
	root *Node
	size int
}

type HitResult struct {
	Keyword string
	Meta    interface{}
	Pos     int
}

func NewTrie() *Trie {
	t := Trie{size: 0}
	t.root = &Node{val: -1, children: make(map[rune]*Node)}
	return &t
}

func (t *Trie) Add(keyword string, meta interface{}) {
	node := t.root
	runes := []rune(keyword)
	for _, ch := range runes {
		_, ok := node.children[ch]
		if !ok {
			node.children[ch] = &Node{val: ch, meta: nil,
				hit:      false,
				children: make(map[rune]*Node)}
		}
		// next level
		node = node.children[ch]
	}
	node.hit = true
	node.meta = meta
	t.size += 1
}

func (t *Trie) Size() int {
	return t.size
}

// keyword word check
func (t *Trie) Query(text string) (result *HitResult, ok bool) {
	var ch rune
	runes := []rune(text)

	for i, _ := range runes {
		node := t.root
		j := i
		for ; j < len(runes); j++ {
			ch = runes[j]
			_, ok = node.children[ch]
			if ok {
				node = node.children[ch]
			}
			if ok && ch == node.val {
				//fmt.Println(string(ch), "i", i, "j", j, string(node.val), node.meta)
				if node.hit {
					result = &HitResult{Keyword: string(runes[i : j+1]),
						Meta: node.meta, Pos: i}
					return result, true
				}
			} else {
				break
			}
		}
	}

	return nil, false
}

func (t *Trie) QueryAll(text string) []HitResult {
	result := make([]HitResult, 0, 10)
	var ok bool

	var ch rune
	runes := []rune(text)

	for i, _ := range runes {
		node := t.root
		j := i
		for ; j < len(runes); j++ {
			ch = runes[j]
			_, ok = node.children[ch]
			if ok {
				node = node.children[ch]
			}
			if ok && ch == node.val {
				//fmt.Println(string(ch), "i", i, "j", j, string(node.val), node.meta)
				if node.hit {
					result = append(result, HitResult{Keyword: string(runes[i : j+1]),
						Meta: node.meta, Pos: i})
				}
			} else {
				break
			}
		}
	}

	return result
}
