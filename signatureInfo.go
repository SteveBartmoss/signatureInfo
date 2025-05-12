package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
	"strings"
)

type SystemInfo struct{
	OS string
	Arch string
	Hostname string
	Kernel string
	Uptime string
	Cpu string
	Gpu string
	Ram string
	Shell string
	User string
	Terminal string
}

func runCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Desconocido"
	}
	return strings.TrimSpace(string(output))
}

func getSystemInfo() SystemInfo {
	info := SystemInfo{
		OS: runtime.GOOS,
		Arch: runtime.GOARCH,
		Hostname: runCommand("hostname"),
		Kernel: runCommand("uname", "-sr"),
		Uptime: runCommand("uptime", "-p"),
		Cpu: runCommand("sh", "-c", "lscpu | grep 'Nombre del modelo' | cut -d ':' -f 2 | xargs"),
		Gpu: runCommand("sh", "-c", "lspci -v | grep -i vga | cut -d ':' -f 3 | xargs"),
		Ram: runCommand("sh", "-c", "free -h | awk '/Mem:/ {print $3 \"/\" $2}'"),
		Shell: os.Getenv("SHELL"),
		User: os.Getenv("USER"),
		Terminal: os.Getenv("TERM"),
	}
	return info
}

func displayInfo(info SystemInfo){

	// Definir colores ANSI
    c1 := "\033[36m"  // Cian
    c2 := "\033[35m"  // Magenta
    reset := "\033[0m" // Reset

    // ASCII art - versión corregida sin marcadores %s
    logo := `
` + c2 + `             ...-:::::-...` + reset + `
` + c2 + `          .-MMMMMMMMMMMMMMM-.` + reset + `
      .-MMMM` + c1 + `` + "`..-:::::::-..`" + c2 + `MMMM-.` + reset + `
    .:MMMM` + c1 + `.:MMMMMMMMMMMMMMM:.` + c2 + `MMMM:.` + reset + `
   -MMM` + c1 + `-M---MMMMMMMMMMMMMMMMMMM.` + c2 + `MMM-` + reset + `
 ` + "`" + `:MMM` + c1 + `:MM  :MMMM:....::-...-MMMM:` + c2 + `MMM:` + "`" + reset + `
 :MMM` + c1 + `:MMM  :MM:  ` + "``" + `    ` + "``" + `  ` + "`" + `:MMM:` + c2 + `MMM:` + reset + `
.MMM` + c1 + `.MMMM  :MM.  -MM.  .MM-  ` + "`" + `MMMM.` + c2 + `MMM.` + reset + `
:MMM` + c1 + `:MMMM  :MM.  -MM-  .MM:  ` + "`" + `MMMM-` + c2 + `MMM:` + reset + `
:MMM` + c1 + `:MMMM  :MM.  -MM-  .MM:  ` + "`" + `MMMM:` + c2 + `MMM:` + reset + `
:MMM` + c1 + `:MMMM  :MM.  -MM-  .MM:  ` + "`" + `MMMM-` + c2 + `MMM:` + reset + `
.MMM` + c1 + `.MMMM  :MM:--:MM:--:MM:  ` + "`" + `MMMM.` + c2 + `MMM.` + reset + `
 :MMM` + c1 + `:MMM-  ` + "`" + `-MMMMMMMMMMMM-` + "`" + `  -MMM-` + c2 + `MMM:` + reset + `
  :MMM` + c1 + `:MMM:` + `                ` + "`" + `:MMM:` + c2 + `MMM:` + reset + `
   .MMM` + c1 + `.MMMM:--------------:MMMM.` + c2 + `MMM.` + reset + `
     '-MMMM` + c1 + `.-MMMMMMMMMMMMMMM-.` + c2 + `MMMM-'` + reset + `
       '.-MMMM` + c1 + `` + "`" + `--:::::--` + "`" + c2 + `MMMM-.'` + reset + `
` + c2 + `            '-MMMMMMMMMMMMM-'` + reset + `
` + c2 + `               ` + "`" + `-:::::-` + "`" + reset + `
`

	
	fmt.Print(logo)
	fmt.Printf("%s %s\n", "OS:", info.OS)
	fmt.Printf("%s %s\n", "Host:", info.Hostname)
	fmt.Printf("%s %s\n", "Kernel:", info.Kernel)
	fmt.Printf("%s %s\n", "Uptime:", info.Uptime)
	fmt.Printf("%s %s\n", "Cpu:", info.Cpu)
	fmt.Printf("%s %s\n", "Gpu:", info.Gpu)
	fmt.Printf("%s %s\n", "Ram:", info.Ram)
	fmt.Printf("%s %s\n", "Shell:", info.Shell)
	fmt.Printf("%s %s\n", "User:", info.User)
	fmt.Printf("%s %s\n", "Terminal:", info.Terminal)

}

func main(){

	info := getSystemInfo()
	displayInfo(info)

	/*
	fmt.Println("OS:",runtime.GOOS)
	fmt.Println("Arquitectura:",runtime.GOARCH)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Desconocido"
	}

	cmd1 := exec.Command("uname","-a")
	kernelInfo, err := cmd1.CombinedOutput()
	if err != nil {
		kernelInfo = []byte("Desconocido")
	}

	cmd2 := exec.Command("cat","/proc/version")
	kernelInfoSecond, err := cmd2.CombinedOutput()
	if err != nil {
		kernelInfoSecond =  []byte("Desconocido")
	}

	cmd := exec.Command("uptime")
	uptimeInfo, err := cmd.CombinedOutput()
	if err != nil {
		uptimeInfo = []byte("Desconocido")
	}

	cmd = exec.Command("lscpu")
	cpuInfo, err := cmd.CombinedOutput()
	if err != nil {
		cpuInfo = []byte("Desconocido")
	}

	cmd = exec.Command("sh", "-c", "lspci -v | grep -i vga")
	gpuInfo, err := cmd.CombinedOutput()
	if err != nil {
		gpuInfo = []byte("Desconocido")
	}

	cmd = exec.Command("free","-m")
	ramInfo, err := cmd.CombinedOutput()
	if err != nil {
		ramInfo = []byte("Desconocido")
	}

	terminal := os.Getenv("TERM")
	
	fmt.Println("Host:",hostname)
	fmt.Println("Usuario:", os.Getenv("USER"))
	fmt.Println("Shell:", os.Getenv("SHELL"))
	fmt.Println("Info Kernel 1:",string(kernelInfo))
	fmt.Println("Info Kernel 2:",string(kernelInfoSecond))
	fmt.Println("uptime:",string(uptimeInfo))
	fmt.Println("cpu:",string(cpuInfo))
	fmt.Println("gpu:",string(gpuInfo))
	fmt.Println("ram:",string(ramInfo))
	fmt.Println("Terminal:",terminal)
	*/
}
