package service

import (
	"log"
	"os"
	"os/exec"
)

// StartTinc - startup Tinc
func StartTinc() {
	setupTincVariables()

	log.Println("starting...")
	// check if tinc is already started
	out, err := exec.Command("tinc", "pid").CombinedOutput()
	log.Println("pid:", string(out))
	if err != nil {
		// we need to start it
		out, err = exec.Command("tinc", "start").CombinedOutput()
		log.Println("start:", string(out))
		if err != nil {
			name, ok := os.LookupEnv("name")
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
				}
			}
		} else {
			log.Println("Could not start tinc because the name is not defined")
		}
	}
}

// copy env variables to tinc
func setupTincVariables() {
	log.Println("setupTincVariables")
	for _, name := range ListParameterKeys() {
		value, ok := os.LookupEnv(name)
		if ok {
			log.Println("set", name, value)
			out, _ := exec.Command("tinc", "set", name, value).CombinedOutput()
			log.Println(" -> set", name, "->", string(out))
		}
	}
}
