<?php 
/*
 * fw - Executa o self-test
 * retorno do self test:
 *   'qualquer coisa'
 *   e na ULTIMA linha 0 ou 1
 */

//$MODO_ = 'dev';
$MODO_ = 'prod';

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

?>
