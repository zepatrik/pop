package cmd

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/cobra"
)

var migrationStepUp int

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply one or more of the 'up' migrations.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mig, err := pop.NewFileMigrator(migrationPath, getConn())
		if err != nil {
			return err
		}
		return mig.UpTo(migrationStepUp)
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateUpCmd.Flags().IntVarP(&migrationStepUp, "step", "s", 0, "Number of migrations to apply. Use 0 to apply all pending.")
}
