package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Convert(inputFile string, outputFormat string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf(("file not found: %s"), inputFile)
	}

	ext := strings.ToLower((filepath.Ext(inputFile)))

	var inputFormat string

	if ext == ".json"{
		inputFormat = "json"
	} else if ext ==".yaml" || ext == ".yml"{
		inputFormat = "yaml"
	} else {
		return fmt.Errorf(("invalid file format %s"), ext)
	}
	if inputFormat == outputFormat {
		return fmt.Errorf("requested format should be different than the input format %s", inputFormat)	
	}
	var resultMap map[string]interface{}
	
	if inputFormat == "json"{
		err = json.Unmarshal(data, &resultMap)
	
	} else {
		err = yaml.Unmarshal(data, &resultMap)
		
	}
	if err != nil {
			return fmt.Errorf("failed to parse data: %s", err)
	}

	var outputData []byte

	if outputFormat == "json"{
		outputData, err = json.MarshalIndent(resultMap, "", "  ")
	} else {
		outputData, err = yaml.Marshal(resultMap)
	}
	if err != nil {
	
		return fmt.Errorf("failed to marshal data: %s", err)
	}

	// 	Print the output to terminal
	fmt.Println(string(outputData))
	// Build the output file path
	dir := filepath.Dir(inputFile)
	baseName := strings.TrimSuffix(filepath.Base(inputFile), ext)
	outputPath := filepath.Join(dir, baseName + "." + outputFormat)
	// Write the file
	err = os.WriteFile(outputPath, outputData, 0644)
	// Check for errors
	if err != nil {
		return fmt.Errorf("failed to write file: %s", err)
	}
	// Return nil
	return nil
}