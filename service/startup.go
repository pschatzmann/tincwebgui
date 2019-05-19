package service

import (
	"log"
	"os"
	"os/exec"
)

// StartTinc - startup Tinc
func StartTinc() {

	log.Println("starting...")
	// check if tinc is already started
	out, err := exec.Command("tinc", "pid").CombinedOutput()
	log.Println("pid:", string(out))
	if err != nil {
		// we need to start it
		out, err = exec.Command("tinc", "start").CombinedOutput()
		log.Println("start:", string(out))
		if err != nil {
			name, ok := os.LookupEnv("NodeName")
			if !ok {
				// if the name is not defined we use the (docker) hostname
				name, ok = os.LookupEnv("HOSTNAME")
			}

			if ok {
				out, err = exec.Command("tinc", "init", name).CombinedOutput()
				log.Println("init:", string(out))
				if err != nil {
					out, err = exec.Command("tinc", "start").CombinedOutput()
					log.Println("start:", string(out))
					// we are ready to setup the variables
					setupTincVariables()
					out, err = exec.Command("tinc", "reload").CombinedOutput()
					log.Println("reload:", string(out))
				}
			}
		} else {
			log.Println("Could not start tinc because the name is not defined")
		}
	} else {
		setupTincVariables()
		out, err = exec.Command("tinc", "reload").CombinedOutput()
		log.Println("reload:", string(out))
	}
}

// copy env variables to tinc (replace NODENAME with NAME)
func setupTincVariables() {
	log.Println("setupTincVariables")
	for _, name := range ListParameterKeys() {
		key := name
		if key == "NodeName" {
			key = "Name"
		}
		value, ok := os.LookupEnv(name)
		if ok {
			log.Println("set", name, value)
			out, _ := exec.Command("tinc", "set", key, value).CombinedOutput()
			log.Println(" -> set", key, "->", string(out))
		}
	}
}
