package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Define the `init` command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize LightVault local secret store",
	Long:  `This command sets up a local file to store your secrets securely.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 🟢 1️⃣ FIXED: Function name is UserHomeDir, not UseHomeDir
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}

		// 🟢 2️⃣ FIXED: filepath.Join, not filepath.join
		lightvaultPath := filepath.Join(homeDir, ".lightvault")

		// 🟢 3️⃣ FIXED: os.Stat checks if file exists, not os.Create
		if _, err := os.Stat(lightvaultPath); err == nil {
			fmt.Println("LightVault store already initialized at:", lightvaultPath)
			return
		}

		// 🟢 4️⃣ FIXED: Create the file
		file, err := os.Create(lightvaultPath)
		if err != nil {
			fmt.Println("Error creating LightVault store:", err)
			os.Exit(1)
		}
		defer file.Close() // 🟢 FIXED: Method is Close(), not close()

		// 🟢 5️⃣ FIXED: WriteString with uppercase W
		file.WriteString("{}")

		// 🟢 6️⃣ FIXED: fmt.Println not fmt.println
		fmt.Println("LightVault store initialized at:", lightvaultPath)
	},
}

// Connect the command to root
func init() {
	rootCmd.AddCommand(initCmd)
}
