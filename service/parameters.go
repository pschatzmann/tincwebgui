package service

/**
*  Handler for /api/parameters
**/
import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/pschatzmann/tincwebgui/auth"
)

type parameter struct {
	Name  string
	Value string
}

// TincParameters - Handle /api/parameters
func TincParameters(w http.ResponseWriter, r *http.Request) {
	err := auth.Check(r)
	if err != nil {
		log.Println("Authorizations check failed")
		http.Error(w, err.Error(), 400)
		return
	}

	var parametersKeys = listParameterKeys()
	var result = make([]parameter, len(parametersKeys))

	for i, v := range parametersKeys {
		parValue := new(parameter)
		parValue.Name = v
		parValue.Value = getParameterValue(parValue.Name)
		result[i] = *parValue
	}

	b, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("parameters")
	w.Write(b)
}

// get the parameter value by calling the get command on the tinc interpreter
func getParameterValue(parameterName string) string {
	// execute tinc command
	cmd := exec.Command("tinc", "get", parameterName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		var errorTxt = err.Error()
		if errorTxt == "exit status 1" {
			return "-"
		}
		return errorTxt
	}
	var resultStr = string(out)
	return strings.TrimSpace(resultStr)
}

// determines all parameter names
func listParameterKeys() []string {
	var parameters = []string{
		"AddressFamily",
		"AutoConnect",
		"BindToAddress",
		"BindToInterface",
		"Broadcast",
		"BroadcastSubnet",
		"ConnectTo",
		"DecrementTTL",
		"Device",
		"DeviceStandby",
		"DeviceType",
		"DirectOnly",
		"Ed25519PrivateKeyFile",
		"ExperimentalProtocol",
		"Forwarding",
		"FWMark",
		"GraphDumpFile",
		"Hostnames",
		"IffOneQueue",
		"Interface",
		"InvitationExpire",
		"KeyExpire",
		"ListenAddress",
		"LocalDiscovery",
		"LogLevel",
		"MACExpire",
		"MaxConnectionBurst",
		"MaxOutputBufferSize",
		"MaxTimeout",
		"Mode",
		"Name",
		"PingInterval",
		"PingTimeout",
		"PriorityInheritance",
		"PrivateKey",
		"PrivateKeyFile",
		"ProcessPriority",
		"Proxy",
		"ReplayWindow",
		"ScriptsExtension",
		"ScriptsInterpreter",
		"StrictSubnets",
		"TunnelServer",
		"UDPDiscovery",
		"UDPDiscoveryKeepaliveInterval",
		"UDPDiscoveryInterval",
		"UDPDiscoveryTimeout",
		"MTUInfoInterval",
		"UDPInfoInterval",
		"UDPRcvBuf",
		"UDPSndBuf",
		"UPnP",
		"UPnPDiscoverWait",
		"UPnPRefreshPeriod",
		"VDEGroup",
		"VDEPort",
		/* Host configuration */
		"Address",
		"Cipher",
		"ClampMSS",
		"Compression",
		"Digest",
		"Ed25519PublicKey",
		"Ed25519PublicKeyFile",
		"IndirectData",
		"MACLength",
		"PMTU",
		"PMTUDiscovery",
		"Port",
		"PublicKey",
		"PublicKeyFile",
		"Subnet",
		"TCPOnly",
		"Weight"}

	return parameters

}