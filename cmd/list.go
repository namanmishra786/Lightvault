package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd defines: lightvault list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secret keys stored in LightVault",
	Long:  `This command lists all the secret keys currently stored in your LightVault store.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Find home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}

		// Build LightVault file path
		lightvaultPath := filepath.Join(homeDir, ".lightvault")

		// Read secrets file
		file, err := os.ReadFile(lightvaultPath)
		if err != nil {
			fmt.Println("Error reading LightVault store:", err)
			os.Exit(1)
		}

		// Parse JSON
		var secrets map[string]string
		if err := json.Unmarshal(file, &secrets); err != nil {
			fmt.Println("Error parsing LightVault store:", err)
			os.Exit(1)
		}

		// List keys
		if len(secrets) == 0 {
			fmt.Println("No secrets found.")
			return
		}

		fmt.Println("Stored secrets:")
		for key := range secrets {
			fmt.Printf("- %s\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
