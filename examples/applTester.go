package main

import (
	"fmt"
	"log/syslog"
	"time"

	"github.com/marcovargas74/m74tester/appliance"
	rlog "github.com/marcovargas74/rLog"
)

const (
	//dEFtimeExec10Min Timer de 10 minutos
	deftimeExec10Min = 10
)

//Ttester Tipo Tester Variaveis Globais usados em todo o Tester
type Ttester struct {
	LogProgLevel      syslog.Priority
	LogProgEnable     bool
	LogProgPrintLocal bool
	LogProgIP         string
	VersionSoftware   string //Versao da aplicação Default
	PidMain           int
	PidRemoteSocket   int
	LOOPMain          bool
	//KillApplication   bool
	//Suporta
	//ServerHTTP bool
	//ServerDir  string
	//ServerPort string

	//Suporta Appliance
	//IsAppliance bool
	//ServerDir  string
	//ServerPort string

	//Suporta Menu
	//MenuEnable bool
}

//TesterCtrl Variaveis Globais usados em todo o Tester
var TesterCtrl Ttester

func init() {
	dt := time.Now()
	TesterCtrl.VersionSoftware = dt.Format("2006-01-02")
	//Configuracao do LOG
	TesterCtrl.LOOPMain = true
	TesterCtrl.LogProgEnable = true
	TesterCtrl.LogProgPrintLocal = true
	TesterCtrl.LogProgLevel = rlog.Debug | rlog.Local4
	TesterCtrl.LogProgIP = "172.31.11.162:514"

	//Configuração do Servidor HTTP
	//AppCtrl.ServerHTTP = false
	//AppCtrl.ServerDir = "appliance/public"
	//AppCtrl.ServerDir = "public"
	//AppCtrl.ServerDir = "public/appPage"
	//AppCtrl.ServerPort = ":8080"

	//Ativa Menu com opcoes
	//AppCtrl.MenuEnable = false

	//Ativa pacote do appliance
	//AppCtrl.IsAppliance = true

	rlog.Clear()
	rlog.StartLogger(TesterCtrl.LogProgEnable, TesterCtrl.LogProgLevel, TesterCtrl.LogProgIP)
	rlog.SetPrintLocal(TesterCtrl.LogProgPrintLocal)
	rlog.AppSyslog(syslog.LOG_INFO, "%s ======== Start Mannager App Version %s\n", rlog.ThisFunction(), TesterCtrl.VersionSoftware)

	//ret := appliance.execLinuxCmd("users")
	//fmt.Printf("return: %s\n", ret)

}

/*
 *
 */
//func executaTimerControl(ptrCountLOOP *uint64, ptrCount30s *byte, ptrCount1MIN *uint64) {
func executaTimerControl(ptrCountLOOP *uint64, ptrCount1MIN *uint64) {
	// 1 minuto
	time.Sleep(time.Minute)
	*ptrCount1MIN++ //Inclrementa 1 Minuto

	if (*ptrCount1MIN % deftimeExec10Min) == 0 {
		rlog.AppSyslog(syslog.LOG_DEBUG, "%s LOOP 10 min..[%d]min\n", rlog.ThisFunction(), *ptrCount1MIN)
	}

} //Excuta_timer_control

/*
FUNCAO INICIAL DO PROGRAMA
*/
func main() {
	var countLOOP uint64 //Contador de minuto
	var count1MIN uint64 //Contador de minuto

	fmt.Println("Tester is running...")
	rlog.AppSyslog(syslog.LOG_INFO, "%s {START LOOP_MAIN!!!}\n", rlog.ThisFunction())

	//Iniciar Pacote do Appliance
	go appliance.StartAppliance(appliance.GetMode())

	for {
		executaTimerControl(&countLOOP, &count1MIN)
		//Termina caso App mandou terminar
		if TesterCtrl.LOOPMain == false {
			break
		}
	} //for

	rlog.AppSyslog(syslog.LOG_DEBUG, "%s {FINISH LOOP_MAIN }\n", rlog.ThisFunction())
	rlog.LoggerClose()
}

/*
func execLinuxCmd(cmd string) []byte {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("return: %s\n", out)
	return out
}*/