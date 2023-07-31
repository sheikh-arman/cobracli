/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	vp := viper.New()

	vp.SetConfigName("test")
	vp.SetConfigType("json")
	vp.AddConfigPath("./viper")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(vp.GetString("foo"))

	vp.Set("name", "arman")
	vp.WriteConfig()

	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("file changed: %s\n", in.Name)
	})
	//cmd.Execute()
	vp.WatchConfig()
	for {
	}
}
