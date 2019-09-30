<?php 

$executavel = "/usr/bin/print_mac.sh";
exec($executavel,$out);
if($out[count($out)-1] != 0){
	echo "<resposta>ERRO</resposta>";
	exit;	
} else {
	echo"<resposta>OK</resposta>";
}
?>
