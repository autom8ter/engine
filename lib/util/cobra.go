package util

import "github.com/spf13/cobra"

type CobraFunc func(cmd *cobra.Command, args []string)

func (c CobraFunc) Cobra(use, short, long string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   c,
	}
}
