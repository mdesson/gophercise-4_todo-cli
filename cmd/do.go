package cmd

import (
	"log"
	"strconv"

	"github.com/mdesson/gophercise-4_todo-cli/taskdb"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do a task",
	Long:  `Complete a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := taskdb.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		key, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		taskdb.CompleteTask(db, key)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
