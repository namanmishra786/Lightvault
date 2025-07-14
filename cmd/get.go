package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// getCmd defines the CLI command: lightvault get <key>
var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get the value of a secret key",
	Long:  `This command retrieves the value of a secret key stored in your LightVault store.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		// Find home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}

		// Build LightVault file path
		lightvaultPath := filepath.Join(homeDir, ".lightvault")

		// Read file
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

		// Look up the key
		value, exists := secrets[key]
		if !exists {
			fmt.Printf("Secret [%s] not found.\n", key)
			return
		}

		fmt.Printf("Secret [%s]: %s\n", key, value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
