package architecture

import (
	"fmt"

	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	CreateStruct.Flags().StringP("name", "n", "", "Name of the microservice architecture")
	//flag for example
	CreateStruct.Flags().BoolP("examples", "e", false, "example for struct")

	return CreateStruct
}

var CreateStruct = &cobra.Command{
	Use:   "init [flags]",
	Short: "Create a microservice architecture",
	Long:  "Create a microservice architecture",

	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		examples, _ := cmd.Flags().GetBool("examples")

		err := createMicroserviceArchitecture(name)
		if err != nil {
			log.Fatalf("Error creating microservice architecture: %v", err)
			os.Exit(1)
		}

		log.Printf("Example: %v", examples)

		if examples {
			folders := 1
			count := 0
			for {
				if _, err := os.Stat(fmt.Sprintf("%s/default.sh", name)); err == nil {
					log.Info("Removing default.sh...")
					removeExample(name + "/default.sh")
					count++
				}

				if count == folders {
					break
				}
		}

		if !examples {
			folders := 3
			count := 0
			log.Info("Removing examples...")
			for {

				//check if internal/infra/mysql and internal/domain/user exists

				if _, err := os.Stat(fmt.Sprintf("%s/internal/infra/mysql", name)); err == nil {
					log.Info("Removing internal/infra/mysql...")
					removeExample(name + "/internal/infra/mysql")
					count++
				}

				if _, err := os.Stat(fmt.Sprintf("%s/internal/domain/user", name)); err == nil {
					log.Info("Removing internal/domain/user...")
					removeExample(name + "/internal/domain/user")
					count++
				}

				if _, err := os.Stat(fmt.Sprintf("%s/default.sh", name)); err == nil {
					log.Info("Removing default.sh...")
					removeExample(name + "/default.sh")
					count++
				}

				if count == folders {
					break
				}

			}

		}

	},
}

func createMicroserviceArchitecture(projectName string) error {
	cmd := exec.Command("git", "clone", "git@github.com:devalexandre/golang-ddd-template.git", projectName)
	return cmd.Run()

}

func removeExample(path string) {
	cmd := exec.Command("rm", "-rf", path)
	cmd.Run()
}
