package bst_test

import (
	"os"
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
	n := suite.btree.Insert("5", []byte("hellworld"))
	suite.NotNil(n)
	n = suite.btree.Insert("3", []byte("hellworld"))
	suite.NotNil(n)
	n = suite.btree.Insert("7", []byte("hellworld"))
	suite.NotNil(n)
	n = suite.btree.Insert("9", []byte("hellworld"))
	suite.NotNil(n)
	n = suite.btree.Insert("1", []byte("hellworld"))
	suite.NotNil(n)

	i := 0

	suite.btree.InOrderTraversal(func(n *bst.Node[string, []byte]) error {

		switch i {
		case 0:
			suite.Equal("1", n.Key)
		case 1:
			suite.Equal("3", n.Key)
		case 2:
			suite.Equal("5", n.Key)
		case 3:
			suite.Equal("7", n.Key)
		case 4:
			suite.Equal("9", n.Key)
		}

		i++

		return nil
	})
}

func (suite *BstSuite) TestToFile() {
	suite.btree.Insert("helloworld", []byte("helloworld"))
	err := suite.btree.ToFile("testfiles/to_file.data")
	suite.NoError(err)
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

func (suite *BstSuite) TestFromFile() {
	btree, err := bst.FromFile[string, []byte]("testfiles/test.data")
	suite.NoError(err)
	suite.Greater(btree.Size(), uintptr(0))
}

func (suite *BstSuite) TearDownSuite() {
	_ = os.Remove("testfiles/to_file.data")
}

func TestBstSuite(t *testing.T) {
	suite.Run(t, new(BstSuite))
}
