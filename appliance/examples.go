package appliance

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

/*
//TApp Tipo Aplicativo Variaveis Globais usados em todo o Aplicativo
type TAppliance struct {
	LogProgLevel    syslog.Priority
	LogProgEnable   bool
	LogProgIP       string
	VersionSoftware string //Versao da aplicação Default
	//Suporta
	ServerHTTP bool
	ServerDir  string
	ServerPort string
}

//AppCtrl Variaveis Globais usados em todo o Aplicativo
var applianceCtrl TAppliance
func init() {

}
*/

func runTestes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Memoria", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "CPU", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "E1", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Rede 1", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Rede 2", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Rede 3", "OK")
	fmt.Fprintf(w, "<h1><font color='#2e802e' size='4'>INFO Teste %s %s</font></h1>", "Rede 4", "OK")
	fmt.Fprintln(w, "<h2>TODOS OS TESTES OK <h2>")
	GetDateNow(w, r)
}

//GetDateNow Mostra a data e hora
func GetDateNow(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprintf(w, "<h1>Data: %s<h1>", s)

	fmt.Println("GetDateNow at", s)

}

//ExampleTesteFunc funcao exemplo de como realizar testes
func ExampleTesteFunc(fileName string) string {
	fmt.Println(fileName)
	return fileName
}

//TReadFile Le arquivo
func TReadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.FormValue("nomeArquivo"))
}

//ReadFileErr Le arquivo
func ReadFileErr(filename string) ([]byte, error) {
	//filename := title + ".txt"
	//filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return body, nil
}

//RxDataFromJS recebe data
func RxDataFromJS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.FormValue("aData"))
}

//SendDataToJS recebe data
func SendDataToJS(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println(r.FormValue("aData"))
	fmt.Fprintf(w, "DataSendToJs:%d", 123)
}

//Imprime valores das memorias do sistema
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", m.Alloc) //filename := title + ".txt"
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc)
	fmt.Printf("\tSys = %v MiB", m.Sys)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

//UpdatePages Atualiza as paginas passa o local onde sera feito os
/*func UpdatePages() {

	fileURL := "/home/intelbras/projetos-go/src/github.com/ma022800/goma/public/*"

	if err := DownloadFile("appPage", fileURL); err != nil {
		panic(err)
	}
}* /

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
*/

func showUname() {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		fmt.Printf("Uname: %v", err)
	}
	fmt.Println(arrayToString(uname.Nodename))
	fmt.Println(arrayToString(uname.Release))
	fmt.Println(arrayToString(uname.Sysname))
	fmt.Println(arrayToString(uname.Version))
	fmt.Println(arrayToString(uname.Machine))
	fmt.Println(arrayToString(uname.Domainname))
}

func arrayToString(x [65]int8) string {
	var buf [65]byte
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}

func getHome() {
	//home := execLinuxCmd("users")
	if home := os.Getenv("HOME"); home != "" {
		fmt.Println("home:", home)
	}
	//home, _ := exec.Command("users").Output()

}


/*
//ExecLinuxCmd execupa comando linux passando como string
func execLinuxCmd(cmd string) []byte {
	out, err := exec.Command(cmd).Output()
	check(err)
	//fmt.Printf("return: %s\n", out)
	return out
}*/

/*
func execLinuxCmd(cmd string) ([]byte, error) {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//fmt.Printf("return: %s\n", out)
	return out, nil
}*/


/*
func showUsb() {
	ctx := gousb.NewContext()
	defer ctx.Close()

}*/

/*
var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

func showUsbs() {
	flag.Parse()

	// Only one context should be needed for an application.  It should always be closed.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.Debug(*debug)

	// OpenDevices is used to find the devices to open.
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// The usbid package can be used to print out human readable information.
		fmt.Printf("%03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))
		fmt.Printf("  Protocol: %s\n", usbid.Classify(desc))

		// The configurations can be examined from the DeviceDesc, though they can only
		// be set once the device is opened.  All configuration references must be closed,
		// to free up the memory in libusb.
		for _, cfg := range desc.Configs {
			// This loop just uses more of the built-in and usbid pretty printing to list
			// the USB devices.
			fmt.Printf("  %s:\n", cfg)
			for _, intf := range cfg.Interfaces {
				fmt.Printf("    --------------\n")
				for _, ifSetting := range intf.AltSettings {
					fmt.Printf("    %s\n", ifSetting)
					fmt.Printf("      %s\n", usbid.Classify(ifSetting))
					for _, end := range ifSetting.Endpoints {
						fmt.Printf("      %s\n", end)
					}
				}
			}
			fmt.Printf("    --------------\n")
		}

		// After inspecting the descriptor, return true or false depending on whether
		// the device is "interesting" or not.  Any descriptor for which true is returned
		// opens a Device which is retuned in a slice (and must be subsequently closed).
		return false
	})

	// All Devices returned from OpenDevices must be closed.
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	// OpenDevices can occasionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
	}

	for _, dev := range devs {
		// Once the device has been selected from OpenDevices, it is opened
		// and can be interacted with.
		_ = dev
	}
}
*/
