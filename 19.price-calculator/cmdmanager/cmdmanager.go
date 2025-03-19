package cmdmanager

import (
	"encoding/json"
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadFileContents() ([]string, error) {
	fmt.Println("Please enter your prices.")

	var prices []string

	for {
		var v string
		fmt.Scan(&v)

		prices = append(prices, v)
		if v == "0" {
			break
		}
	}

	return prices, nil
}

func (cmd CMDManager) WriteJSON(data any) error {

	jsonData, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
