
//var MODO_ = "dev";
//var MODO_ = "prod";

var TYPE_APPLIANCE  = "appliance";
var TYPE_UNNITI = "unniti";


/*if ( MODO_ == "dev")
	var file_hard_conf = "/home/intelbras/projetos-go/src/github.com/ma022800/goma/goma/appliance/public/static/hard.conf";
else if ( MODO_ == "prod")*/
var file_hard_conf = "hard.conf";

var fezselftest = false;
var numeroSelfTest = 0;
var erro=0;

/* Array com os parametros que cada script devera receber */
var arraySelfTest = [];
var type_hardware_functions;

selfTest_init()

function validar(){
	document.getElementById('msg').innerHTML = "";
	if(!validaArquivo()){
		return false;
	}

	document.getElementById('recovery').submit();
}

function validaArquivo(){
	var arq = document.getElementById('filename').value;
	if(arq.length < 1){
		document.getElementById('msg').innerHTML = "Favor selecionar um arquivo antes de enviar";
		return false;
	}
	return true;
}

// Valida o Mac Address
function validaMac(n){
	document.getElementById('msg').innerHTML = "";
	var mac1 = document.getElementById('mac'+n).value;
	
	if(mac1.length != 17){
		document.getElementById('msg').innerHTML = "Mac "+n+" Est&aacute; incorreto";
		return false;
	}
	mac1 = mac1.split(":");
	if(mac1.length != 6){ // Tamanho incorreto
		document.getElementById('msg').innerHTML = "Mac "+n+" Est&aacute; incorreto";
		return false;
	}
	var ok=true;
	for(var i=0;i<mac1.length;i++){
		var inteiro = parseInt(mac1[i],16);
		var digito1 = parseInt(mac1[i][0],16);
		var digito2 = parseInt(mac1[i][1],16);
		
		if( isNaN(inteiro) || isNaN(digito1) || isNaN(digito2)) { // Verifica se alguma conversao deu problema
			ok = false;
			break;
		}
		
		if(inteiro >= 0 && inteiro < 256){ // Verifica se esta no range
			ok = true;
		} else {
			ok = false;
			break;
		}
	}
	if(ok == false){
		document.getElementById('msg').innerHTML = "Mac "+n+" Est&aacute; incorreto";
		return false;
	}
	return true;
}

