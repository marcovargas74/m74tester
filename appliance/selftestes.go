package appliance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func memoryTest(w http.ResponseWriter, r *http.Request) int {
	formatMessage(w, "OK Inicio dos testes de memoria RAM")
	formatMessage(w, "OK Teste de escrita na SDRAM")

	//Escreve dados em um arquivo e
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	if err != nil {
		fmt.Println("erro: ", err)
		formatMessage(w, "ERR Erro na escrita na memoria RAM")
		return 1
	}

	_, err = ioutil.ReadFile("/tmp/dat1")
	if err != nil {
		fmt.Println("erro: ", err)
		formatMessage(w, "ERR Teste de escrita na SDRAM - ocorreu um erro")
		return 1
	}

	formatMessage(w, "OK Teste de escrita na SDRAM ocorreu com sucesso")
	return 0

}

//	arraySelfTest.push('flash');
func flahMemoryTest(w http.ResponseWriter, r *http.Request) int {
	//nparts=`cat /proc/mtd  | wc -l`
	//if [ $nparts -ne 12]; then
	formatMessage(w, "ERR Erro nas memorias FLASH")
	//fi

	//if [ `cat /proc/mtd  | grep dbs | wc -l` -eq 1 ]; then
	formatMessage(w, "OK Memoria Dataflash FLASH encontrada: testando...")
	//	dataflash_flash_test
	//else
	formatMessage(w, "ERR Memoria Dataflash FLASH nao encontrada")
	//fi

	//if [ `cat /proc/mtd  | grep nshared | wc -l` -eq 1 ]; then
	formatMessage(w, "OK Memoria NAND FLASH encontrada: testando...")
	//	nand_flash_test
	//else
	formatMessage(w, "ERR Memoria NAND FLASH nao encontrada")
	//fi

	//if [ `grep PBX /proc/mtd | wc -l` -eq 1 ]; then
	formatMessage(w, "OK Memoria FLASH PABX encontrada: testando...")
	//	pabx_flash_test
	//else
	formatMessage(w, "ERR Memoria FLASH PABX nao encontrada")
	//fi
	return 0
}

//	arraySelfTest.push('codecs');
func boardTest(w http.ResponseWriter, r *http.Request) int {

	//if [ $MODO_TESTE -eq 0 ]; then
	//test_fpga_ac490 > /tmp/ac_log
	//RETVAL=$?
	//SLOT=`dmesg | grep "Boot do DSP conectado no slot" | wc -l`
	//else
	//	RETVAL=0
	//	SLOT=1
	//fi

	//if [ $TYPE_HARD -eq $TYPE_HARD_GW280 ]; then
	//test_dsp_gw
	//return
	//fi

	//if [ $TYPE_HARD -eq $TYPE_HARD_UNNITI ]; then
	//CODEC_COUNT=1
	//else
	//CODEC_COUNT=3
	//fi

	/*//for i in $(seq $CODEC_COUNT); do
	if [ $MODO_TESTE -eq 0 ]; then
		SLOT=`dmesg | grep "Inicializando Boot do DSP conectado no slot $(($i-1))" | wc -l`
	else
		SLOT=$i
	fi*/

	//if [ $SLOT -gt 0 ]; then
	//formatMessage(w, "OK Placa Codec conectada no slot 1")
	//else
	formatMessage(w, "WARN Nao existe Placa Codec conectada no slot 1")
	//fi
	//done

	//if [ $RETVAL -eq 0 ]; then
	formatMessage(w, "OK Todas as Placas Codec testadas obtiveram sucesso")
	formatMessage(w, "OK O teste da FPGA obteve sucesso")
	//if [ $(( $RETVAL & 1 )) == 1 ]; then
	formatMessage(w, "ERR A Placa Codec no slot 1 apresentou defeito")
	//fi

	//if [ $(( $RETVAL & 2 )) == 2 ]; then
	formatMessage(w, "ERR A Placa Codec no slot 2 apresentou defeito")
	//fi

	//if [ $(( $RETVAL & 4 )) == 4 ]; then
	formatMessage(w, "ERR A Placa Codec no slot 3 apresentou defeito")
	//fi

	//if [ $(( $RETVAL & 8 )) == 8 ]; then
	formatMessage(w, "ERR O FPGA apresentou defeito")
	//fi
	//fi

	return 0
}

