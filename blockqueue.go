package DAG_Rider

import "fmt"

type BlockQueue struct {
	blocks []Block
}

func (q *BlockQueue) Empty() bool {
	return len(q.blocks) == 0
}

func (q *BlockQueue) Push(block *Block) error {
	q.blocks = append(q.blocks, *block)
	return nil
}

func (q *BlockQueue) Pop() (block *Block, err error) {
	if q.Empty() {
		block = nil
		err = fmt.Errorf("poping an empty BlockQueue")
		return
	}
	*block = q.blocks[0]
	q.blocks = q.blocks[1:]
	err = nil
	return
}

func (q *BlockQueue) Front() (*Block, error) {
	if q.Empty() {
		return nil, fmt.Errorf("taking the front element from an empty BlockQueue")
	}
	return &(q.blocks[0]), nil
}