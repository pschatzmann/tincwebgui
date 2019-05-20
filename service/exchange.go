package service

/**
*  Parser for tinc dump commands: It Parses a string with the help of an array
*  of delimiter fields into a Map
**/
import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// RemoteExchange - exchange the host configuration files
func RemoteExchange(w http.ResponseWriter, r *http.Request) {
	log.Println("RemoteExchange")

	url := getURL(r)
	if url != "" {
		http.Error(w, "The url parameter is not defined", 400)
		return
	}

	importRemoteConfig(url, w, r)
	exportLocalConfig(url, w, r)

	// success message
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("The host configurations were successfully exchanged"))
}

func importRemoteConfig(url string, w http.ResponseWriter, r *http.Request) {
	// import the remote configuration file
	cmd := exec.Command("tinc", "import")
	config, err := getRemoteConfiguration(url + "/export")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	cmd.Stdin = strings.NewReader(config)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

// get the configuration file from the remote system
func getRemoteConfiguration(url string) (string, error) {
	log.Println("URL:", url)

	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return string(body), nil
}

func exportLocalConfig(url string, w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("tinc", "export").CombinedOutput()
	if err != nil {
		http.Error(w, getErrorText(string(out), err.Error()), 400)
		return
	}
	err = writeRemoteConfiguration(url, string(out))
	if err != nil {
		http.Error(w, getErrorText(string(out), err.Error()), 400)
		return
	}
}

// post the configuration file to the remote system
func writeRemoteConfiguration(url string, configData string) error {
	log.Println("URL:", url)

	req, err := http.NewRequest("POST", url+"/import", strings.NewReader(configData))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return nil
}

func getURL(r *http.Request) string {
	result := getHTTPParameterValue(r, "url")
	if result == "" {
		return ""
	}
	if !strings.HasPrefix(result, "http") {
		xURL := r.URL
		scheme := xURL.Scheme
		result = scheme + result
	}
	if !strings.HasSuffix(result, "api") {
		if !strings.HasSuffix(result, "/") {
			result = result + "/"
		}
		result = result + "api"
	}
	return result
}