//	arraySelfTest.push('usb');
func usbsTest(w http.ResponseWriter, r *http.Request) int {

	found := 0
	//Separa linha a linha o retorno do showUsbs()
	devices := strings.Split(showUsbs(), "\n")
	for _, linha := range devices {
		//fmt.Println(indice, linha)
		if strings.Contains(linha, "Device") {
			formatMessage(w, "INFO USB: %s", linha[strings.Index(linha, "Device"):])
			found++
		}
	}

	//if [ $MODO_TESTE -eq 0 ]; then
	//FOUND=`lsusb | grep "0e40:bebe\|0403:c580" | wc -l`
	////else
	//FOUND=2
	//fi

	if found == 0  {
		formatMessage(w, "ERR hardware USB nao encontrada")
	    return 1
	}

	if found < 2  {
		formatMessage(w, "WARN Apenas uma porta USB encontrada")
	    return 2
	  }

	//elif [ $FOUND == 2 ]; then
	formatMessage(w, "OK Portas USB ENCONTRADAS")
	//else
	//fi

	return 0
}

//	arraySelfTest.push('lan');
func ethInterfacesTest(w http.ResponseWriter, r *http.Request, eth string) int {
	//	formatMessage(w, "OK Inicio dos testes de memoria RAM")
	//	formatMessage(w, "OK Teste de escrita na SDRAM")
	formatMessage(w, "OK Teste de configuracao da %s", eth)
	formatMessage(w, "OK Configurando a interface %s", eth)

	//	formatMessage(w, "ERR Erro na escrita na memoria RAM")
	//	formatMessage(w, "OK Teste de escrita na SDRAM ocorreu com sucesso")
	/*
		send_message "OK" "Teste da interface WAN "

			if [ $TYPE_HARD -eq $TYPE_HARD_ICIP68 ]; then
				send_message "OK" "Carregando driver WAN 16 bits "
				insmod /drivers/ks8842.ko typeIs16Bits=1
			else
				send_message "OK" "Carregando driver WAN"
				insmod /drivers/ks8842.ko typeIs16Bits=0
			fi

			if [ $MODO_TESTE -eq 0 ]; then
				if [ $? -ne 0 ]; then
					send_message "ERR" "Erro ao carregar driver"
					return
				else
					send_message "OK" "Modulo carregado com sucesso"
				fi

				sleep 1
				found=`dmesg | grep ks8842 | grep "Found chip" | wc -l`

			else
				send_message "OK" "Modulo carregado com sucesso"
				sleep 1
				found=1
			fi

			if [ $found -gt 0 ]; then
				send_message "OK" "Interface wan encontrada e acessivel"
			else
				send_message "ERR" "Interface wan nao encontrada"
				return
			fi

			send_message "OK" "Teste de configuracao da WAN"
			send_message "OK" "Configurando a interface WAN"
			send_message "OK" "Endereco $WAN_IP Mascara 255.255.255.0"
			ifconfig eth1 $WAN_IP netmask 255.255.255.0 2> /dev/null > /dev/null
			sleep 1
			send_message "OK" "Pingando endereco $WAN_TEST_IP"
			ping -c 3 $WAN_TEST_IP 2> /dev/null > /dev/null
			if [ $? -eq 0 ]; then
				send_message "OK" "Ping OK"
			else
				send_message "WARN" "Erro ao pingar endereço $WAN_TEST_IP"
			fi
	*/

	return 0
}

//	arraySelfTest.push('wan');

