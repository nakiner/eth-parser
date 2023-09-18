package cache

const currentBlockKey = "current_block"

func (p *Provider) SetCurrentBlock(block int64) {
	p.currentBlock.Set(currentBlockKey, block)
}

func (p *Provider) GetCurrentBlock() int64 {
	val, ok := p.currentBlock.Get(currentBlockKey)
	if !ok {
		return 0
	}
	return val.(int64)
}
