package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func modifyParameters(parameter, env string) string {
	switch env {
	case "dev":
		return strings.Replace(parameter, "/local/", "/dev/", 1)
	case "qa":
		return strings.Replace(parameter, "/local/", "/asgard/", 1)
	case "stage":
		return strings.Replace(parameter, "/local/", "/stage/", 1)
	case "build":
		return strings.Replace(parameter, "/local/", "/local/", 1)
	default:
		fmt.Println("Do nothing")
	}
	return parameter
}

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	srcParameterName := os.Getenv("SRC_PARAMETER_NAME")
	destParameterName := os.Getenv("DEST_PARAMETER_NAME")
	region := os.Getenv("AWS_REGION")

	if srcParameterName == "" || destParameterName == "" || region == "" {
		fmt.Println("SRC_PARAMETER_NAME, DEST_PARAMETER_NAME, and AWS_REGION environment variables are required.")
		printHelp()
		os.Exit(1)
	}

	// Define the list of environments
	environments := []string{"dev", "asgard", "stage", "local"}

	// Loop through each environment and modify parameters
	for _, env := range environments {
		modifiedSrcParameterName := modifyParameters(srcParameterName, env)
		modifiedDestParameterName := modifyParameters(destParameterName, env)

		// Create an AWS session and SSM client
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
		ssmClient := ssm.New(sess)

		// Check if the destination parameter exists
		destParamInput := &ssm.GetParameterInput{
			Name: aws.String(modifiedDestParameterName),
		}
		destParamOutput, err := ssmClient.GetParameter(destParamInput)
		if err != nil {
			fmt.Printf("Error getting destination parameter %s: %v\n", modifiedDestParameterName, err)
			os.Exit(1)
		}

		// Get the source parameter value
		srcParamInput := &ssm.GetParameterInput{
			Name:           aws.String(modifiedSrcParameterName),
			WithDecryption: aws.Bool(true),
		}
		srcParamOutput, err := ssmClient.GetParameter(srcParamInput)
		if err != nil {
			fmt.Println("Error getting source parameter:", err)
			os.Exit(1)
		}

		// Print the values before overwriting
		fmt.Printf("\nBefore Overwrite:\nSource Parameter (%s): %s\nDestination Parameter (%s): %s\n",
			modifiedSrcParameterName, *srcParamOutput.Parameter.Value,
			modifiedDestParameterName, *destParamOutput.Parameter.Value)

		// Create or update the destination parameter
		destParamInputPut := &ssm.PutParameterInput{
			Name:      aws.String(modifiedDestParameterName),
			Value:     srcParamOutput.Parameter.Value,
			Type:      srcParamOutput.Parameter.Type,
			Overwrite: aws.Bool(true),
		}
		_, err = ssmClient.PutParameter(destParamInputPut)
		if err != nil {
			fmt.Println("Error copying parameter:", err)
			os.Exit(1)
		}

		// Print the values after overwriting
		fmt.Printf("\nAfter Overwrite:\nSource Parameter (%s): %s\nDestination Parameter (%s): %s\n",
			modifiedSrcParameterName, *srcParamOutput.Parameter.Value,
			modifiedDestParameterName, *srcParamOutput.Parameter.Value)

		fmt.Printf("\nParameter %s successfully copied to %s\n", modifiedSrcParameterName, modifiedDestParameterName)
	}
}

func printHelp() {
	fmt.Println("Usage: AWS_REGION=${PUT-YOUR-REGION} SRC_PARAMETER_NAME=${PUT-YOUR-SOURCE-SSM-NAME} DEST_PARAMETER_NAME=${PUT-YOUR-DESTINATION-SSM-NAME} copy-ssm-parameters")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}
