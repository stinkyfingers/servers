package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/enroll", enrollHandler)
	mux.HandleFunc("/logger", loggerHandler)
	mux.HandleFunc("/config", configHandler)
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":8889", "server.crt", "server.key", mux))
}

type EnrollmentResponse struct {
	NodeKey     string `json:"node_key"`
	NodeInvalid bool   `json:"node_invalid"`
}

type Config struct {
	Options Options `json:"options"`
}
type Options struct {
	ConfigFileName  string `flag:"-" json:"-"`
	FlagFileName    string `flag:"-" json:"-"`
	ConfigPath      string `flag:"config_path" json:"config_path"`
	DatabasePath    string `flag:"database_path" json:"database_path"`
	TlsHostname     string `flag:"tls_hostname" json:"tls_hostname"`
	TlsServerCerts  string `flag:"tls_server_certs" json:"tls_server_certs"`
	EnrollSecretEnv string `flag:"enroll_secret_env" json:"enroll_secret_env"`

	ConfigPlugin             string `flag:"config_plugin" json:"config_plugin"`
	ConfigTlsEndpoint        string `flag:"config_tls_endpoint" json:"config_tls_endpoint"`
	ConfigRefresh            string `flag:"config_refresh" json:"config_refresh"`
	ConfigAcceleratedRefresh string `flag:"config_accelerated_refresh" json:"config_accelerated_refresh"`

	LoggerPlugin      string `flag:"logger_plugin" json:"logger_plugin"`
	LoggerTlsEndpoint string `flag:"logger_tls_endpoint" json:"logger_tls_endpoint"`

	EnrollTlsEndpoint string `flag:"enroll_tls_endpoint" json:"enroll_tls_endpoint"`

	PidFile string `flag:"pidfile" json:"pidfile"`

	DisableExtensions  bool   `flag:"-" json:"disable_extensions"`
	ExtensionsAutoload string `flag:"extensions_autoload" json:"extensions_autoload"`
	ExtensionsSocket   string `flag:"extensions_socket" json:"extensions_socket"`

	ScheduleSplayPercent string `flag:"schedule_splay_percent" json:"schedule_splay_percent"`
	DisableTables        bool   `flag:"disable_tables" json:"disable_tables"`
	UTC                  bool   `flag:"utc" json:"utc"`
}

func enrollHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print("enroll ", string(b))
	er := EnrollmentResponse{NodeKey: "key"}
	j, _ := json.Marshal(er)
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}

func loggerHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print("logger ", string(b))
	w.Write([]byte("{}"))
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print("config ", string(b))
	cnf, _ := ioutil.ReadFile("osquery.conf")
	log.Print("sending config ", string(cnf))
	w.Header().Add("Content-Type", "application/json")
	w.Write(cnf)
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print(string(b))
	w.Write([]byte("OK"))
}
