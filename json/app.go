package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Your JSON data
	jsonData := []byte(`{
	"InstanceStatuses": [
	 {
	  "AttachedEbsStatus": null,
	  "AvailabilityZone": "ap-northeast-1c",
	  "Events": null,
	  "InstanceId": "i-0694c6f5f0648ed0e",
	  "InstanceState": {
	   "Code": 80,
	   "Name": "stopped"
	  },
	  "InstanceStatus": {
	   "Details": null,
	   "Status": "not-applicable"
	  },
	  "Operator": {
	   "Managed": false,
	   "Principal": null
	  },
	  "OutpostArn": null,
	  "SystemStatus": {
	   "Details": null,
	   "Status": "not-applicable"
	  }
	 }
	],
	"NextToken": null,
	"ResultMetadata": {}
   }`)

	var result map[string]any
	// converting []bytes to Golang DataStructure
	err := json.Unmarshal(jsonData, &result)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	//fmt.Println(result)

	instances, _ := result["InstanceStatuses"].([]any)
	fmt.Println(instances)

	for _, inst := range instances {
		instance, ok := inst.(map[string]any)

		if !ok {
			continue
		}

		fmt.Println(instance["AvailabilityZone"])
	}

}
