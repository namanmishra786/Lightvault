package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a secret key-value pair",
	Long:  `This command sets a new secret or updates an existing one in your LightVault store.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}

		lightvaultPath := filepath.Join(homeDir, ".lightvault")

		file, err := os.ReadFile(lightvaultPath)
		if err != nil {
			fmt.Println("Error reading LightVault store:", err)
			os.Exit(1)
		}

		var secrets map[string]string
		if err := json.Unmarshal(file, &secrets); err != nil {
			fmt.Println("Error parsing LightVault store:", err)
			os.Exit(1)
		}

		secrets[key] = value

		newData, err := json.MarshalIndent(secrets, "", "  ")
		if err != nil {
			fmt.Println("Error encoding secrets:", err)
			os.Exit(1)
		}

		if err := os.WriteFile(lightvaultPath, newData, 0600); err != nil {
			fmt.Println("Error writing LightVault store:", err)
			os.Exit(1)
		}

		fmt.Printf("Secret [%s] set successfully!\n", key)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
