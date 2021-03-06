package cmd

import (
	"log"
	"strings"

	"github.com/mdesson/gophercise-4_todo-cli/taskdb"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  `Add a task to your to todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := taskdb.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		taskdb.AddTask(db, strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