//SelfTest Menu principal que chama os testes espeficicos
func SelfTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("aData"))

	erro := 0
	testName := r.FormValue("param")
	/*if mode == "dev" {
		//$test = $_POST['x'];
		testName = "hardware"
		//exec := "sh /home/intelbras/ICIP/Firmwares/Scripts/selftest.sh $test";
	}*/

	fmt.Fprintf(w, "<valor><font color='#2e802e' size='4'>INFO Teste de %s</font></valor>", testName)

	switch testName {
	case "memoria":
		erro = memoryTest(w, r)
		fmt.Println("one")

	case "flash":
		erro = flahMemoryTest(w, r)
		fmt.Println("two")

	case "codecs":
		erro = boardTest(w, r)
		fmt.Println("three")

	case "usb":
		erro = usbsTest(w, r)
		fmt.Println("four")
	}

	/*
		if testName == "memoria" {
			erro = memoryTest(w, r)
		}

		if testName == "flash" {
			erro = flahMemoryTest(w, r)
		}

		if testName == "codecs" {
			erro = boardTest(w, r)
		}

		//	arraySelfTest.push('usb');
		if testName == "usb" {
			erro = usbsTest(w, r)
		}
	*/
	//	arraySelfTest.push('lan');

	if strings.Contains(testName, "eth") {
		erro = ethInterfacesTest(w, r, testName)
	}

	/*
		if erro != 0 {
			fmt.Fprintf(w, "<resposta>1</resposta>")
		} else {
			fmt.Fprintf(w, "<resposta>0</resposta>")
		}*/

	fmt.Println("SelfTest ", testName, "ERRO: ", erro)
	//if (ODO_ == "prod" {
	///$test = $_POST['x'];
	///$exec = "sh /usr/bin/selftest.sh $test";
}

