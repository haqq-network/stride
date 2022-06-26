package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/Stride-Labs/stride/x/interchainquery/types"
	"github.com/spf13/cobra"
)

// GetTxCmd creates and returns the intertx tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		// TODO(TEST-53) remove cli access to ICQ queries pre-launch
		SubmitQueryResponse(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func SubmitQueryResponse() *cobra.Command {
	cmd := &cobra.Command{
		// TODO(TEST-77) camel case
		Use:   "submitqueryresponse [chain_id] [query_id] [result] [height] [from_address]",
		Short: `Submit Query Response.`,
		Long: `
		e.g. "submitqueryresponse GAIA_1 8e3451aea1ca8438f4ba9292a3add814d50d45d163ff28f859836cbb074584c0 ChUKBXVhdG9tEgw0OTkwMDAwMDAwMDASAhAB 40 stride1wlgadk2gndm96tvf0v6207jckqu8e2huyfhsp5"`,
		Example: `submitqueryresponse GAIA_1 cosmos1pcag0cj4ttxg8l7pcg0q4ksuglswuuedcextl2 uatom`,
		Args:    cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			chain_id := args[0]
			// query_id := args[1]
			result := args[2]
			// height := args[3]
			from_address := args[4]
			from_addr_sdk, _ := sdk.AccAddressFromBech32(from_address)
			caller := clientCtx.GetFromAddress().String()

			// TODO cleanup
			if len(caller) < 1 {
				return fmt.Errorf("Error: empty --from address.")
			}
			fmt.Println(caller)

			// TODO(TEST-50) create message based on parsed json
			msg := types.NewMsgSubmitQueryResponse(chain_id, result, from_addr_sdk)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionID)
	_ = cmd.MarkFlagRequired(FlagConnectionID)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}