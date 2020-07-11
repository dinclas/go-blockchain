package blockchain

import "testing"

func TestBlockchain_AddBlock(t *testing.T) {
	bc := New()
	bc.AddBlock("foo")

	if len(bc.Blocks) != 2 {
		t.Errorf("Blockchain is expected to have 2 blocks. It has %d", len(bc.Blocks))
	}

	if bc.Blocks[len(bc.Blocks)-1].Data != "foo" {
		t.Errorf("Blockchain's last block data expected to be \"foo\". Got \"%s\"", bc.Blocks[len(bc.Blocks)-1].Data)
	}
}

func TestBlockchain_ReplaceChain(t *testing.T) {
	firstBc := New()
	secondBc := New()

	firstBc.AddBlock("foo")
	secondBc.ReplaceChain(firstBc.Blocks)

	if len(secondBc.Blocks) != 2 {
		t.Errorf("Blockchain is expected to have 2 blocks. It has %d", len(secondBc.Blocks))
	}

	secondBc.AddBlock("bar")
	firstBc.ReplaceChain(secondBc.Blocks)

	if len(firstBc.Blocks) != 3 {
		t.Errorf("Blockchain is expected to have 3 blocks. It has %d", len(firstBc.Blocks))
	}

	thirdBc := New()
	secondBc.ReplaceChain(thirdBc.Blocks)

	if len(firstBc.Blocks) != 3 {
		t.Errorf("Blockchain is expected to have 3 blocks. It has %d", len(secondBc.Blocks))
	}
}