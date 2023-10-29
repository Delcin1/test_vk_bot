package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

const (
	payload = "payload"
	color   = ""
)

func main() {
	token := "vk1.a.g1guX9aOPyB4c73pgw3JvBy_t4Zs4deVxGvWILff-u75W1Vvusc2N1HD4_DvHASlDTJI43KEhPNC8fYeVYGV-YyJyHvYNGI30jjesooIU9__VvB4fY-lyrkrjeKoF66eXwRT6g2pchYLWgu0UP0xZ7dk9bh41hFGFhRt2Hc1VnO0Bs0rck-3vGdIC1Zb1O0EPIK1GlufhDsVtDxIUz_9xg" // рекомендуется использовать os.Getenv("TOKEN")
	vk := api.NewVK(token)

	// Получаем информацию о группе
	group, err := vk.GroupsGetByID(api.Params{})
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем longpoll
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	mainMenu := getMainMenu()

	// Событие нового сообщения
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

		messagesSendBuilder := params.NewMessagesSendBuilder()

		if obj.Message.Text == "Доступные биты" {
			messagesSendBuilder.Message("Здесь будут доступные биты")
			messagesSendBuilder.RandomID(0)
			messagesSendBuilder.PeerID(obj.Message.PeerID)
			messagesSendBuilder.Keyboard(getBeatsMenu(true, false))

			_, err := vk.MessagesSend(messagesSendBuilder.Params)
			if err != nil {
				log.Fatal(err)
			}
		} else if obj.Message.Text == "Главное меню" || obj.Message.Text == "Начать" {
			messagesSendBuilder.Message("Главное меню")
			messagesSendBuilder.RandomID(0)
			messagesSendBuilder.PeerID(obj.Message.PeerID)
			messagesSendBuilder.Keyboard(mainMenu)

			_, err := vk.MessagesSend(messagesSendBuilder.Params)
			if err != nil {
				log.Fatal(err)
			}
		} else if obj.Message.Text == "Тех. поддержка" {
			messagesSendBuilder.Message("Напишите ваш вопрос и с вами свяжутся в ближайшее время")
			messagesSendBuilder.RandomID(0)
			messagesSendBuilder.PeerID(obj.Message.PeerID)
			messagesSendBuilder.Keyboard(mainMenu)

			_, err := vk.MessagesSend(messagesSendBuilder.Params)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			messagesSendBuilder.Message(fmt.Sprintf("Вам писал пользователь @id%d : \n%s ", obj.Message.PeerID, obj.Message.Text))
			messagesSendBuilder.RandomID(0)
			messagesSendBuilder.PeerID(205624380)

			_, err := vk.MessagesSend(messagesSendBuilder.Params)
			if err != nil {
				log.Fatal(err)
			}

			messagesSendBuilder.Message("Главное меню")
			messagesSendBuilder.RandomID(0)
			messagesSendBuilder.PeerID(obj.Message.PeerID)
			messagesSendBuilder.Keyboard(mainMenu)

			_, err = vk.MessagesSend(messagesSendBuilder.Params)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	// Запускаем Bots Longpoll
	log.Println("Start longpoll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
