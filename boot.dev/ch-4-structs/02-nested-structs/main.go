package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func isValidUser(u user) bool {
	return u.name != "" && u.number != 0
}

func canSendMessage(mToSend messageToSend) bool {
	return isValidUser(mToSend.sender) && isValidUser(mToSend.recipient)
}
