package architecture

import (
	"fmt"
	"github.com/devalexandre/forge/src/config"
	"github.com/devalexandre/forge/src/log"
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
	CreateStruct.Flags().BoolP("initial", "i", true, "initial  struct")

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
		initial, _ := cmd.Flags().GetBool("initial")
		template, _ := cmd.Flags().GetString("template")

		err := createMicroserviceArchitecture(name, template)
		if err != nil {
			log.Fatalf("Error creating microservice architecture: %v", err)
			os.Exit(1)
		}

		var databaseAdapter, userExample bool
		var hasDatabaseAdapter, hasUserExample string
		var folders []string

		if initial {
			color.Magenta("Do you want to create a user example? (y/n) ")
			_, err := fmt.Scan(&hasUserExample)

			if err != nil {
				log.Fatalf("Error reading the user's choice: %v", err)
			}

			if hasUserExample == "y" {
				userExample = true
			}
		}

		if initial {
			color.Magenta("Do you want to create a database adapter? (y/n) ")
			_, err := fmt.Scan(&hasDatabaseAdapter)

			if err != nil {
				log.Fatalf("Error reading the user's choice: %v", err)
			}

			if hasDatabaseAdapter == "y" {
				databaseAdapter = true
			}
		}

		if databaseAdapter {
			// question: what database do you want to use?
			color.Magenta("What database do you want to use? ")
			color.White("1 - mysql")
			color.White("2 - postgres")
			color.White("3 - sqlite")

			var db int

			// get the user's choice
			_, err := fmt.Scan(&db)

			if err != nil {
				log.Fatalf("Error reading the user's choice: %v", err)

			}

			color.Cyan("Creating infrastructure for the database...")
			switch db {
			case 1:
				folders = append(folders, "internal/infra/database/pgx", "internal/infra/database/sqlite")
				break
			case 2:
				folders = append(folders, "internal/infra/database/mysql", "internal/infra/database/sqlite")
				break
			case 3:
				folders = append(folders, "internal/infra/database/mysql", "internal/infra/database/pgx")
				break
			default:
				log.Fatalf("Invalid option: %v", db)
				os.Exit(1)

			}

		}

		if !userExample {
			folders = append(folders, "internal/domain/user")
		}

		if !databaseAdapter {
			folders = append(folders, "internal/infra/database")
		}

		if len(folders) > 0 {
			count := 0

			log.Debugf("Removing initial...")
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
