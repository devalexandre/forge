package architecture

import (
	"fmt"
	"os/exec"

	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func createMicroserviceArchitecture(projectName string) error {
	cmd := exec.Command("git", "clone", "git@github.com:devalexandre/golang-ddd-template.git", projectName)
	return cmd.Run()

}

func removeExample(path string) {
	log.Printf("removing : %v", path)
	cmd := exec.Command("rm", "-rf", path)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing example: %v", err)

	}
}

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

		}

		if !examples {
			count := 0
			folders := []string{"internal/domain/user", "internal/infra/pgx"}
			log.Info("Removing examples...")
			for {

				//check if internal/infra/mysql and internal/domain/user exists
				for _, folder := range folders {
					if _, err := os.Stat(fmt.Sprintf("%s/%s", name, folder)); err == nil {
						log.Printf("Removing %s...", folder)
						removeExample(name + "/" + folder)
						count++
					}

					if count == len(folders) {
						os.Exit(0)
					}

				}

			}

		}

		fmt.Printf("Microservice architecture created successfully! (%s)", name)
	},
}
