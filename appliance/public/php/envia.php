<html>
<head>
<title>Teste</title>
</head>
<body>
<?
// Diretorio onde o arquivo vai ser salvo
$target_path = "/tmp/";
// Colocar aqui o nome do seu executavel :)
$executavel = "/usr/bin/recover.sh";

// Debug
$debug=1;

$target_path = $target_path . basename( $_FILES['filename']['name']);
// Pega o valor da checkbox
// on => marcado
// vazio => desmarcado
if(isset($_REQUEST['ctz'])) {
	$ctz = $_REQUEST['ctz'];
} else {
	$ctz = "";
}
$out = "";

// Se conseguir mover, moveu..
if(move_uploaded_file($_FILES['filename']['tmp_name'], $target_path)) {
	echo "O arquivo " . basename( $_FILES['filename']['name']) . " foi enviado com sucesso e a flag: ". $ctz. "<br>";
	exec($executavel.' '. $target_path . ' ' . $ctz,$out);
	if($out[count($out)-1] != 0){
		echo "<resposta>ERRO</resposta>";
	} else {
		echo"<resposta>OK</resposta>";
	}
	
	if($debug == 1){
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
	}
} else {
    	echo "Erro ao enviar o arquivo.";
}
?>
</body>
</html>
