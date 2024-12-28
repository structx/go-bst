package bella_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/bella"
)

type BstSuite struct {
	suite.Suite
	btree *bella.BST
}

func (suite *BstSuite) SetupSuite() {
	suite.btree = bella.New()
}

func (suite *BstSuite) TestInsert() {
	suite.btree.Insert("hello", []byte("world"))
}

func (suite *BstSuite) TestSearch() {

	assert := suite.Assert()

	suite.btree.Insert("hello", []byte("world"))
	assert.Greater(suite.btree.Size(), int64(0))

	value, err := suite.btree.Search("hello")
	assert.NoError(err)

	assert.Equal([]byte("world"), value)
}

func (suite *BstSuite) TestEmpty() {

	assert := suite.Assert()

	suite.btree.Empty()

	assert.Equal(int64(0), suite.btree.Size())
}

func TestBstSuite(t *testing.T) {
	suite.Run(t, new(BstSuite))
}
