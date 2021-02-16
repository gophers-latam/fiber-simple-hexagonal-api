all: win32 win64 linux64

win32:
	GOOS=windows GOARCH=386 go build -o api_server_x32.exe

win64:
	GOOS=windows GOARCH=amd64 go build -o api_server_x64.exe

linux64:
	GOOS=linux GOARCH=amd64 go build -o api_server

clean:
	rm api_server_x32.exe && rm api_server_x64.exe && rm api_server