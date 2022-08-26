##Примеры использования

go run main.go -c "twenty" .\in.txt

go run main.go -v "twenty" .\in.txt

go run main.go -v -F "twenty" .\in.txt

go run main.go -C 2 "twenty" .\in.txt

go run main.go -A 2 -B 1 "twenty" .\in.txt

go run main.go "twenty" .\in.txt

go run main.go "[0-9]+" .\in.txt

go run main.go -n 'href=.[\w[:punct:]]+\"' .\in.txt

go run main.go -n -F 'href=.[\w[:punct:]]+\"' .\in.txt


