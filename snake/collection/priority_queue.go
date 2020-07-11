package collection

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/mazrean/SnakeGame/snake/heuristic"
)

// PriorityQueue 優先度付きキューの構造体
type PriorityQueue struct {
	nodes *priorityNodes
	HeuristicFunc heuristic.HeuristicFunc
}

// NewPriorityQueue ProprityQueueのコンストラクタ
func NewPriorityQueue(f heuristic.HeuristicFunc) *PriorityQueue {
	nds := new(priorityNodes)
	heap.Init(nds)

	priorityQueue := &PriorityQueue{
		nodes: nds,
		HeuristicFunc: f,
	}

	return priorityQueue
}

// Push 要素の追加
func (p *PriorityQueue) Push(n *Node) error {
	h, err := p.HeuristicFunc(n.State, n.Direction)
	if err != nil {
		return fmt.Errorf("PriorityFunc Error: %w", err)
	}

	node := &nodeWithPriority{
		node: n,
		priority: h + n.State.Turn,
	}

	heap.Push(p.nodes, node)

	return nil
}

// Pop 値の取り出し
func (p *PriorityQueue) Pop() (*Node, error) {
	if len(*p.nodes) == 0 {
		return nil, errors.New("Empty Priority Queue")
	}

	interfaceNode := heap.Pop(p.nodes)
	
	node, ok := interfaceNode.(*Node)
	if !ok {
		return nil, errors.New("Node Parse Error")
	}

	return node, nil
}

// Size サイズ
func (p *PriorityQueue) Size() (int,error) {
	return len(*p.nodes), nil
}

// Empty 空か
func (p *PriorityQueue) Empty() (bool, error) {
	return len(*p.nodes) == 0, nil
}

type nodeWithPriority struct {
	node *Node
	priority int
}

type priorityNodes []*nodeWithPriority

func (p *priorityNodes) Swap(i,j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityNodes) Push(node interface{}) {
	value := node.(*nodeWithPriority)
	*p = append(*p, value)
}

func (p *priorityNodes) Pop() interface{} {
	n := len(*p)
	node := (*p)[n-1].node
	(*p)[n-1] = nil
	*p = (*p)[0 : n-1]
	return node
}

func (p *priorityNodes) Less(i,j int) bool {
	return (*p)[i].priority < (*p)[j].priority
}

func (p *priorityNodes) Len() int {
	return len(*p)
}