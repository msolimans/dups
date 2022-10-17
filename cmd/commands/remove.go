package commands

import (
	"github.com/msolimans/dups/pkg/dups"
	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	var (
		source 		string	
		destination string 
		concurrent 		bool  
		pretty 		bool  
		out 		bool 
	)

	command := &cobra.Command{
		Use:   "remove",
		Short: "Remove duplicates",
		Run: func(cmd *cobra.Command, args []string) {
			dups.Start(source, destination, concurrent, pretty, out)
		},
	}

	//define flags
	command.Flags().StringVarP(&source, "src", "s", "", "Source file path")
	command.Flags().StringVarP(&destination, "dest", "d", "clean_application.json", "Destination file path")
	command.Flags().BoolVarP(&pretty, "pretty", "p", true, "Enable pretty format, adding indentation to json output")
	command.Flags().BoolVarP(&concurrent, "concurrent", "c", false, "Enable concurrency")
	command.Flags().BoolVarP(&out, "out", "o", false, "Output summary in stdout")

	
	//mark flags as required
	command.MarkFlagRequired("src")
	return command
}
