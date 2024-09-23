package expenses

import (
	"budgetme/sqldb"
	"budgetme/utils" // Ensure this points to your utils package
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	db     *sqldb.Database
	log    *logrus.Logger
	config string
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
		log = utils.GetLogger()

		if config != "" {
			viper.SetConfigFile(config)
			err := viper.ReadInConfig() // Find and read the config file
			if err != nil {
				log.Fatalf("Error reading config file: %v", err)
				os.Exit(1)
			}
			log.Infof("Config file loaded: %s", config)
		} else {
			log.Info("No config file provided, proceeding without it.")
		}

		// Initialize the DB connection as part of the Database struct
		var err error
		db, err = sqldb.InitDB(log)
		if err != nil {
			log.Fatal("Failed to initialize the database: ", err)
			os.Exit(1)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db.DB != nil {
			db.DB.Close()
			db.Log.Info("Database connection closed successfully.")
		}
	},
}

func init() {
	// Add a global config flag for all subcommands
	ExpensesCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "Path to config file")
}
