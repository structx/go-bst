package bst_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/go-bst"
)

type BstSuite struct {
	suite.Suite
	btree *bst.BtreeGN[string, []byte]
}

func (suite *BstSuite) SetupSuite() {
	suite.btree = &bst.BtreeGN[string, []byte]{}
}

func (suite *BstSuite) TestInsert() {
	suite.btree.Insert("hello", []byte("world"))
}

func (suite *BstSuite) TestSearch() {

	suite.btree.Insert("hello", []byte("world"))
	suite.Greater(suite.btree.Size(), uintptr(0))

	value, err := suite.btree.Search("hello")
	suite.NoError(err)

	suite.Equal([]byte("world"), value)

	_, err = suite.btree.Search("missing")
	suite.Equal(bst.ErrNotFound, err)
}

func (suite *BstSuite) TestFlush() {

	suite.btree.Insert("rick", []byte("morty"))
	suite.GreaterOrEqual(suite.btree.Size(), uintptr(24))

	suite.btree.Flush()
	suite.Equal(uintptr(0), suite.btree.Size())
}

func TestBstSuite(t *testing.T) {
	suite.Run(t, new(BstSuite))
}
