package create

import (
	"fmt"
	"github.com/devalexandre/forge/src/log"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	CreateDomain.Flags().StringP("name", "n", "", "Name of the domain")
	CreateDomain.Flags().StringP("type", "t", "domain", "Type to create: domain, contract, repository, service, factory, mock")
	return CreateDomain
}

var CreateDomain = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a domain, contract, repository, service, factory or mock",
	Long:  "Create a domain , contract, repository, service, factory or mock",
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
			log.Printf("Type not found")
		}

	},
}

func updateName(path string, name string) {
	//update package in all files .go

	src, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	//update package name
	src = []byte(strings.ReplaceAll(string(src), "user", name))
	//update struct name
	src = []byte(strings.ReplaceAll(string(src), "User", strings.Title(name)))

	//update file

	if ioutil.WriteFile(path, src, 0644) != nil {
		log.Errorf(err.Error())
		return
	}

}

func createDomain(name string) {
	// check if internal folder exists
	var paths []string
	if _, err := os.Stat("./internal"); err != nil {
		log.Printf("You need to be in the root of the project")
		return
	}

	cmd := exec.Command("mkdir", "-p", "./internal/domain/"+name)
	cmd.Run()

	//create contract
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/contracts.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/contracts.go")
	cmd.Run()
	paths = append(paths, "./internal/domain/"+name+"/contracts.go")
	//create repository
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/repository.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/repository.go")
	cmd.Run()
	paths = append(paths, "./internal/domain/"+name+"/repository.go")

	//create service
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/service.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/service.go")
	cmd.Run()
	paths = append(paths, "./internal/domain/"+name+"/service.go")

	//create factory
	cmd = exec.Command("curl", "-o", "./internal/domain/"+name+"/factory.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/factory.go")
	cmd.Run()
	paths = append(paths, "./internal/domain/"+name+"/factory.go")

	//mock directory
	cmd = exec.Command("mkdir", "-p", "./internal/domain/"+name+"/mock")
	cmd.Run()

	for _, path := range paths {
		//update package in all files .go
		updateName(path, name)
	}

	//update package in all files .go

	fmt.Printf("Domain %s created successfully", name)

}

func createContract(name string) {

	cmd := exec.Command("curl", "-o", "./contracts.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/contracts.go")
	cmd.Run()

	//update package in all files .go
	updateName("./contracts.go", name)

}

func createRepository(name string) {

	cmd := exec.Command("curl", "-o", "./repository.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/repository.go")
	cmd.Run()

	updateName("./repository.go", name)

}

func createService(name string) {

	cmd := exec.Command("curl", "-o", "./service.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/service.go")
	cmd.Run()

	updateName("./service.go", name)
}

func createFactory(name string) {

	cmd := exec.Command("curl", "-o", "./factory.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/factory.go")
	cmd.Run()

	updateName("./factory.go", name)
}

func createMock(name string) {

	cmd := exec.Command("curl", "-o", "./mock/mock.go", "https://raw.githubusercontent.com/devalexandre/golang-ddd-template/main/internal/domain/user/mock/mock.go")
	cmd.Run()

	updateName("./mock/mock.go", name)
}
