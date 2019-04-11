
all: winexec.exe

winexec.exe: *.go
	GOOS=windows GOARCH=amd64 go build -o winexec.exe .

test: winexec.exe
	./winexec.exe "c:\\windows\\system32\\notepad.exe"
