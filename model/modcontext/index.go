/*
Copyright 2015 anzi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/
package modcontext

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Dburi    string `json:"dburi"`
	Maxopen  int    `json:"maxopen"`
	Maxidle  int    `json:"maxidle"`
	Pagesize int    `json:"pagesize"`
}

var configuration = Configuration{}

func Coninit() error {
	filebytes, _ := ioutil.ReadFile(os.Args[1])
	err := json.Unmarshal(filebytes, &configuration)
	return err
}
func Default() *Configuration {
	return &configuration
}
