//go:build unix

package uci

import (
	. "gopkg.in/check.v1"
)

func (s *UCISuite) TestNewEngineNice(c *C) {
	// Test creating engine with nice level 10 (lower priority)
	eng, err := NewEngineNice(10, "./stockfish")
	c.Assert(err, IsNil)
	defer eng.Close()

	// Verify engine is functional
	err = eng.UCI()
	c.Assert(err, IsNil)
}

func (s *UCISuite) TestSetNice(c *C) {
	eng, err := NewEngine("./stockfish")
	c.Assert(err, IsNil)
	defer eng.Close()

	// Set nice level after engine is running
	err = eng.SetNice(10)
	c.Assert(err, IsNil)

	// Verify engine is still functional
	err = eng.UCI()
	c.Assert(err, IsNil)
}

func (s *UCISuite) TestNewEngineNiceBadPath(c *C) {
	_, err := NewEngineNice(10, "/bad/path/to/engine")
	c.Assert(err, NotNil)
}