$(document).ready(function(){
	$('div .hide').hide();
	$('#btnself').click(function(){
		$('#btnself').attr('disabled','disabled');
		$('#saida1').empty();
		//$('#saida1').append("<font color='#2e802e' size='4'>INFO Teste de HARDWARE</font><br />");
		//$('#saida1').append("<font color='#0066FF' size='3'>" + get_msg_harware() + "</font><br />");
		/*$.ajax({
			//url: 'http://localhost:8080/testes',
			url: '/testes',
			method: "GET",
			success: function(data) {
				$("#response").html(data);
				$('#saida1').append(data);
		
			},
		});*/
		//$('#saida1').append("<font color='#2e802e' size='4'>Teste Hardware la vai</font><br />");
		
	//	$('#saida1').append("<font color='#2e802e' size='4'>Inicio dos Testes em</font><br />");
		//get_date();
		$('#btnself').hide();
		config_test();
		selfTest();
		//sendParamJsToGo(file_hard_conf)
		//sendParamGoToJs(file_hard_conf)
		
	});
	
	$('#skipself').click(function(){
		if($(this).attr('checked') != undefined){
			$('#envia_firmware').slideDown();
		} else {
			if(!fezselftest){
				$('#envia_firmware').slideUp();
			}
		}
	});
	
	$('#recovery').submit(function(e){
		$('#btnEnviar').attr('disabled','disabled');
		leFrame();
	});
	
	$('#busca_mac').click(function(){
		$('#busca_mac').attr('disabled','disabled');
		$.post('mac.php',{busca:'1'},function(data){
			$('#swap').html(data);
			var resposta = $('#swap resposta').html();
			if(resposta == "OK" || resposta == "GET"){
				var mac1 = $('#swap mac1').html();
				var mac2 = $('#swap mac2').html();
				var mac3 = $('#swap mac3').html();
				var mac4 = $('#swap mac4').html();

				$('#mac_log').append("<label>MAC LAN: </label>" + mac1 + "<br>");
				$('#mac_log').append("<label>MAC WAN: </label>" + mac2 + "<br>");
				$('#mac_log').append("<label>MAC WAN: </label>" + mac3 + "<br>");
				$('#mac_log').append("<label>MAC WAN: </label>" + mac4 + "<br>");

				if (resposta == "OK")
					$('#mac_log').append("<br><h3><font color='#006600'>MAC lido com sucesso.<font></h3>");
				else
					$('#mac_log').append("<br><h3><font color='#006600'>MAC buscado da rede com sucesso.<font></h3>");

				$('#mac_log').slideDown();
				$('#imprimir_mac').slideDown();
			} else {
				$('#gm_grava').slideDown();
				show_mac_eth1_field();
			}
		});
	});
	
	$('#imprimir_mac').click(function(){
		$('#imprimir_mac').attr('disabled','disabled');
		$.post('imprimir.php',{busca:'1'},function(data){
			$('#swap').html(data);
			var resposta = $('#swap resposta').html();
			if(resposta == "OK"){
				$('#imprimir_log').append("<br><h3><font color='#006600'>MAC impresso com sucesso<font></h3>");
				$('#imprimir_log').slideDown();
			} else {
				$('#imprimir_log').append("<br><h3><font color='#006600'>Erro ao imprimir MAC<font></h3>");
				$('#imprimir_log').slideDown();
			}
		});
	});
	
	$('#frm_mac').submit(function(e){
		e.preventDefault();
		$('#msg').html("");
		if(!validaMac(1)){
			$('#msg').html("Favor preencher o MAC ETH0 corretamente, ex: 00:1a:3f:01:01:22");
			return;
		}
	
		if(!validaMac(2)){
			$('#msg').html("Favor preencher o MAC ETH1 corretamente, ex: 00:1a:3f:01:01:23");
			return;
		}

		if(!validaMac(3)){
			$('#msg').html("Favor preencher o MAC ETH2 corretamente, ex: 00:1a:3f:01:01:23");
			return;
		}

		if(!validaMac(4)){
			$('#msg').html("Favor preencher o MAC ETH3 corretamente, ex: 00:1a:3f:01:01:23");
			return;
		}

		$.post('mac.php',$(this).serialize(),function(data){
			$('#swap').html(data);		

			var resposta = $('#swap resposta').html();
			if(resposta == "OK"){
				$('#saida3').html("Mac gravado com sucesso!");
				$('#imprimir_mac').slideDown();
			}
		});
	});
});

var globLoop = 0;
var MAX_LOOP = 5;
function leFrame(){
	var leu = false;
	var erro = false;
	globLoop++;
	var resposta = $('#enviar').contents().find('resposta').html();
	if(resposta == null){
		setTimeout('leFrame()',500);
	} else {
		if(resposta == "OK"){
			$('#grava_mac').slideDown();
			$('#btnEnviar').attr('disabled','disabled');
		} else {
			$('#msg').html('Falha no envio do arquivo de Firmware');
		}
		$('#saida2').empty();
		$('#enviar').contents().find('valor').each(function(){
			$('#saida2').append($(this).html()+'<br />');
		});
	}
}

