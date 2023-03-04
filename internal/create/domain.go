package create

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Init() *cobra.Command {
	CreateDomain.Flags().StringP("name", "n", "", "Name of the domain")
	CreateDomain.Flags().StringP("type", "t", "domain", "Type to create: domain, contract, repository, service, factory, mock")
	return CreateDomain
}

var CreateDomain = &cobra.Command{
	Use:   "domain [flags]",
	Short: "Create a domain",
	Long:  "Create a domain",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		tpe, _ := cmd.Flags().GetString("type")

		switch tpe {
		case "domain":
			createDomain(name)
		case "contract":
			createContract(name)
		case "repository":
			createRepository(name)
		case "service":
			createService(name)
		case "factory":
			createFactory(name)
		case "mock":
			createMock(name)
		default:
			log.Println("Type not found")
		}

	},
}

func createDomain(name string) {
	// check if internal folder exists
	if _, err := os.Stat("./internal"); err != nil {
		log.Println("You need to be in the root of the project")
		return
	}

	cmd := exec.Command("mkdir", "-p", "./internal/domain/"+name)
	cmd.Run()
	//create contract
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/contracts.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/contracts.go")
	cmd.Run()

	//create repository
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/repository.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/repository.go")
	cmd.Run()

	//create service
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/service.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/service.go")
	cmd.Run()

	//create factory
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/factory.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/factory.go")
	cmd.Run()

	//mock directory
	cmd = exec.Command("mkdir", "-p", "./internal/domain/"+name+"/mock")
	cmd.Run()

	//update package in all files .go
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./internal/domain/"+name+"/*.go")
	cmd.Run()

	//update User to name in all files .go
	cmd = exec.Command("sed", "-i", "s/User/"+strings.Title(name)+"/g", "./internal/domain/"+name+"/*.go")
	cmd.Run()
}

func createContract(name string) {

	cmd := exec.Command("curl", "-o", "./contracts.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/contracts.go")
	cmd.Run()

	//update package name user to name
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./"+name+"/*.go")
	cmd.Run()

	//update User to name using case
	name = cases.Title(language.English, cases.Compact).String(name)
	cmd = exec.Command("sed", "-i", "s/User/"+name+"/g", "./"+name+"/*.go")
	cmd.Run()
}

func createRepository(name string) {

	cmd := exec.Command("curl", "-o", "./repository.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/repository.go")
	cmd.Run()

	//update package in all files .go
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./"+name+"/*.go")
	cmd.Run()

	//update User to name using case
	name = cases.Title(language.English, cases.Compact).String(name)

	cmd = exec.Command("sed", "-i", "s/User/"+name+"/g", "./"+name+"/*.go")
}

func createService(name string) {

	cmd := exec.Command("curl", "-o", "./service.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/service.go")
	cmd.Run()

	//update package in all files .go
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./"+name+"/*.go")
	cmd.Run()

	//update User to name using case
	name = cases.Title(language.English, cases.Compact).String(name)

	cmd = exec.Command("sed", "-i", "s/User/"+name+"/g", "./"+name+"/*.go")
}

func createFactory(name string) {

	cmd := exec.Command("curl", "-o", "./factory.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/factory.go")
	cmd.Run()

	//update package in all files .go
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./"+name+"/*.go")
	cmd.Run()

	//update User to name in all files .go
	cmd = exec.Command("sed", "-i", "s/User/"+strings.Title(name)+"/g", "./internal/domain/"+name+"/*.go")
	cmd.Run()
}

func createMock(name string) {

	cmd := exec.Command("curl", "-o", "./mock/mock.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/mock/mock.go")
	cmd.Run()

	//update package in all files .go
	cmd = exec.Command("sed", "-i", "s/user/"+name+"/g", "./mock/*.go")
	cmd.Run()

	//update User to name in all files .go
	cmd = exec.Command("sed", "-i", "s/User/"+strings.Title(name)+"/g", "./mock/*.go")
	cmd.Run()
}
