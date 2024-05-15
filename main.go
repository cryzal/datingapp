package main

import (
	"datingapp/app"
	"datingapp/shared/driver"
	"flag"
	"fmt"
)

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"user": app.NewUser(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
