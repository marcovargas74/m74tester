package appliance

import (
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"os/exec"
	"strings"

	rlog "github.com/marcovargas74/rLog"
)

const (
	serverDirDev  = "/home/intelbras/projetos-go/src/github.com/marcovargas74/m74tester/appliance/public"
	serverDirProd = "/home/iap/appliance/public"
	serverPort    = ":8080"
	//mode          = "prod"
	//mode             = "dev"

	//fileHardConfDev  = "/home/intelbras/projetos-go/src/github.com/marcovargas74/m74tester/appliance/public/static/hard.conf"
	fileHardConfDev  = serverDirDev + "/static/hard.conf"
	fileHardConfProd = serverDirProd + "/static/hard.conf"
)

//Mode é o modo do teste se "prod" ou "dev"
var Mode = "prod"

//StartAppliance inicia o servidor http do appliance
func StartAppliance(modo string) {

	if modo != "" {
		Mode = modo
	}

	serverDir := serverDirProd
	if Mode == "dev" {
		serverDir = serverDirDev
	}

	server := http.FileServer(http.Dir(serverDir))
	http.Handle("/", server)

	HandleFuncions()

	//MODO_ := "dev" //prod
	rlog.AppSyslog(syslog.LOG_INFO, "%s %s{START SERVER HTTP!!!}\n", rlog.ThisFunction(), Mode)
	log.Fatal(http.ListenAndServe(serverPort, nil))

}

//HandleFuncions Prepara funcoe para serem usadas
func HandleFuncions() {
	http.HandleFunc("/date", GetDateNow)
	http.HandleFunc("/testes", runTestes)
	http.HandleFunc("/rxdata", RxDataFromJS)
	http.HandleFunc("/txdata", SendDataToJS)
	http.HandleFunc("/readfile", ReadFile)
	http.HandleFunc("/iniselftest", SelfTestIni)
	http.HandleFunc("/selftest", SelfTest)

}

//ReadFile Le arquivo
func ReadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("nomeArquivo")
	//filename := title + ".txt"
	//filename = "title"
	//Nao da bola para o dado que vem do js
	filename = fileHardConfProd
	if Mode == "dev" {
		filename = fileHardConfDev
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", filename, "Err:", err)
		//return nil
		formatMessage(w, "ERR Erro ao abrir Arquivo de configuracao")
		return
	}

	fmt.Fprintf(w, "%s", body)
	fmt.Println("ReadFile: ", filename)
	fmt.Println("Body: ", string(body))
	//fmt.Fprintf(w, "<resposta>0</resposta>")
}

//SelfTestIni inicia os testes
func SelfTestIni(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println(r.FormValue("aData"))
	//fmt.Fprintf(w, "SelfTestIni OK")
	//fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Memoria", "OK")

	erro := 0
	testName := "config"
	if Mode == "dev" {
		//$test = $_POST['x'];
		testName = "hardware"
		//exec := "sh /home/intelbras/ICIP/Firmwares/Scripts/selftest.sh $test";
	}

	//HoraCerta(w, r)
	fmt.Fprintf(w, "<valor><font color='#2e802e' size='4'>INFO Teste de %s</font></valor>", testName)
	if erro != 0 {
		fmt.Fprintf(w, "<resposta>1</resposta>")
	} else {
		fmt.Fprintf(w, "<resposta>0</resposta>")
	}

	fmt.Println("SelfTestIni OK ")
	//if (ODO_ == "prod" {
	///$test = $_POST['x'];
	///$exec = "sh /usr/bin/selftest.sh $test";

	//showInterfaces()
}

//func formatMessage(w http.ResponseWriter, message string) string {
func formatMessage(w http.ResponseWriter, format string, a ...interface{}) {

	message := fmt.Sprintf(format, a...)

	//default: PRETA
	color := "#000000"
	erro := 0

	//#INFO Verde
	if strings.Contains(message, "INFO") {
		color = "#2e802e"
	}

	//#OK Azul
	if strings.Contains(message, "OK") {
		color = "#0066FF"
	}

	//#WARN Laranja
	if strings.Contains(message, "WARN") {
		color = "#FF9900"
	}

	//#ERR vermelho
	if strings.Contains(message, "ERR") {
		color = "#FF0000"
		erro = 1
	}

	fmt.Fprintf(w, "<valor><font color='%s' size='3'>\t%s</font></valor>", color, message)

	if erro != 0 {
		fmt.Fprintf(w, "<resposta>1</resposta>")
	} else {
		fmt.Fprintf(w, "<resposta>0</resposta>")
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		//panic(e)
	}
}

//GetMode return the type of teste development or production
func GetMode() string {
	//home := execLinuxCmd("users")
	/*if home := os.Getenv("HOME"); home != "" {
		fmt.Println("home:", home)
	}*/
	//home, err := exec.Command("users").Output()

	home, err := os.UserHomeDir()
	check(err)
	fmt.Println("home:", home)
	//if strings.Contains(string(execLinuxCmd("users")[:]), "intelbras") {
	if strings.Contains(home, "intelbras") {
		//fmt.Println("dev")
		return "dev"
	}

	//fmt.Println("modo is prod")
	return "prod"
}

func showUsbs() string {

	out, err := exec.Command("lsusb").Output()
	check(err)
	//out := execLinuxCmd("lsusb")
	message := string(out[:])
	return message
}


/* //////////////////////// LIXO //////////////////////////////////////////////////////////////////////

if ($MODO_ == 'dev')
{
 //$test = $_POST['x'];
 $test = "hardware";
 $exec = "sh /home/intelbras/ICIP/Firmwares/Scripts/selftest.sh $test";
}

if ($MODO_ == 'prod')
{
 $test = $_POST['x'];
 $exec = "sh /usr/bin/selftest.sh $test";
}

$erro = 0;
$name=strtoupper($test);
echo "<valor><font color='#2e802e' size='4'>INFO Teste de $name</font></valor>";
flush();

exec($exec,$out);

if($out[count($out)-1] != 0)
 {
   $erro++;
 }

for($i=0;$i<count($out)-1;$i++){
	   $start=$out[$i][1];
	   switch ($start) {
		   case "E":
		   case "R":
			   #ERR vermelho
			   $color = "#FF0000";
			   break;
		   case "W":
		   case "A":
			   #WARN Laranja
			   $color = "#FF9900";
			   break;
		   case "I":
		   case "N":
			   $color = "#2e802e";
			   break;
		   case "O":
		   case "K":
			   #INFO e OK Azul
			   $color = "#0066FF";
			   break;
		   default:
			   $color = "#000000";
	   }
	   echo "<valor><font color='$color' size='3'>\t$out[$i]</font></valor>";
	   flush();
}

if($erro != 0){
   echo "<resposta>1</resposta>";
} else {
   echo"<resposta>0</resposta>";
}
*/
