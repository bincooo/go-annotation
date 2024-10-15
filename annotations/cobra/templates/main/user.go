//go:build user

// Example: no need to do with this something. This just provided example of desired generated code
// The file and above tag is generated by Cobra(build)
package main

/*
import (
	// Added with command executor
	"github.com/YReshetko/go-annotation/examples/cobra/commands"

	// Generated by default
	"github.com/spf13/cobra"
)

func init() {
	//####################### BLOCK ##########################//
	// generated with no :=, but =
	root = &cobra.Command{}
	// Generate from build commands tree Cobra(usage)
	root.Use = "cli"
	// Generate from build commands tree Cobra(example)
	root.Example = "cli [-F file | -D dir] ... [-f format] profile"
	// Generate from Cobra(short)
	root.Short = "Root command of the application (short)"
	// Generate from Cobra(long)
	root.Long = "Root command of the application (long)"

	// Generated by RootCommand fields with 'flag' tag with no 'persistent' adds
	// Type is resolved by field type
	// 'P' suffix is resolved by 'short' tag
	// Default value is empty is missing, so default type value
	// Description is resolved by 'description' tag
	root.Flags().StringP("output", "o", "", "output file name")
	// Generated by flag adds 'required'
	if err := root.MarkFlagRequired("output"); err != nil {
		fatal(err)
	}
	root.Flags().Int("num", 42, "some number for command")
	// Generated by RootCommand fields with 'flag' and 'persistent' adds
	root.PersistentFlags().BoolP("is-ok", "i", true, "some persistent flag")

	// Generated because @CobraRun (always use RunE, because we need to track parse and other errors)
	root.RunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.RootCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.Run(cmd, args)
	}

	// Generated because @CobraPreRun (always use PreRunE, because we need to track parse and other errors)
	root.PreRunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.RootCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.PreRun(cmd, args)
	}

	// Generated because @CobraPersistPreRun (always use PersistentPreRunE, because we need to track parse and other errors)
	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.RootCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.PersistPreRun(cmd, args)
	}
	// DO NOT NEED to add the root command to any other commands
	// root.AddCommand(root)

	//####################### BLOCK ##########################//
	// generated with :=, but not =
	child := &cobra.Command{}

	// Generate from build commands tree Cobra(usage) we take the last word from usage, and select the target command by previous
	child.Use = "get"
	// Generate from build commands tree Cobra(example)
	child.Example = "cli get [-F file | -D dir] ... [-f format] profile"
	// Generate from Cobra(short)
	child.Short = "Child command of the application (short)"
	// Generate from Cobra(long)
	child.Long = "Child command of the application (long)"

	// Generated by RootCommand fields with 'flag' tag with no 'persistent' adds
	// Type is resolved by field type
	// 'P' suffix is resolved by 'short' tag
	// Default value is empty is missing, so default type value
	// Description is resolved by 'description' tag
	child.Flags().StringP("output", "o", "", "output file name")
	// Generated by flag adds 'required'
	if err := child.MarkFlagRequired("output"); err != nil {
		fatal(err)
	}
	child.Flags().Int("num", 42, "some number for command")

	// DO NOT generate add the flag is it marked as 'inherited' and should be taken from root command
	//child.PersistentFlags().BoolP("is-ok", "i", true, "some persistent flag")
	//child.InheritedFlags().BoolP("is-ok", "i", true, "some persistent flag")

	// Generated because @CobraRun (always use RunE, because we need to track parse and other errors)
	child.RunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.ChildCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.Run(cmd, args)
	}

	// Generated because @CobraPostRun (always use PostRunE, because we need to track parse and other errors)
	child.PostRunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.ChildCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.PostRun(cmd, args)
	}

	// Generated because @CobraPersistPostRun (always use PersistentPostRunE, because we need to track parse and other errors)
	child.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		// Get the command type
		executor := commands.ChildCommand{}
		// Parse flags
		if err := parseFlags(cmd, &executor); err != nil {
			return err
		}
		// Generated as the end of function. The method name is taken from annotation
		return executor.PersistPostRun(cmd, args)
	}

	// Generated as usage has 'cli get'? that marks root command as a receiver
	root.AddCommand(child)
}
*/