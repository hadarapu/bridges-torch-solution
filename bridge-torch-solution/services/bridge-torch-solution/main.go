package main

import (
	"fmt"
	"git.sqad.io/bridge-torch-solution/services/bridge-torch-solution/handler"
	. "git.sqad.io/bridge-torch-solution/services/common"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFile   string
	numOfWorkers int
)

func main() {
	//read the arguments
	//first - yaml filename
	//second - number of workers

	//yaml file has an entry called problem
	//it is a map with key as bridge id and value is an array of person ids..
	//you can add any number of bridges and any number of persons for each bridge

	rootCmd.Flags().StringVarP(&configFile, "configFile", "i", "/services/resources/config.yaml", "path to yaml file")
	rootCmd.Flags().IntVarP(&numOfWorkers, "numOfWorkers", "n", 5, "number of worker go routines to run")
	Execute()
}

// Execute  executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command {
	Use: "bridge-torch-solution",
	Short: "solution for bridges torch problem",
	Long: "This is the solution for people crossing bridges with a torch",
	RunE: func(cmd *cobra.Command, args []string) error {
		if configFile == "" || numOfWorkers == 0 {
			panic("program didnt take default arguments.. please check the default values")
		}

		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(path+configFile)

		conf, err := LoadYamlFile(path+configFile)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		workerInput := make(chan *InputObject)
		workerOutput := make(chan *OutputObject)

		//instead of calculating time one after another, lets create a routine pool so that things are done in parallel
		//create a go routine pool to avoid creating many routines
		for i:=0; i<numOfWorkers; i++ {
			go handler.CalculateQuickestTime(workerInput, workerOutput, conf)
		}

		//push messages to input queue
		go func(workerInput chan *InputObject) {
			for bridgeId, personsList := range conf.Problem {
				workerInput <- &InputObject{BridgeId: bridgeId, PersonIdsList: personsList, Cfg:conf}
			}
			close(workerInput)
		}(workerInput)

		//wait on the output queue
		var totalTime float64
		for bridgeId, _ := range conf.Problem {
			fmt.Sprintf("%d", bridgeId)
			output := <- workerOutput
			totalTime += output.QuickestTime
		}
		//close to make sure that no one is sending messages anymore
		close(workerOutput)

		fmt.Println(fmt.Sprintf("total optimal time it took for all people to cross all bridges: %v", totalTime ))

		return nil
	},
}