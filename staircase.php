<?php

$handle = fopen ("php://stdin","r");
fscanf($handle,"%d",$n);
for($i=1; $i<=$n; $i++){
    for($j=$n; $j>0; $j--){
        if($j <= $i){
            echo "#";
        }else{
            echo " ";
        }
    }
    echo "\n";
}
?>