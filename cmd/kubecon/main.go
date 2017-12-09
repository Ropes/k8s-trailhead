package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	log "github.com/Sirupsen/logrus"

	"github.com/bmizerany/pat"
)

func helloKubecon(w http.ResponseWriter, req *http.Request) {
	str := `
  __  ___  __    __   ______    _______   ______   ______    __   __   __  
 |  |/  / |  |  |  | |   _  \  |   ____| /      | /  __  \  |  \ |  | |  | 
 |  '  /  |  |  |  | |  |_)  | |  |__   |  ,----'|  |  |  | |   \|  | |  | 
 |    <   |  |  |  | |   _  <  |   __|  |  |     |  |  |  | |  .    | |  | 
 |  .  \  |  '--'  | |  |_)  | |  |____ |  '----.|  '--'  | |  |\   | |__| 
 |__|\__\  \______/  |______/  |_______| \______| \______/  |__| \__| (__) 

 Greetings Kubernaughts from %s %s

 Num CPU: %d
 GOMAXPROCS(): %d
 GOMAXPROCS[env]: %q
`
	usr, ok := os.LookupEnv("USER")
	if !ok {
		log.Warnf("user envvar undefined")
	}
	host, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		log.Warnf("hostname envvar undefined")
	}
	resp := fmt.Sprintf(str, usr, host, runtime.NumCPU(), runtime.GOMAXPROCS(-1), os.Getenv("GOMAXPROCS"))
	io.WriteString(w, resp)
}

func main() {
	m := pat.New()
	m.Get("/kubecon", http.HandlerFunc(helloKubecon))

	http.Handle("/", m)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Errorf("http error: %v", err)
	}
}
