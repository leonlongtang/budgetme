package expenses

import (
	"budgetme/sqldb"
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	db        *sql.DB
	direction string
	orderBy   string
)

// expensesCmd represents the expenses command
var ExpensesCmd = &cobra.Command{
	Use:   "expenses",
	Short: "Handling Expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			// You can either log the error or handle it in some meaningful way
			fmt.Println("Error displaying help:", err)
		}

	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = sqldb.InitDB() // Initialize the DB at the root level
		if err != nil {
			fmt.Println("Failed to initialize the database:", err)
			os.Exit(1)
		}
		direction = viper.GetString("direction")
		orderBy = viper.GetString("order_by")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			db.Close() // Close the DB connection when the CLI app finishes
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expensesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expensesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
