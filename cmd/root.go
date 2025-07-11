package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "lightvault",
	Short:"lightvault - Secure local secreats manager",
	Long:`LightVault is a CLI tool to securely store and manage secrets locally.`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("LightVault cli is ready!")
	},
	
}

func Execute (){
	err := rootCmd.Execute()
	if err != nil{
		os.Exit(1)
	}


}