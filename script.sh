#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'


list="list.txt"
main_program="cmd/app/main.go"
#add "input" and "expected output file.txt"
tests=("HELLO" "test1.txt"
       "hello" "test2.txt" 
       "HeLlo HuMaN" "test3.txt"
       "1Hello 2There" "test4.txt"
       "Hello\nThere" "test5.txt"
       "Hello\n\nThere" "test6.txt"
		   "abcdefghijklmnopqrstuvwxyz" "test17.txt"
		   "hello There 1 to 2!" "test7.txt"
		   "{Hello & There #}" "test8.txt"
		   "MaD3IrA&LiSboN" "test9.txt"
		   "1a\"#FdwHywR&/()=" "test10.txt"
		   "{|}~" "test11.txt"
		   "[\\]^_ 'a" "test12.txt"
		   "RGB" "test13.txt"
		   ":;<=>?@" "test14.txt"
		   "\\!\" #$%&'()*+,-./" "test15.txt"
		   "ABCDEFGHIJKLMNOPQRSTUVWXYZ" "test16.txt"
       "\n\n\n" "test18.txt"
       )

count=0
for value in "${tests[@]}"; do
((count++))

 if [ "$count" -eq 2 ]; then
  if diff "temp.txt" "./testcases/$value" 2>&1 > /dev/null || diff "temp.txt" "./testcases/error.txt" 2>&1 > /dev/null; then  # 2>&1 > /dev/null hides error outputs
    printf "${GREEN}---$value Bim!---\n${NC}"
  else
    printf "${RED}---$value Oh Snap!---\n${NC}"
  fi
  count=0
elif [ "$count" -eq 1 ]; then
  go run "$main_program" "$value" > temp.txt
fi


done

rm temp.txt
