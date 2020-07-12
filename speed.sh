#!/bin/bash -eu
types=("bfs" "iddfs" "A*")
for ((i=2; i<=20; i++))
do
for t in ${types[@]}
do
cat ./testData/$i.txt | ./main solve $t >> ./speed_result/$t\_$i.txt
cat ./speed_result/$t\_$i.txt
done
done