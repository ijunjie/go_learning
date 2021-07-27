/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"cluster-register/cmd"
	"fmt"
)

func main() {
	cmd.Execute()
	fmt.Printf("host: %s \n", cmd.KdeHost)
	fmt.Printf("port: %d \n", cmd.KdePort)
	fmt.Printf("username: %s \n", cmd.KdeUsername)
	fmt.Printf("password: %s \n", cmd.KdePassword)
	fmt.Printf("write-to-db: %t \n", cmd.WriteToDB)
	if cmd.WriteToDB {

		fmt.Printf("mysql host: %s \n", cmd.MySqlHost)
		fmt.Printf("mysql port: %d \n", cmd.MySqlPort)
		fmt.Printf("mysql username: %s \n", cmd.MySqlUsername)
		fmt.Printf("mysql password: %s \n", cmd.MySqlPassword)
	}
}
