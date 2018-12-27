package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NewCmdUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Conrtol Slack users",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("users not implemented")
		},
	}

	cmd.AddCommand(SubCmdUserList())
	return cmd
}

func SubCmdUserList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return all user lists",
		Run: func(cmd *cobra.Command, args []string) {
			users, _ := api.GetUsers()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "RealName", "ID", "IsBot", "IsAdmin"})
			table.SetBorder(false)
			table.SetRowLine(false)

			for _, user := range users {
				d := []string{
					user.Name,
					user.RealName,
					user.ID,
					fmt.Sprintf("%t", user.IsBot),
					fmt.Sprintf("%t", user.IsAdmin),
				}
				table.Append(d)
			}
			table.Render()
		},
	}
	return cmd
}
