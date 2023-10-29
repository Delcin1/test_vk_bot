package main

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

func getMainMenu() object.MessagesKeyboard {
	mainMenu := object.NewMessagesKeyboard(true)
	mainMenu.AddRow()
	mainMenu.AddTextButton("Доступные биты", payload, color)
	mainMenu.AddTextButton("Мои покупки", payload, color)
	mainMenu.AddRow()
	mainMenu.AddTextButton("Профиль", payload, color)
	mainMenu.AddTextButton("Тех. поддержка", payload, color)

	return *mainMenu
}

func getBeatsMenu(isFirstPage, isLastPage bool) object.MessagesKeyboard {
	beatsMenu := object.NewMessagesKeyboard(true)
	beatsMenu.AddRow()
	if !isFirstPage {
		beatsMenu.AddTextButton("Предыдущая", payload, color)
	}
	if !isLastPage {
		beatsMenu.AddTextButton("Следующая", payload, color)
	}
	beatsMenu.AddRow()
	beatsMenu.AddTextButton("Главное меню", payload, color)

	return *beatsMenu
}
