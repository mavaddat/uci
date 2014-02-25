package uci

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type UCISuite struct{}

var _ = Suite(&UCISuite{})

// these are somewhat fragile and depend on having stockfish in the folder
// geared toward CI

func (s *UCISuite) TestUCIDepth4(c *C) {
	eng, err := NewEngine("./stockfish")
	c.Assert(err, IsNil)
	eng.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")
	eng.SetOptions(Options{MultiPV: 4})
	resultOpts := HighestDepthOnly | IncludeUpperbounds | IncludeLowerbounds
	res, err := eng.GoDepth(10, resultOpts)
	c.Assert(err, IsNil)
	c.Assert(res.BestMove, Equals, "c8d7")
	c.Assert(len(res.Results), Equals, 4)
}

func (s *UCISuite) TestUCIDepth15(c *C) {
	eng, err := NewEngine("./stockfish")
	c.Assert(err, IsNil)
	eng.SetFEN("r1b1k1nr/ppq2pbp/2n1p1p1/1B2pN2/5P2/2N1B3/PPP3PP/R2QK2R w KQkq - 2 11")
	res, err := eng.GoDepth(19)
	c.Assert(err, IsNil)
	c.Assert(res.BestMove, Equals, "b5c6")
}
