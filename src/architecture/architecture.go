package architecture

import (
	"fmt"
	"github.com/devalexandre/go-forge/src/config"
	"github.com/devalexandre/go-forge/src/log"
	"github.com/fatih/color"

	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func createMicroserviceArchitecture(projectName string, template string) error {
	log.Debugf("Clone template: %v", template)
	cmd := exec.Command("git", "clone", template, projectName)
	return cmd.Run()

}

func removeExample(path string) {
	log.Debugf("removing : %v", path)
	log.Debugf("rm -rf %v", path)
	cmd := exec.Command("rm", "-rf", path)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing example: %v", err)

	}
}

func Init() *cobra.Command {
	CreateStruct.Flags().StringP("name", "n", "", "Name of the microservice architecture")
	//flag for example
	CreateStruct.Flags().BoolP("examples", "e", true, "example for struct")

	CreateStruct.Flags().StringP("template", "t", config.TemplateTDD, "template to use: tdd, ddd, hexagonal")

	CreateStruct.MarkFlagRequired("name")

	return CreateStruct
}

var CreateStruct = &cobra.Command{
	Use:   "init [flags]",
	Short: "Create a microservice architecture",
	Long:  "Create a microservice architecture",

	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		examples, _ := cmd.Flags().GetBool("examples")
		template, _ := cmd.Flags().GetString("template")

		err := createMicroserviceArchitecture(name, template)
		if err != nil {
			log.Fatalf("Error creating microservice architecture: %v", err)
			os.Exit(1)
		}

		if examples {
			// question: what database do you want to use?
			color.Magenta("What database do you want to use? ")
			color.White("1 - mysql")
			color.White("2 - postgres")
			color.White("3 - sqlite")

			var db int
			var folders []string
			// get the user's choice
			_, err := fmt.Scan(&db)

			if err != nil {
				log.Fatalf("Error reading the user's choice: %v", err)

			}

			color.Cyan("Creating infrastructure for the database...")
			switch db {
			case 1:
				folders = []string{"internal/infra/database/pgx", "internal/infra/database/sqlite"}
				break
			case 2:
				folders = []string{"internal/infra/database/mysql", "internal/infra/database/sqlite"}
				break
			case 3:
				folders = []string{"internal/infra/database/mysql", "internal/infra/database/pgx"}
				break
			default:
				log.Fatalf("Invalid option: %v", db)
				os.Exit(1)

			}

			// remove the folders that the user didn't choose
			count := 0
			for {

				for _, folder := range folders {
					if _, err := os.Stat(fmt.Sprintf("%s/%s", name, folder)); err == nil {
						removeExample(name + "/" + folder)
						count++
					} else {
						log.Debugf("Folder %s not found", folder)
						count++
					}

					if count == len(folders) {
						// go out of the loop
						color.Green("Microservice architecture created successfully! %s", name)
						os.Exit(0)
					}
				}

			}
		}

		if !examples {
			count := 0
			folders := []string{"internal/domain/user", "internal/infra/database"}
			log.Debugf("Removing examples...")
			for {

				//check if internal/infra/mysql and internal/domain/user exists
				for _, folder := range folders {
					if _, err := os.Stat(fmt.Sprintf("%s/%s", name, folder)); err == nil {
						log.Debugf("Removing %s...", folder)
						removeExample(name + "/" + folder)
						count++
					} else {
						log.Debugf("Folder %s not found", folder)
						count++
					}

					if count == len(folders) {
						color.Green("Microservice architecture created successfully! %s", name)
						os.Exit(0)
					}

				}

			}

		}

	},
}