function selfTest(){
	param=arraySelfTest[numeroSelfTest];
	numeroSelfTest++;
	if(param==undefined)
	{
		
		/* finalizou a execucao de todos OS scripts entao verifica se tudo foi OK */
		if(erro == 0)
		{
			/* abre o envio de firmware */
			fezselftest = true;
			$('#grava_mac').slideDown();
			 $('#gm_grava').slideDown();
			 $('#frm_mac').slideDown();
			show_mac_eth1_field();
			show_mac_eth2_field();
			show_mac_eth3_field();
			//print_log("selfTest_FIM: DOWN:" + erro);
			erro=0;
			numeroSelfTest=0;
			return;
	
		}
		//print_log("selfTest_FIM: NumTestes:"+ numeroSelfTest +"Erros:" + erro);
		//$('#envia_firmware').hide();
		erro=0;
		numeroSelfTest=0;
        return;
	 			
		/* finalizou a execucao de todos OS scripts entao verifica se tudo foi OK * /
		if(erro != 0)
		{
			/* se estava aberto fecha o envio * /
			$('#envia_firmware').slideUp();
			print_log("selfTest_FIM: UP Erros:" + erro);
	
		} else {
			/* abre o envio de firmware * /
			fezselftest = true;
			//$('#envia_firmware').slideDown();
			$('#envia_firmware').hide();
			print_log("selfTest_FIM: DOWN:" + erro);
	
		}
		//print_log("selfTest_FIM: NumTestes:"+ numeroSelfTest +"Erros:" + erro);
		erro=0;
		numeroSelfTest=0;
		return;
		*/

	}

	/*if ( MODO_ == "dev")
	{
		selfTest();
		return ;
	}*/

	set_wait();
    $.ajax({
		url: '/selftest', //name function
		method: "POST",
		data: {param} ,
		success: function(data) {
			$('#swap').empty();
 		    $('#swap').html(data);
 		    var resposta = $('#swap resposta').html();
 		    if(resposta != "0")
 			   erro++;

			$('#swap valor').each(function(){
				$('#saida1').append($(this).html() + "<br />");
			});
			selfTest();
			//$('#saida1').append("<font color='#2e802e' size='4'>Inicio dos Testes</font><br />");

		},
		

	});
	//$('#saida1').append("<font color='#2e802e' size='4'>Inicio dos Testes FIM</font><br />");


	/* Para ler os parametros irao chegar em x * /
	$.post('self-test.php',{x:param}, function(data){
		$('#swap').empty();
		$('#swap').html(data);
		var resposta = $('#swap resposta').html();
		if(resposta != "0")
			erro++;

		$('#swap valor').each(function(){
			$('#saida1').append($(this).html() + "<br />");
		});
		selfTest();
	});
	*/
}

/*
 * Alteravcoes da Nova versao de Auto Teste
 */

/*
 * Descobre qual tipo de hardware e monta os tipos de 
 * testes que serão feitos
 */
function selfTest_init()
{
	$.ajax({
		url: '/iniselftest', //name function
		method: "POST",
		//data: {nomeArquivo} ,
		success: function(data) {
			$('#swap').empty();
 		    $('#swap').html(data);
 		    var resposta = $('#swap resposta').html();
 		    if(resposta != "0")
 			   erro++;

			read_file_typeHardware(file_hard_conf);
			//$('#saida1').append("<font color='#2e802e' size='4'>Inicio dos Testes as</font><br />");
			//get_hour(); 
	
		},
		

	});
	//$('#saida1').append("<font color='#2e802e' size='4'>Inicio dos Testes FIM</font><br />");
	//$('#saida1').append("<font color='#2e802e' size='4'>Abre Arquivo</font><br />");
 
	/*	
	$.post('php/self-test.php', function(data)
 	{
 		$('#swap').empty();
 		$('#swap').html(data);
 		var resposta = $('#swap resposta').html();
 		if(resposta != "0")
 			erro++;

		read_file_typeHardware(file_hard_conf);
 	});*/
}

function read_file_typeHardware(nomeArquivo)
{
 	/*$.post('read_file.php', {x:nomeArquivo}, function(data){
 		set_typeHardware(data);
 		set_typeHardware("unniti"); //Enquanto nao se decide se fw do recover vai ser exclusivo da unniti...
	 });*/
	 

	$.ajax({
		url: '/readfile', //name function
		method: "GET",
		data: {nomeArquivo} ,
		success: function(data) {
			$("#response").html(data);
 		    var resposta = $('#response resposta').html();
 		    if(resposta != "0")
				$('#saida1').append(data);

			set_typeHardware(data);
//			set_typeHardware("unniti"); //Enquanto nao se decide se fw do recover vai ser exclusivo da unniti...
		},
		
	});
	//$('#saida1').append("<font color='#2e802e' size='4'>Abre Arquivo</font><br />");
	//print_log("erro leitura");
}

