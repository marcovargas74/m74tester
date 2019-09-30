<?php 
if(isset($_REQUEST['busca'])){
	/*
	 * Seta aqui o executavel que busca os MACs
	 * resposta esperada:
	 * XX:XX:XX:XX:XX:XX (somente em caso OK)
	 * YY:YY:YY:YY:YY:YY (somente em caso OK)
	 * 0 (OK)
	 * ou
	 * 1 (NOK)
	 */
	$executavel = "/usr/bin/read_mac.sh";
	exec($executavel,$out);
	if($out[count($out)-1] != 0){
		$executavel = "/usr/bin/get_mac.sh";
		exec($executavel,$out);
		if($out[count($out)-1] != 0) {
			echo "<resposta>ERRO</resposta>";
		} else {
			echo "<mac1>$out[0]</mac1>";
			echo "<mac2>$out[1]</mac2>";
			echo"<resposta>GET</resposta>";
		}
		exit;	
	} else {
		/* ASSUMINDO QUE A PRIMEIRA LINHA é o mac1, a segunda é o mac2 */
		echo "<mac1>$out[0]</mac1>";
		echo "<mac2>$out[1]</mac2>";
		echo"<resposta>OK</resposta>";
	}
} else { // fw - gravando
	
	$mac1 = $_REQUEST['mac1'];
	$mac2 = $_REQUEST['mac2'];
	/*
	 * Seta aqui o executavel que recebe como parametro mac1 e mac2
	 * resposta esperada:
	 * 0 ou 1;
	 */
	$executavel = "/usr/bin/mac_only_write.sh $mac1 $mac2";
	exec($executavel,$out);
	if($out[count($out)-1] != 0){
		echo "<resposta>ERRO</resposta>";
		exit;	
	} else {
		echo"<resposta>OK</resposta>";
	}
}
?>
