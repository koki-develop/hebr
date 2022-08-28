package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/koki-develop/hebr"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "cmd",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var src []byte
		if len(args) == 0 {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			src = data
		} else {
			filename := args[0]
			f, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			data, err := io.ReadAll(f)
			if err != nil {
				return err
			}
			src = data
		}

		dec, err := cmd.Flags().GetBool("decode")
		if err != nil {
			return err
		}

		if dec {
			decoded, err := hebr.Decode([]byte(strings.TrimRightFunc(string(src), unicode.IsSpace)))
			if err != nil {
				return err
			}
			fmt.Print(string(decoded))
		} else {
			encoded, err := hebr.Encode(src)
			if err != nil {
				return err
			}
			fmt.Print(string(encoded))
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("decode", "d", false, "decodes data")
}
