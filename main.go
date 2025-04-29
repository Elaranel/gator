package main

import (
	"fmt"

	"github.com/elaranel/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("elaranel")
	if err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

}
