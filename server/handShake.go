package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
)

func HandShakeServer(player Player) bool {
	return startUpMenu(player)
}

func HandShakeClient(currentSocket net.Conn) bool {

	reader := bufio.NewReader(os.Stdin)
	for {
		messageFromServer, err := common.Receive(currentSocket)
		if err != nil {
			fmt.Println("Server disconnected. Client exiting...")
			return false
		}
		if messageFromServer == CloseConnectionCommand {
			fmt.Println("Client exiting...")

			return false
		}
		fmt.Println("->: " + messageFromServer)

		fmt.Print(">> ")
		textFromPrompt, _ := reader.ReadString('\n')

		common.Send(currentSocket, textFromPrompt)
	}

}

// Devuelve true si puede comenzar a buscar partida correctamente tras el menú , false en caso contrario
func startUpMenu(player Player) bool {

	puedeBuscarPartida := false
	for !puedeBuscarPartida {

		messageFromClient, err := sendMainMenuOptions(player)
		if err != nil {
			logger.LogError(err)
			return false
		}
		fmt.Println("-> ", messageFromClient)

		if messageFromClient == common.OptionOne {

			fmt.Println("Player " /*player.name*/, player.id, " selected option 1, searching match...")
			puedeBuscarPartida = true
			// ... //

		} else if messageFromClient == common.OptionTwo {

			err = sendHelpSubMenuOptions(player)
			if err != nil {
				logger.LogError(err)
				return false
			}

		} else if messageFromClient == common.OptionThree {

			disconnectPlayerFromMenu(player)
			return false

		}

	}

	return puedeBuscarPartida

}

//Muestra opciones del menú y le pide al cliente que elija una
func sendMainMenuOptions(player Player) (string, error) {

	defer fmt.Println("Player " /*player.name*/, player.id, " redirected to main menu")

	// Saluda al usuario y le muestra el menú

	common.Send(player.socket, common.WelcomeMessage+strconv.Itoa(player.id)+common.MainMenuOptions)
	// Recibe su respuesta
	messageFromClient, err := common.Receive(player.socket)
	return messageFromClient, err

}

//Muestra opciones del Submenú de HELP y le pide al cliente que elija una
func sendHelpSubMenuOptions(player Player) error {

	fmt.Println("Player" /*player.name*/, player.id, "selected option 2, showing help...")

	defer fmt.Println("Player" /*player.name*/, player.id, "redirected to main menu")

	var (
		messageFromClient string
		err               error
		volverAMainMenu   bool
	)

	for !volverAMainMenu {

		common.Send(player.socket, common.HelpMessage+common.HelpMenuOptions)

		messageFromClient, err = common.Receive(player.socket)

		if messageFromClient == common.OptionOne {
			volverAMainMenu = true
		}

	}

	return err

}

// Desconecta al cliente del jugador que solicitó opcion 3 (Exit) del Menú
func disconnectPlayerFromMenu(player Player) {
	fmt.Println("Player selected option 3, disconnecting client...")
	common.Send(player.socket, CloseConnectionCommand)
	fmt.Println("Client disconnected")
}
