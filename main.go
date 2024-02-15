package main

import (
	"github.com/devalexandre/forge/src/architecture"
	"github.com/devalexandre/forge/src/create"
	"github.com/devalexandre/forge/src/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Init() {
	log.LogLevel = "all"
	log.PrettyPrint = false
}

func displayArt() string {
	art := `
FFFFFFFFFFFFFFFFFFFFFF                                                                          
F::::::::::::::::::::F                                                                          
F::::::::::::::::::::F                                                                          
FF::::::FFFFFFFFF::::F                                                                          
  F:::::F       FFFFFFooooooooooo   rrrrr   rrrrrrrrr      ggggggggg   ggggg    eeeeeeeeeeee    
  F:::::F           oo:::::::::::oo r::::rrr:::::::::r    g:::::::::ggg::::g  ee::::::::::::ee  
  F::::::FFFFFFFFFFo:::::::::::::::or:::::::::::::::::r  g:::::::::::::::::g e::::::eeeee:::::ee
  F:::::::::::::::Fo:::::ooooo:::::orr::::::rrrrr::::::rg::::::ggggg::::::gge::::::e     e:::::e
  F:::::::::::::::Fo::::o     o::::o r:::::r     r:::::rg:::::g     g:::::g e:::::::eeeee::::::e
  F::::::FFFFFFFFFFo::::o     o::::o r:::::r     rrrrrrrg:::::g     g:::::g e:::::::::::::::::e 
  F:::::F          o::::o     o::::o r:::::r            g:::::g     g:::::g e::::::eeeeeeeeeee  
  F:::::F          o::::o     o::::o r:::::r            g::::::g    g:::::g e:::::::e           
FF:::::::FF        o:::::ooooo:::::o r:::::r            g:::::::ggggg:::::g e::::::::e          
F::::::::FF        o:::::::::::::::o r:::::r             g::::::::::::::::g  e::::::::eeeeeeee  
F::::::::FF         oo:::::::::::oo  r:::::r              gg::::::::::::::g   ee:::::::::::::e  
FFFFFFFFFFF           ooooooooooo    rrrrrrr                gggggggg::::::g     eeeeeeeeeeeeee  
                                                                    g:::::g                     
                                                        gggggg      g:::::g                     
                                                        g:::::gg   gg:::::g                     
                                                         g::::::ggg:::::::g                     
                                                          gg:::::::::::::g                      
                                                            ggg::::::ggg                        
                                                               gggggg                           
`
	art += "\n"
	art += "Microservice CLI\n"
	art += "Version: 0.0.4\n"
	return art
}

func main() {
	Init()

	color.Cyan(displayArt())
	rootCmd := cobra.Command{
		Use: "forge",
		//Short: displayArt(),
	}
	rootCmd.AddCommand(architecture.Init())
	rootCmd.AddCommand(create.Init())
	rootCmd.AddCommand(create.InitAdpters())

	rootCmd.Execute()
}
