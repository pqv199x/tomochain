package tomox_state

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
	"testing"
)

func TestTomoxTrieTest(t *testing.T) {
	db, _ := ethdb.NewMemDatabase()
	stateCache := NewDatabase(db)
	trie, _ := stateCache.OpenStorageTrie(EmptyHash, EmptyHash)
	min := common.BigToHash(big.NewInt(1)).Bytes()
	max := common.BigToHash(big.NewInt(2)).Bytes()
	trie.TryUpdate(min, min)
	trie.TryUpdate(max, max)
	left, _, _ := trie.TryGetBestLeftKeyAndValue()
	right, _, _ := trie.TryGetBestRightKeyAndValue()
	fmt.Println(left, right)
	trie.TryDelete(min)
	left, _, _ = trie.TryGetBestLeftKeyAndValue()
	right, _, _ = trie.TryGetBestRightKeyAndValue()
	fmt.Println(left, right)
}
