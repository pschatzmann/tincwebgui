package service

/**
*  Command hander for tinc commands. We determine the commands which need to be executed and
*  execute them.
**/
import (
	"errors"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/pschatzmann/tincwebgui/auth"
)

// Command - struct with command
type Command struct {
	commands      []string
	parameterName string
	valueName     string
	formatter     func([]byte) ([]byte, Mime, error)
}

// NewCommand - Constructor
func NewCommand(commands ...string) *Command {
	return &Command{commands, "", "", CommandFormatter}
}

// NewCommand2 - Constructor with 2 arguments: function, commands
func NewCommand2(formatter func([]byte) ([]byte, Mime, error), commands ...string) *Command {
	var usedFormatter = formatter
	if usedFormatter == nil {
		usedFormatter = CommandFormatter
	}
	return &Command{commands, "", "", formatter}
}

// NewCommand3 - Constructor with 3 arguments. The 3d parameter is determined from the http request with the help of the parameterName
func NewCommand3(formatter func([]byte) ([]byte, Mime, error), parameterName string, commands ...string) *Command {
	var usedFormatter = formatter
	if usedFormatter == nil {
		usedFormatter = CommandFormatter
	}
	return &Command{commands, parameterName, "", formatter}
}

// NewCommand4 - Constructor with 3 arguments. The 3d and 4th parameters are determined from the http request with the help of the parameterName
func NewCommand4(formatter func([]byte) ([]byte, Mime, error), parameterName string, valueName string, commands ...string) *Command {
	var usedFormatter = formatter
	if usedFormatter == nil {
		usedFormatter = CommandFormatter
	}
	return &Command{commands, parameterName, valueName, formatter}
}

// Handler - Returns the Http Handler
func (cmd *Command) Handler() func(http.ResponseWriter, *http.Request) {
	return cmd.commandHandler
}

// StartHandler - Returns the result as array of strings
func (cmd *Command) commandHandler(w http.ResponseWriter, r *http.Request) {
	// check authentication
	err := auth.Check(r)
	if err != nil {
		errorMsg := "Authorizations check failed"
		log.Println(errorMsg)
		http.Error(w, errorMsg, 401)
		return
	}

	// determine commands
	commands, err := cmd.getCommands(r)
	if err != nil {
		errorMsg := "Commands could not be determined"
		log.Println(errorMsg)
		http.Error(w, errorMsg, 400)
		return
	}
	commandsStr := strings.Join(commands, " ")
	log.Println("Tinc Commands: ", commandsStr)

	// execute command
	out, err := exec.Command("tinc", commands...).CombinedOutput()
	if err != nil {
		log.Println("Execution of tinc command returned in error:", commandsStr)
		http.Error(w, getErrorText(string(out), err.Error()), 400)
		return
	}

	// format result
	result, contentType, err := cmd.formatter(out[:])
	if err != nil {
		log.Println("Formatting of result of tinc command returned in error")
		http.Error(w, getErrorText(string(out), err.Error()), 400)
		return
	}

	// return result
	w.Header().Set("Content-Type", string(contentType))
	w.Write(result)
}

// getCommands - Determine the commands from the Command and optionally from the url
func (cmd *Command) getCommands(r *http.Request) ([]string, error) {
	commands := cmd.commands[:]
	if cmd.parameterName != "" {
		var parNameValue = r.URL.Query().Get(cmd.parameterName)
		if parNameValue == "" {
			return commands, errors.New("The parameter is mandatory: " + cmd.parameterName)
		}
		commands = append(commands, parNameValue)
		if cmd.valueName != "" {
			parValueValue := r.URL.Query().Get(cmd.valueName)
			if parValueValue == "" {
				return commands, errors.New("The value is mandatory: " + cmd.valueName)
			}
			commands = append(commands, parValueValue)
		}
	}
	return commands, nil
}

func getErrorText(t1 string, t2 string) string {
	if t1 != "" {
		return t1
	}
	return t2
}