/*
*/
function get_date()
{
 	$.ajax({
		url: '/date', //name function
		method: "GET",
		success: function(data) {
			$("#response").html(data);
			$('#saida1').append(data);
			//return data
			//set_typeHardware(data);
//			set_typeHardware("unniti"); //Enquanto nao se decide se fw do recover vai ser exclusivo da unniti...
		},
	});
	//$('#saida1').append("<font color='#2e802e' size='4'>Le Hora</font><br />");

}

function set_wait()
{
 	$.ajax({
		url: '/wait', //name function
		method: "GET",
		success: function(data) {
			$('#saida1').append(data);
		},
	});

}


function get_msg_harware()
{
 	if(is_hardware_UnniTI())
		 return "Hardware UnniTI";
    if(is_hardware_Appliance())
 		return "Hardware Appliance";	 
}

function config_test()
{
 	if(is_hardware_UnniTI())
		set_test_unniti();
	else if(is_hardware_Appliance())
 		set_test_appliance();	 
}

function set_typeHardware(new_type)
{
	type_hardware_functions=new_type;
	//$('#saida1').append("<font color='#2e802e' size='4'>set_typeHardware type_hardware_functions</font><br />");
	//$('#saida1').append(type_hardware_functions);
	//print_log("erro leitura");
}

function get_typeHardware()
{
	return type_hardware_functions;
}

function set_test_appliance()
{
	//arraySelfTest.push('memoria');
	//arraySelfTest.push('flash');
	//arraySelfTest.push('codecs');
	arraySelfTest.push('usb');
//T	arraySelfTest.push('memoria');
//T	arraySelfTest.push('eths');
	//arraySelfTest.push('eth1');
	//arraySelfTest.push('eth2');
	//arraySelfTest.push('eth3');
    //arraySelfTest.push('wan');
//T	arraySelfTest.push('ssd');

}

function set_test_unniti()
{
	set_test_appliance();
}


function is_hardware_Appliance()
{
	var type_hard = get_typeHardware();

	if(type_hard.indexOf(TYPE_APPLIANCE) > -1 )
		return true;

	return false;
}

function is_hardware_UnniTI()
{
	var type_hard = get_typeHardware();

	if(type_hard.indexOf(TYPE_UNNITI) > -1 )
		return true;

	return false;
}

function show_mac_eth1_field()
{
	$('#mac_eth1').slideDown();
}

function show_mac_eth2_field()
{
	$('#mac_eth2').slideDown();
}

function show_mac_eth3_field()
{
	$('#mac_eth3').slideDown();
}



function print_log(msg_log)
{
	//if ( MODO_ == "prod")
	//	return ;

	alert(msg_log);
}


//Interfaces em GO 

//Passando parametros do JS para funcoes em GO
function sendParamJsToGo(aData)
{
    $.ajax({
		//url: 'http://localhost:8080/testes',
		url: '/rxdata', //name function
		method: "PUT",
		data: {aData} ,
		success: function(data) {
			$("#response").html(data);
			$('#saida1').append(data);
			
		},
	});
	$('#saida1').append("<font color='#2e802e' size='4'>sendParamJsToGo</font><br />");

}

//Passando parametros do GO para JS 
function sendParamGoToJs(aData)
{
    $.ajax({
		//url: 'http://localhost:8080/testes',
		url: '/txdata', //name function
		method: "GET",
		data: {aData} ,
		success: function(data) {
			$("#response").html(data);
			$('#saida1').append(data);
			
		},
	});
	$('#saida1').append("<font color='#2e802e' size='4'>sendParamGoToJS</font><br />");



}

