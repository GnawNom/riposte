package store

import "github.com/DataWraith/vptree"

type HashStore interface {
	FindDuplicate(hash PerceptionHash, threshold uint64)
	Store(hash PerceptionHash)
}

type PerceptionHash struct {
	Value uint64
}

func PerceptionHashMetric(a, b interface{}) float64 {
	c1 := a.(PerceptionHash)
	c2 := b.(PerceptionHash)
	hamming := c1.Value ^ c2.Value
	return float64(popcnt(hamming))
}

type VPTreeHashStore struct {
	Name   string
	Buffer []PerceptionHash
	Tree   *vptree.VPTree
}

func popcnt(x uint64) int {
	diff := 0
	for x != 0 {
		diff += int(x & 1)
		x >>= 1
	}

	return diff
}

func NewVPTreeHashStore(name string) *VPTreeHashStore {
	return &VPTreeHashStore{Name: name, Buffer: make([]PerceptionHash, 0, 100), Tree: nil}
}

func (s *VPTreeHashStore) Store(hash PerceptionHash) {
	if len(s.Buffer) == cap(s.Buffer) {
		s.buildTree(s.Buffer)
		s.Buffer = make([]PerceptionHash, 0, 100)
	}
	s.Buffer = append(s.Buffer, hash)
}

func (s *VPTreeHashStore) buildTree(hashes []PerceptionHash) {
	vpitems := make([]interface{}, len(hashes))
	for i, v := range hashes {
		vpitems[i] = interface{}(v)
	}
	s.Tree = vptree.New(PerceptionHashMetric, vpitems)
}

func (s *VPTreeHashStore) FindDuplicate(target PerceptionHash, threshold float64) bool {
	for _, entry := range s.Buffer {
		distance := PerceptionHashMetric(target, entry)
		if distance <= threshold {
			return true
		}
	}
	if s.Tree != nil {
		results, _ := s.Tree.Search(target, 1)
		if len(results) > 0 {
			return true
		}
	}
	return false
}
