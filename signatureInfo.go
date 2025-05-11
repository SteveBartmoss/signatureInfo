package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
)

func main(){
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
	
	fmt.Println("Host:",hostname)
	fmt.Println("Usuario:", os.Getenv("USER"))
	fmt.Println("Shell:", os.Getenv("SHELL"))
	fmt.Println("Info Kernel 1:",string(kernelInfo))
	fmt.Println("Info Kernel 2:",string(kernelInfoSecond))
	fmt.Println("uptime:",string(uptimeInfo))
	fmt.Println("cpu:",string(cpuInfo))
	fmt.Println("gpu:",string(gpuInfo))
	fmt.Println("ram:",string(ramInfo))

}
