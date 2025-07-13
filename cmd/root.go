package cmd

import (

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "updog",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Aquí cobra-cli no registra comandos.
	// Cada subcomando generado (con `cobra-cli add`) lo hace por separado en su archivo.
}

