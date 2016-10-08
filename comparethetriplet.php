<?php

$handle = fopen ("php://stdin","r");
fscanf($handle,"%d %d %d",$a0,$a1,$a2);
fscanf($handle,"%d %d %d",$b0,$b1,$b2);
compare($a0, $b0)." ".compare($a1, $b1)." ".compare($a2, $b2);

function compare($a, $b){
    if($a < $b){
        echo $b." ";
    }else if($a > $b){
        echo $a." ";
    }
}
?>
