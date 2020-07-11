#!/bin/bash -eu
read -p "name: " name
timeout 600 ./main measure $1 $2 >> ./data/$name.txt
types=("bfs" "iddfs" "A*")
for t in ${types[@]}
do
mv cpu_$t.prof ./results/cpu_$t\_$name.prof
mv mem_$t.prof ./results/mem_$t\_$name.prof
go tool pprof -png ./results/cpu_$t\_$name.prof
mv profile001.png ./image/cpu_$t\_$name.png
go tool pprof -png ./results/mem_$t\_$name.prof
mv profile001.png ./image/mem_$t\_$name.png
done
cat ./data/$name.txt