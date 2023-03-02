MAIN_FILE = main.go
OUT_FILE = tmachine.exe

run:
	go build -o ${OUT_FILE} ${MAIN_FILE}
	./${OUT_FILE} busybeaver2