#!/bin/sh
#Script usado para Compila aplicação para ARM
#
#

#Dados para compilar para o Apliance
APP_NAME=appliance
DIR_WORK_APP=~/projetos-go/src/github.com/ma022800/goma
DIR_WORK_PUBLIC=$DIR_WORK_APP/goma/appliance/public/
#DIR_FIRMWARE=~/UNNITI_GIT/firmware/packetsAPP_bin/release_voip
IP_REMOTE=172.31.1.123


echo "--------------------------INICIO------------------------------------"
echo "1. INICIANDO a Compilacao do app!!!!"

#Compila pra ARM
#env GOOS=linux GOARCH=arm go build -o $APP_NAME

env GOOS=linux go build -o $APP_NAME

if [ -e $APP_NAME ]; then
  echo "3. OK APP Compilado!!!!"
else
  echo "4. NOK Problema ao compilar APP!!!!"
  exit
fi	

ls -l -h 


#copia a aplicacao para pasta do gerador de firmware
#

#TESTE copia firmware para testar na Unniti
#diretorio usado para gerar as Versoes de Firmware da Versao desenvolvimento
#echo ""
##echo ""
#cp -f $DIR_WORK_APP/$APP_NAME  $DIR_FIRMWARE
#echo "Copia o arquivo binario para a pasta do Firmware!!!"
#echo "--------------------------INICIO------------------------------------"
#cd $DIR_WORK_APP
#se Nao existir apaga o antigo 
#if [ ! -e fpga.ko ]; then
# echo "ARQUIVO DO Drive FPGA nao Existe!!! "
#fi	

echo "Copia o arquivo binario para o REMOTE $IP_REMOTE!!!!"
#ls -l

scp $DIR_WORK_APP/$APP_NAME  iap@$IP_REMOTE:/home/iap/appliance

echo "Arquivo copiado para para $IP_REMOTE!!!"

echo "Copia as paginas para o REMOTE $IP_REMOTE!!!!"
scp -r $DIR_WORK_PUBLIC  iap@$IP_REMOTE:/home/iap/appliance

echo "PAginas copiado para para $IP_REMOTE!!!"

echo "--------------------------FIM------------------------------------"
echo ""

cd ..


