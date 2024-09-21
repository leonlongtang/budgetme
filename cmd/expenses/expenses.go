package expenses

import (
	"budgetme/sqldb"
	"budgetme/utils" // Ensure this points to your utils package
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	db        *sql.DB
	direction string
	orderBy   string
	log       *logrus.Logger
)

// expensesCmd represents the expenses command
var ExpensesCmd = &cobra.Command{
	Use:   "expenses",
	Short: "Handling Expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Error("Error displaying help: ", err)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log = utils.GetLogger() // Initialize the logger
		var err error
		db, err = sqldb.InitDB() // Initialize the DB at the root level
		if err != nil {
			log.Fatal("Failed to initialize the database: ", err)
			os.Exit(1)
		}
		direction = viper.GetString("direction")
		orderBy = viper.GetString("order_by")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			db.Close() // Close the DB connection when the CLI app finishes
			log.Info("Database connection closed successfully.")
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
