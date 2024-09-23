package expenses

import (
	"budgetme/services"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ExpenseRaw struct {
	Amount   float64
	Category string
	Date     string
}

var (
	amount   float64
	category string
	date     string
	config   string
)

// createNewCmd represents the createNew command
var createNewCmd = &cobra.Command{
	Use:   "createNew",
	Short: "Adds a new expense in the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the service layer
		expenseService := services.NewExpenseService(db)

		if config != "" {
			log.Info("Loading expenses from config file")

			// Set the config file path for viper
			viper.SetConfigFile(config)

			// Read the configuration file
			if err := viper.ReadInConfig(); err != nil {
				log.Error("Failed to read config file:", err)
				fmt.Println("Error: could not load config file")
				return
			}

			// Debugging: Print the entire config to check if it was loaded correctly
			fmt.Println("Raw config data:", viper.AllSettings()) // This will print the raw configuration

			var expenses []ExpenseRaw
			err := viper.UnmarshalKey("createNew.expenses", &expenses)
			if err != nil {
				log.Error("Failed to load expenses from config:", err)
				return
			}

			// Check if expenses were loaded correctly
			fmt.Println("Loaded expenses:", expenses)

			// Add each expense from the config file
			for _, exp := range expenses {
				log.Infof("Adding expense from config: %+v", exp)
				err := expenseService.AddExpense(exp.Amount, exp.Category, exp.Date)
				if err != nil {
					log.Errorf("Failed to add expense from config: %+v, error: %v", exp, err)
				} else {
					fmt.Println("Added expense:", exp)
				}
			}
		} else {
			// Validate required flags
			if amount == 0 || category == "" || date == "" {
				log.Error("Missing required flags: amount, category, or date")
				fmt.Println("Error: missing required flags --amount, --category, and --date")
				return
			}

			// Add a single expense using the flags
			log.Infof("Adding expense from flags: amount=%f, category=%s, date=%s", amount, category, date)
			err := expenseService.AddExpense(amount, category, date)
			if err != nil {
				log.Errorf("Failed to add expense from flags: %v", err)
			} else {
				fmt.Printf("Added expense: Amount=%.2f, Category=%s, Date=%s\n", amount, category, date)
			}
		}
	},
}

func init() {
	createNewCmd.Flags().StringVarP(&config, "file", "f", "", "Path to config file with expenses")
	createNewCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Expense amount")
	createNewCmd.Flags().StringVarP(&category, "category", "c", "", "Expense category")
	createNewCmd.Flags().StringVarP(&date, "date", "d", "", "Expense date (YYYY-MM-DD)")
	ExpensesCmd.AddCommand(createNewCmd)
}
