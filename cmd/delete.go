package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// deleteCmd deletes a secret by key
var deleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a secret key",
	Long:  `This command deletes a secret key-value pair from your LightVault store.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		// Find home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}

		// Build file path
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

		// Check if key exists
		if _, exists := secrets[key]; !exists {
			fmt.Printf("Secret [%s] not found.\n", key)
			return
		}

		// Delete key
		delete(secrets, key)

		// Save file back
		newData, err := json.MarshalIndent(secrets, "", "  ")
		if err != nil {
			fmt.Println("Error encoding secrets:", err)
			os.Exit(1)
		}

		if err := os.WriteFile(lightvaultPath, newData, 0600); err != nil {
			fmt.Println("Error writing LightVault store:", err)
			os.Exit(1)
		}

		fmt.Printf("Secret [%s] deleted successfully!\n", key)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
