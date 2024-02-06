package create

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

func InitAdpters() *cobra.Command {
	CreateAdpters.Flags().StringP("name", "n", "", "Name of the domain")
	CreateAdpters.Flags().StringP("type", "t", "domain", "Type to create: domain, contract, repository, service, factory, mock")
	return CreateAdpters
}

var CreateAdpters = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a domain, contract, repository, service, factory or mock",
	Long:  "Create a domain , contract, repository, service, factory or mock",
	Run: func(cmd *cobra.Command, args []string) {
		tpe, _ := cmd.Flags().GetString("type")

		switch tpe {
		case "pgx":
			createPGX()

		default:
			log.Println("Type not found")
		}

	},
}

func createPGX() {

	cmd := exec.Command("mkdir", "-p", "./internal/infra/pgx")
	cmd.Run()

	cmd = exec.Command("curl", "-o", "./internal/infra/pgx/contracts.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/pgx/contracts.go")
	cmd.Run()

	cmd = exec.Command("curl", "-o", "./internal/infra/pgx/connection.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/pgx/connection.go")
	cmd.Run()

	cmd = exec.Command("curl", "-o", "./internal/infra/pgx/connection_test.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/pgx/connection_test.go")
	cmd.Run()

	//create mocks dir
	cmd = exec.Command("mkdir", "-p", "./internal/infra/pgx/mocks")
	cmd.Run()

	cmd = exec.Command("curl", "-o", "./internal/infra/pgx/mocks/connection.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/pgx/mocks/connection.go")
	cmd.Run()

	cmd = exec.Command("curl", "-o", "./internal/infra/pgx/mocks/db.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/infra/pgx/mocks/db.go")
	cmd.Run()

	color.Green("pgx adapter created !!! ")

}