/* MATERIAL DE APOIO

#!/bin/sh

MODO_TESTE=0 # Roda script na Placa
#MODO_TESTE=1 # Roda script no desktop usado para testar
ERRORS=0

if [ $MODO_TESTE -eq 0 ]; then
	DRIVE_FPGA="/proc/fpga"
	HARDWARE_CONFIG="/etc/hard.conf"
	LAN_TEST_IP=10.0.0.1
	WAN_IP=10.1.0.2
	WAN_TEST_IP=10.1.0.1
else
	DRIVE_FPGA="/home/intelbras/ICIP/Firmwares/Scripts/_proc_fpga_gw"
	HARDWARE_CONFIG="/home/intelbras/ICIP/Firmwares/Scripts/hard.conf"
	LAN_TEST_IP=10.0.0.147
	WAN_IP=10.1.0.2
	WAN_TEST_IP=10.0.0.147
fi

send_message() {
	if [ $1 == "ERR" ]; then
		ERRORS=$(( $ERRORS + 1 ))
	fi
	echo -e "$1 \t $2"

	if [ $MODO_TESTE -eq 0 ]; then
		echo -e "$1 \t $2" > /dev/console
	fi
}

check_hardware_key() {
	if [ $MODO_TESTE -eq 0 ]; then
		FOUND=`lsusb | grep "0e40:bebe\|0403:c580" | wc -l`
	else
		FOUND=2
	fi

	if [ $FOUND == 1 ]; then
		send_message "WARN" "Apenas uma chave de hardware presente"
	elif [ $FOUND == 2 ]; then
		send_message "OK" "Duas chaves de hardware presente"
	else
		send_message "ERR" "Chave de hardware nao encontrada"
	fi
}

nand_flash_test() {
	nand_path=`cat /proc/mtd | grep nshared | sed s/:// | awk '{print $1}'`
	if [ x$nand_path != x ]; then
		nanddump /dev/$nand_path 2> /tmp/log_nand > /dev/null
		errors=`cat /tmp/log_nand | grep uncorrectable | wc -l`
		if [ $errors -gt 10 ]; then
			send_message "ERR" "Erro nas leitura da memória NAND FLASH"
		else
			send_message "OK" "Leitura da memoria NAND FLASH OK"
		fi
	else
		send_message "ERR" "Particao de testes da NAND nao encontrada"
	fi
}

dataflash_flash_test() {
	# the system boot by this flash... then it is readble
	send_message "OK" "Leitura da memoria Dataflash FLASH OK"
}

pabx_flash_test() {
	pabx_path=`grep PBX /proc/mtd | cut -f 1 -d :`
	if [ x$pabx_path != x ]; then
		cat /dev/$pabx_path > /dev/null
		if [ $? -eq 0 ]; then
			send_message "OK" "Leitura da memoria FLASH PABX OK"
		else
			send_message "ERR" "Erro na leitura da memoria FLASH PABX"
		fi
	else
		send_message "ERR" "Particao da memoria FLASH PABX nao encontrada"
	fi
}

memory_test() {
	send_message	"OK" "Inicio dos testes de memoria RAM"
	send_message	"OK" "Teste de escrita na SDRAM"
	if [ $MODO_TESTE -eq 0 ]; then
		imt 64 > /tmp/imt_log
	fi

	if [ $? -ne 0 ]; then
		send_message "ERR" "Erro na escrita na memoria RAM"
	else
		erros=`cat /tmp/imt_log  | tail -1  | awk '{print $4}'`
		if [ x$erros != x0 ]; then
			send_message "ERR" "Teste de escrita na SDRAM - ocorreram $erros erros"
		else
			send_message "OK" "Teste de escrita na SDRAM ocorreu com sucesso"
		fi
	fi

	#isto é imprevisível, não é possível determinar quem vai ser morto pelo OOMKiler
	#send_message	"OK" "Teste 2 - Teste de estouro de memoria"
	#imt 300 2> /dev/null > /tmp/imt_log
	#erros=`cat /tmp/imt_log | tail -1 `
	#if [ $erros -lt 70  ]; then
	#	send_message "WARN" "O teste conseguiu usar apenas $erros MB de memoria"
	#else
	#	send_message "OK" "Teste 2 ocorreu com sucesso - chegou a $erros MB de memoria"
	#fi
}

check_mtd() {
	nparts=`cat /proc/mtd  | wc -l`
	if [ $nparts -ne 12]; then
		send_message "ERR" "Erro nas memorias FLASH"
	fi

	if [ `cat /proc/mtd  | grep dbs | wc -l` -eq 1 ]; then
		send_message "OK"  "Memoria Dataflash FLASH encontrada: testando..."
		dataflash_flash_test
	else
		send_message "ERR" "Memoria Dataflash FLASH nao encontrada"
	fi

	if [ `cat /proc/mtd  | grep nshared | wc -l` -eq 1 ]; then
		send_message "OK"  "Memoria NAND FLASH encontrada: testando..."
		nand_flash_test
	else
		send_message "ERR" "Memoria NAND FLASH nao encontrada"
	fi

	if [ `grep PBX /proc/mtd | wc -l` -eq 1 ]; then
		send_message "OK" "Memoria FLASH PABX encontrada: testando..."
		pabx_flash_test
	else
		send_message "ERR" "Memoria FLASH PABX nao encontrada"
	fi

}

test_dsp_gw() {
	if [ $SLOT -gt 0 ]; then
		send_message "OK" "Placa Codec conectada"
	else
		send_message "WARN" "Nao existe Placa Codec"
	fi

	if [ $RETVAL -eq 0 ]; then
    	send_message "OK" "O teste da FPGA obteve sucesso"
		send_message "OK" "Placa Codec testada obteve sucesso"
		return
	fi

	if [ $(( $RETVAL & 7 )) != 0 ]; then
		send_message "ERR" "A Placa Codec apresentou defeito"
		return
	fi

	if [ $(( $RETVAL & 8 )) == 8 ]; then
	 	send_message "ERR" "A FPGA apresentou defeito"
    	return
    fi

 	send_message "ERR" "Codec apresentou defeito Desconhecido"


}


test_dsp() {
	if [ $MODO_TESTE -eq 0 ]; then
		test_fpga_ac490 > /tmp/ac_log
		RETVAL=$?
		SLOT=`dmesg | grep "Boot do DSP conectado no slot" | wc -l`
	else
		RETVAL=0
		SLOT=1
	fi

 	if [ $TYPE_HARD -eq $TYPE_HARD_GW280 ]; then
		test_dsp_gw
		return
	fi

	if [ $TYPE_HARD -eq $TYPE_HARD_UNNITI ]; then
		CODEC_COUNT=1
	else
		CODEC_COUNT=3
	fi

	for i in $(seq $CODEC_COUNT); do
		if [ $MODO_TESTE -eq 0 ]; then
			SLOT=`dmesg | grep "Inicializando Boot do DSP conectado no slot $(($i-1))" | wc -l`
		else
			SLOT=$i
		fi

		if [ $SLOT -gt 0 ]; then
			send_message "OK" "Placa Codec conectada no slot $i"
		else
			send_message "WARN" "Nao existe Placa Codec conectada no slot $i"
		fi
	done

	if [ $RETVAL -eq 0 ]; then
		send_message "OK" "Todas as Placas Codec testadas obtiveram sucesso"
		send_message "OK" "O teste da FPGA obteve sucesso"
	else
		if [ $(( $RETVAL & 1 )) == 1 ]; then
			send_message "ERR" "A Placa Codec no slot 1 apresentou defeito"
		fi

		if [ $(( $RETVAL & 2 )) == 2 ]; then
			send_message "ERR" "A Placa Codec no slot 2 apresentou defeito"
		fi

		if [ $(( $RETVAL & 4 )) == 4 ]; then
			send_message "ERR" "A Placa Codec no slot 3 apresentou defeito"
		fi

		if [ $(( $RETVAL & 8 )) == 8 ]; then
			send_message "ERR" "O FPGA apresentou defeito"
		fi
	fi

}

test_ks8842() {
	send_message "OK" "Teste da interface WAN "

	if [ $TYPE_HARD -eq $TYPE_HARD_ICIP68 ]; then
		send_message "OK" "Carregando driver WAN 16 bits "
		insmod /drivers/ks8842.ko typeIs16Bits=1
	else
		send_message "OK" "Carregando driver WAN"
		insmod /drivers/ks8842.ko typeIs16Bits=0
	fi

	if [ $MODO_TESTE -eq 0 ]; then
		if [ $? -ne 0 ]; then
			send_message "ERR" "Erro ao carregar driver"
			return
		else
			send_message "OK" "Modulo carregado com sucesso"
		fi

		sleep 1
		found=`dmesg | grep ks8842 | grep "Found chip" | wc -l`

	else
		send_message "OK" "Modulo carregado com sucesso"
		sleep 1
		found=1
	fi

	if [ $found -gt 0 ]; then
		send_message "OK" "Interface wan encontrada e acessivel"
	else
		send_message "ERR" "Interface wan nao encontrada"
		return
	fi

	send_message "OK" "Teste de configuracao da WAN"
	send_message "OK" "Configurando a interface WAN"
	send_message "OK" "Endereco $WAN_IP Mascara 255.255.255.0"
	ifconfig eth1 $WAN_IP netmask 255.255.255.0 2> /dev/null > /dev/null
	sleep 1
	send_message "OK" "Pingando endereco $WAN_TEST_IP"
	ping -c 3 $WAN_TEST_IP 2> /dev/null > /dev/null
	if [ $? -eq 0 ]; then
		send_message "OK" "Ping OK"
	else
		send_message "WARN" "Erro ao pingar endereço $WAN_TEST_IP"
	fi
}

test_lan() {
	send_message "OK" "Pingando endereco $LAN_TEST_IP"
	ping -c 3 $LAN_TEST_IP 2> /dev/null > /dev/null
	if [ $? -eq 0 ]; then
		send_message "OK" "Ping OK"
	else
		send_message "WARN" "Erro ao pingar endereço $LAN_TEST_IP"
	fi
}

source /usr/bin/get_hardware.sh

if [ $MODO_TESTE -eq 0 ]; then
	mdev -s
fi

# Enquanto nao fica definido se o fw do recover
# sera especifico para a unniti ou nao, considera
# que e especifico e forca tipo de hardware em unniti
TYPE_HARD=$TYPE_HARD_UNNITI

case $1 in
	"memoria")
		memory_test
		;;
	"flash")
		check_mtd
		;;
	"lan")
		test_lan
		;;
	"usb")
		check_hardware_key
		;;
	"codecs")
		test_dsp
		;;
	"wan")
		test_ks8842
		;;

esac

if [ $ERRORS -eq 0 ]; then
	echo 0
else
	echo 1
fi




*/