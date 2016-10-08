<?php

$handle = fopen ("php://stdin","r");
fscanf($handle,"%d",$n);
$arr_temp = fgets($handle);
$arr = explode(" ",$arr_temp);
array_walk($arr,'intval');
$numberOfPositives = 0;
$numberOfNegatives = 0;
$numberOfZeros = 0;
for($i=0; $i < $n; $i++){
    if($arr[$i] > 0){
        $numberOfPositives++;
    }
    else if($arr[$i] == 0){
        $numberOfZeros++;
    }
    else{
        $numberOfNegatives++;
    }
}
echo number_format(floatval($numberOfPositives/$n),4)."\n";
echo number_format(floatval($numberOfNegatives/$n),4)."\n";
echo number_format(floatval($numberOfZeros/$n),4)."\n";
?>