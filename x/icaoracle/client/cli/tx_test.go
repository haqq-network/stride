package cli_test

import (
	"github.com/Stride-Labs/stride/v22/x/icaoracle/client/cli"
)

func (s *ClientTestSuite) TestCmdRestoreOracleICA() {
	args := []string{
		HostChainId,
	}

	cmd := cli.CmdRestoreOracleICA()
	s.ExecuteTxAndCheckSuccessful(cmd, args, "restore oracle ICA")
}
