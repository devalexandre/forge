package create

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"

	"os/exec"
)

func InitAdpters() *cobra.Command {
	CreateAdpters.Flags().StringP("type", "t", "", "mysql or postgres or sqlite")
	return CreateAdpters
}

var CreateAdpters = &cobra.Command{
	Use:   "adapter [flags]",
	Short: "Create a adapter",
	Long:  "Create a adapter for mysql, postgres or sqlite",
	Run: func(cmd *cobra.Command, args []string) {
		tpe, _ := cmd.Flags().GetString("type")

		color.Red(fmt.Sprintf("Creating %s", tpe))
		switch tpe {
		case "mysql":
			createMySQL()
			break
		case "postgres":
			createPGX()
			break
		case "sqlite":
			createSQLite()
			break
		default:
			chooseAdapter()
			break
		}

	},
}

func chooseAdapter() {
	color.Magenta("Choose the adapter to create: ")
	color.White("1 - mysql")
	color.White("2 - postgres")
	color.White("3 - sqlite")

	var adapter int
	// get the user's choice
	_, err := fmt.Scan(&adapter)

	if err != nil {
		log.Fatalf("Error reading the user's choice: %v", err)

	}

	switch adapter {
	case 1:
		createMySQL()
		break
	case 2:
		createPGX()
		break
	case 3:
		createSQLite()
		break
	default:
		log.Println("Adapter not found")
	}

}

func createPGX() {
	createAdapter("pgx")
}

func createMySQL() {
	createAdapter("mysql")
}

func createSQLite() {
	createAdapter("sqlite")
}

func createAdapter(adapter string) {

	cmd := exec.Command("mkdir", "-p", fmt.Sprintf("./internal/infra/database/%s", adapter))
	cmd.Run()

	cmd = exec.Command("curl", "-o", fmt.Sprintf("./internal/infra/database/%s/contracts.go", adapter), fmt.Sprintf("https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/database/%s/contracts.go", adapter))
	cmd.Run()

	cmd = exec.Command("curl", "-o", fmt.Sprintf("./internal/infra/database/%s/connection.go", adapter), fmt.Sprintf("https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/database/%s/connection.go", adapter))
	cmd.Run()

	cmd = exec.Command("curl", "-o", fmt.Sprintf("./internal/infra/database/%s/connection_test.go", adapter), fmt.Sprintf("https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/database/%s/connection_test.go", adapter))
	cmd.Run()

	//verify if the mocks dir exists
	if _, err := exec.Command("ls", fmt.Sprintf("./internal/infra/database/%s/mocks", adapter)).Output(); err != nil {
		cmd = exec.Command("mkdir", "-p", fmt.Sprintf("./internal/infra/database/%s/mocks", adapter))
		cmd.Run()
	}

	// verify if the mocks files exists
	if _, err := exec.Command("ls", fmt.Sprintf("./internal/infra/database/%s/mocks/connection.go", adapter)).Output(); err != nil {
		cmd = exec.Command("curl", "-o", fmt.Sprintf("./internal/infra/database/%s/mocks/connection.go", adapter), fmt.Sprintf("https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/database/%s/mocks/connection.go", adapter))
		cmd.Run()
	}

	if _, err := exec.Command("ls", fmt.Sprintf("./internal/infra/database/%s/mocks/db.go", adapter)).Output(); err != nil {
		cmd = exec.Command("curl", "-o", fmt.Sprintf("./internal/infra/database/%s/mocks/db.go", adapter), fmt.Sprintf("https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/database/%s/mocks/db.go", adapter))
		cmd.Run()
	}
	color.Green(fmt.Sprintf("%s adapter created !!! ", adapter))

}
