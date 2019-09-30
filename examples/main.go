package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ma022800/goma/goma"
	"github.com/ma022800/goma/goma/appliance"
	"github.com/marcovargas74/configSimulator/includes"
	rlog "github.com/marcovargas74/rLog"
)

//AppCtrl Variaveis Globais usados em todo o Aplicativo
var AppCtrl goma.TApp

func init() {
	dt := time.Now()
	AppCtrl.VersionSoftware = dt.Format("2006-01-02")
	//Configuracao do LOG
	AppCtrl.LOOPMain = true
	AppCtrl.LogProgEnable = true
	AppCtrl.LogProgPrintLocal = true
	AppCtrl.LogProgLevel = rlog.Debug | rlog.Local4
	AppCtrl.LogProgIP = "172.31.11.162:514"

	//Configuração do Servidor HTTP
	AppCtrl.ServerHTTP = false
	//AppCtrl.ServerDir = "appliance/public"
	//AppCtrl.ServerDir = "public"
	AppCtrl.ServerDir = "public/appPage"
	AppCtrl.ServerPort = ":8080"

	//Ativa Menu com opcoes
	AppCtrl.MenuEnable = false

	//Ativa pacote do appliance
	AppCtrl.IsAppliance = true

	rlog.Clear()
	rlog.StartLogger(AppCtrl.LogProgEnable, AppCtrl.LogProgLevel, AppCtrl.LogProgIP)
	rlog.SetPrintLocal(AppCtrl.LogProgPrintLocal)
	rlog.AppSyslog(syslog.LOG_INFO, "%s ======== Start Mannager App Version %s\n", rlog.ThisFunction(), AppCtrl.VersionSoftware)

	ret := execLinuxCmd("users")
	fmt.Printf("return: %s\n", ret)

}

/*
 *
 */
//func executaTimerControl(ptrCountLOOP *uint64, ptrCount30s *byte, ptrCount1MIN *uint64) {
func executaTimerControl(ptrCountLOOP *uint64, ptrCount1MIN *uint64) {
	// 1 minuto
	time.Sleep(time.Minute)
	*ptrCount1MIN++ //Inclrementa 1 Minuto

	if (*ptrCount1MIN % includes.DEFtimeExec10Min) == 0 {
		rlog.AppSyslog(syslog.LOG_DEBUG, "%s LOOP 10 min..[%d]min\n", rlog.ThisFunction(), *ptrCount1MIN)
	}

} //Excuta_timer_control

/*
FUNCAO INICIAL DO PROGRAMA
*/
func main() {
	var countLOOP uint64 //Contador de minuto
	var count1MIN uint64 //Contador de minuto

	//appliance.UpdatePages()
	fmt.Println("GoMa is running...")
	//config := inicialConfigs(1)
	//fmt.Println(config)

	rlog.AppSyslog(syslog.LOG_INFO, "%s {START LOOP_MAIN!!!}\n", rlog.ThisFunction())

	//Iniciar Pacote do Appliance
	if AppCtrl.IsAppliance {
		go appliance.StartAppliance(getMode())
	}

	for {
		executaTimerControl(&countLOOP, &count1MIN)
		//Termina caso App mandou terminar
		if AppCtrl.LOOPMain == false {
			break
		}
	} //for

	rlog.AppSyslog(syslog.LOG_DEBUG, "%s {FINISH LOOP_MAIN }\n", rlog.ThisFunction())
	rlog.LoggerClose()
}

func execLinuxCmd(cmd string) []byte {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("return: %s\n", out)
	return out
}

func getMode() string {
	//home := execLinuxCmd("users")
	/*if home := os.Getenv("HOME"); home != "" {
		fmt.Println("home:", home)
	}*/

	home, _ := os.UserHomeDir()
	fmt.Println("home:", home)

	if strings.Contains(string(execLinuxCmd("users")[:]), "intelbras") {
		//fmt.Println("dev")
		return "dev"
	}

	//fmt.Println("modo is prod")
	return "prod"
}
