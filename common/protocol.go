package common

const (
	AskForNameMessage = "Por favor,introduce un nombre para jugar: "
	WelcomeMessage    = "Bienvenidos a  FIUBADOS: "
	ObjectiveMessage  = "E̳l̳ o̳b̳j̳e̳t̳i̳v̳o̳ d̳e̳ e̳s̳t̳e̳ j̳u̳e̳g̳o̳ e̳s̳ i̳n̳t̳r̳o̳d̳u̳c̳i̳r̳ a̳ l̳o̳s̳ a̳l̳u̳m̳n̳o̳s̳ q̳u̳e̳ i̳n̳g̳r̳e̳s̳a̳n̳ a̳ l̳a̳ f̳a̳c̳u̳l̳t̳a̳d̳ e̳n̳ d̳i̳v̳e̳r̳s̳a̳s̳ c̳u̳e̳s̳t̳i̳o̳n̳e̳s̳ a̳d̳m̳i̳n̳i̳s̳t̳r̳a̳t̳i̳v̳a̳s̳ d̳e̳ s̳u̳ f̳u̳n̳c̳i̳o̳n̳a̳m̳i̳e̳n̳t̳o̳,̳ s̳i̳e̳n̳d̳o̳ e̳n̳ e̳l̳ f̳u̳t̳u̳r̳o̳ m̳u̳y̳ ú̳t̳i̳l̳ p̳a̳r̳a̳ e̳l̳ d̳e̳s̳a̳r̳r̳o̳l̳l̳o̳ d̳e̳ s̳u̳ c̳a̳r̳r̳e̳r̳a̳ p̳r̳o̳f̳e̳s̳i̳o̳n̳a̳l̳.̳"
	MainMenuOptions   = "(1) Jugar  (2) Ayuda  (3) Salir"

	HelpMessage     = "Este juego consiste en partidas 1vs1 en las que dos jugadores responden varias preguntas de opción múltiple. Cada jugador contestará el número de la opción que considere correcta en cada pregunta. El jugador que responda de forma correcta aumenta su puntuación. Si ambos jugadores responden correctamente una pregunta, el primero que haya respondido se lleva puntos adicionales. Asímismo, si solo uno de los jugadores responde correctamente, se lleva puntos adicionales. El jugador que responda incorrectamente no suma ningun punto."
	HelpMenuOptions = "(1) Volver al Menu principal"

	CloseConnectionCommand = "STOP"
	Success                = "OK"
	ReadyToPlay            = "LISTO"

	DisconnectAndExitMessage = "SERVER DISCONNECTED. Client exiting..."
	ExitMessage              = "Client exiting..."

	SearchingMatchMessage = "A̳r̳m̳a̳n̳d̳o̳ u̳n̳a̳ p̳a̳r̳t̳i̳d̳a̳:̳ B̳u̳s̳c̳a̳n̳d̳o̳ j̳u̳g̳a̳d̳o̳r̳e̳s̳.̳.̳.̳"

	PlayOption = "1"
	HelpOption = "2"
	ExitOption = "3"

	MatchingPlayerMessage = "Está jugando contra el jugador: "
	ReadyToPlayMessage    = ". Introduzca LISTO cuando esté listo para jugar"

	QuestionMessage = "Pregunta "
	ColonMessage    = ": "
	FirstOption     = " (1)"
	SecondOption    = " (2)"
	ThirdOption     = " (3)"

	CorrectAnswerMessage               = "Respuesta correcta! "
	WasFirstToAnswerMessage            = "Respondiste de primero! Recibes 3 puntos adicionales. "
	WasSecondToAnswerMessage           = "Respondiste de segundo! Recibes solo 1 punto. "
	OpponentAnsweredIncorrectlyMessage = "Tu contrincante respondió incorrectamente, te llevas 3 puntos adicionales. "
	IncorrectAnswerMessage             = "Respuesta incorrecta! "
	WhichWasCorrectAnswerMessage       = "La respuesta correcta era: ("

	PlayerMessage                  = "Jugador "
	WinnerMessage                  = " ha ganado! ¡Gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"
	TieMessage                     = "¡Juego empatado! ¡Gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"
	OtherPlayerDisconnectedMessage = "El otro jugador se desconectó ¡Ganaste, gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"

	ColorCyan  = "\033[96m"
	ColorReset = "\033[0m"
	ColorGreen = "\033[92m"

	ServerArrow = `
🤖->:`
	AsciAskForNameMessage = `

    █▀█ █▀█ █▀█   █▀▀ ▄▀█ █░█ █▀█ █▀█ ░   █ █▄░█ ▀█▀ █▀█ █▀█ █▀▄ █░█ █▀▀ █▀▀   █░█ █▄░█   █▄░█ █▀█ █▀▄▀█ █▄▄ █▀█ █▀▀
    █▀▀ █▄█ █▀▄   █▀░ █▀█ ▀▄▀ █▄█ █▀▄ █   █ █░▀█ ░█░ █▀▄ █▄█ █▄▀ █▄█ █▄▄ ██▄   █▄█ █░▀█   █░▀█ █▄█ █░▀░█ █▄█ █▀▄ ██▄

    █▀█ ▄▀█ █▀█ ▄▀█   ░░█ █░█ █▀▀ ▄▀█ █▀█ ▀
    █▀▀ █▀█ █▀▄ █▀█   █▄█ █▄█ █▄█ █▀█ █▀▄ ▄
        `
	AsciWelcomeMessage = `
		
				██████╗░██╗███████╗███╗░░██╗██╗░░░██╗███████╗███╗░░██╗██╗██████╗░░█████╗░░██████╗  ░█████╗░  
				██╔══██╗██║██╔════╝████╗░██║██║░░░██║██╔════╝████╗░██║██║██╔══██╗██╔══██╗██╔════╝  ██╔══██╗  
				██████╦╝██║█████╗░░██╔██╗██║╚██╗░██╔╝█████╗░░██╔██╗██║██║██║░░██║██║░░██║╚█████╗░  ███████║  
				██╔══██╗██║██╔══╝░░██║╚████║░╚████╔╝░██╔══╝░░██║╚████║██║██║░░██║██║░░██║░╚═══██╗  ██╔══██║  
				██████╦╝██║███████╗██║░╚███║░░╚██╔╝░░███████╗██║░╚███║██║██████╔╝╚█████╔╝██████╔╝  ██║░░██║  
				╚═════╝░╚═╝╚══════╝╚═╝░░╚══╝░░░╚═╝░░░╚══════╝╚═╝░░╚══╝╚═╝╚═════╝░░╚════╝░╚═════╝░  ╚═╝░░╚═╝  

				███████╗██╗██╗░░░██╗██████╗░░█████╗░██████╗░░█████╗░░██████╗██╗
				██╔════╝██║██║░░░██║██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝╚═╝
				█████╗░░██║██║░░░██║██████╦╝███████║██║░░██║██║░░██║╚█████╗░░░░
				██╔══╝░░██║██║░░░██║██╔══██╗██╔══██║██║░░██║██║░░██║░╚═══██╗░░░
				██║░░░░░██║╚██████╔╝██████╦╝██║░░██║██████╔╝╚█████╔╝██████╔╝██╗
				╚═╝░░░░░╚═╝░╚═════╝░╚═════╝░╚═╝░░╚═╝╚═════╝░░╚════╝░╚═════╝░╚═╝
		`
	AsciMainMenuOptions = `
	
	▄▀ ▄█─ ▀▄ 　 ───░█ █──█ █▀▀▀ █▀▀█ █▀▀█ 　 　 ▄▀ █▀█ ▀▄ 　 ─█▀▀█ █──█ █──█ █▀▀▄ █▀▀█ 　 　 ▄▀ █▀▀█ ▀▄  ░█▀▀▀█ █▀▀█ █── ─▀─ █▀▀█ 
	█─ ─█─ ─█ 　 ─▄─░█ █──█ █─▀█ █▄▄█ █▄▄▀ 　 　 █─ ─▄▀ ─█ 　 ░█▄▄█ █▄▄█ █──█ █──█ █▄▄█ 　 　 █─ ──▀▄ ─█  ─▀▀▀▄▄ █▄▄█ █── ▀█▀ █▄▄▀
	▀▄ ▄█▄ ▄▀ 　 ░█▄▄█ ─▀▀▀ ▀▀▀▀ ▀──▀ ▀─▀▀ 　 　 ▀▄ █▄▄ ▄▀ 　 ░█─░█ ▄▄▄█ ─▀▀▀ ▀▀▀─ ▀──▀ 　 　 ▀▄ █▄▄█ ▄▀  ░█▄▄▄█ ▀──▀ ▀▀▀ ▀▀▀ ▀─▀▀	
				
		`
	AsciHelpMessage = `	E̳s̳t̳e̳ j̳u̳e̳g̳o̳ c̳o̳n̳s̳i̳s̳t̳e̳ e̳n̳ p̳a̳r̳t̳i̳d̳a̳s̳ 1̳v̳s̳1̳ e̳n̳ l̳a̳s̳ q̳u̳e̳ d̳o̳s̳ j̳u̳g̳a̳d̳o̳r̳e̳s̳ r̳e̳s̳p̳o̳n̳d̳e̳n̳ v̳a̳r̳i̳a̳s̳ p̳r̳e̳g̳u̳n̳t̳a̳s̳ d̳e̳ o̳p̳c̳i̳ó̳n̳ m̳ú̳l̳t̳i̳p̳l̳e̳.̳ C̳a̳d̳a̳ j̳u̳g̳a̳d̳o̳r̳ c̳o̳n̳t̳e̳s̳t̳a̳r̳á̳ e̳l̳ n̳ú̳m̳e̳r̳o̳ d̳e̳ l̳a̳ o̳p̳c̳i̳ó̳n̳ q̳u̳e̳ c̳o̳n̳s̳i̳d̳e̳r̳e̳ c̳o̳r̳r̳e̳c̳t̳a̳ e̳n̳ c̳a̳d̳a̳ p̳r̳e̳g̳u̳n̳t̳a̳.̳E̳l̳ j̳u̳g̳a̳d̳o̳r̳ q̳u̳e̳ r̳e̳s̳p̳o̳n̳d̳a̳ d̳e̳ f̳o̳r̳m̳a̳ c̳o̳r̳r̳e̳c̳t̳a̳ a̳u̳m̳e̳n̳t̳a̳ s̳u̳ p̳u̳n̳t̳u̳a̳c̳i̳ó̳n̳.̳S̳i̳ a̳m̳b̳o̳s̳ j̳u̳g̳a̳d̳o̳r̳e̳s̳ r̳e̳s̳p̳o̳n̳d̳e̳n̳ c̳o̳r̳r̳e̳c̳t̳a̳m̳e̳n̳t̳e̳ u̳n̳a̳ p̳r̳e̳g̳u̳n̳t̳a̳,̳ e̳l̳ p̳r̳i̳m̳e̳r̳o̳ q̳u̳e̳ h̳a̳y̳a̳ r̳e̳s̳p̳o̳n̳d̳i̳d̳o̳ s̳e̳ l̳l̳e̳v̳a̳ p̳u̳n̳t̳o̳s̳ a̳d̳i̳c̳i̳o̳n̳a̳l̳e̳s̳.̳E̳l̳ j̳u̳g̳a̳d̳o̳r̳ q̳u̳e̳ r̳e̳s̳p̳o̳n̳d̳a̳ i̳n̳c̳o̳r̳r̳e̳c̳t̳a̳m̳e̳n̳t̳e̳ n̳o̳ s̳u̳m̳a̳ n̳i̳n̳g̳u̳n̳ p̳u̳n̳t̳o̳.̳
		`
)
