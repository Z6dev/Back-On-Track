package todoutils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Z6dev/Back-On-Track/structs"
)

// LoadTodos loads a JSON file into a slice of TODO.
// If the file does not exist, it creates one and returns an empty slice.
func LoadTodos(filename string) ([]structs.TODO, error) {
	var todos []structs.TODO

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Create empty file with []
		err = os.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		return todos, nil
	}

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Decode JSON
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return todos, nil
}

// SaveTodos saves a slice of TODO to a JSON file.
func SaveTodos(filename string, todos []structs.TODO) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
