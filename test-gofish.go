// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"fmt"
	"os"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	config := gofish.ClientConfig{
		Endpoint: "https://xxx.xxx.xxx.xxx",
		Username: "username",
		Password: "password",
		Insecure: true,
	}
	c, err := gofish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Retrieve the service root
	service := c.Service

	// Query the computer systems
	ss, err := service.Systems()
	if err != nil {
		panic(err)
	}

	for _, system := range ss {
		secureBoot, err := system.SecureBoot()
		if err != nil {
			fmt.Errorf("error getting secureboot")
			os.Exit(1)
		}
		fmt.Printf("Original Secureboot: %#v\n\n", secureBoot.SecureBootEnable)
		secureBoot.SecureBootEnable = true
		err = secureBoot.Update()

		if err != nil {
			fmt.Errorf("error getting secureboot")
			os.Exit(1)
		}
		secureBoot, err = system.SecureBoot()
		if err != nil {
			fmt.Errorf("error refresing secureboot")
			os.Exit(1)
		}
		fmt.Printf("secure boot refreshed: %#v\n\n", secureBoot.SecureBootEnable)

		err = system.Reset(redfish.ForceRestartResetType)
		if err != nil {
			panic(err)
		}
	}
}
