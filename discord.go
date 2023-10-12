package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)



func main() {
	discord, err := discordgo.New("Bot " + "APIKEY")
	if err != nil{
		log.Fatal(err)
	}

	discord.AddHandler(func(sess *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == sess.State.User.ID {
			return
		}
		log.Printf("Received message: %s", message.Author.ID)

		log.Printf("Received message: %s", sess.State.User.ID)
		log.Printf("Received message: %s", message.Content) // Add this line for debugging
	
		if message.Content == "Hello" {
			sess.ChannelMessageSend(message.ChannelID, "Where's Ivan?")
		}
		if message.Author.ID == "535525711495299073" {
			log.Printf(message.Content) // Add this line for debugging
			sess.ChannelMessageSend(message.ChannelID, "Hi Will")
		}
	})
	
	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil{
		log.Fatal(err)
	}
	defer discord.Close()

	sc:= make(chan os.Signal, 1)
	signal.Notify(sc,syscall.SIGINT,syscall.SIGTERM,os.Interrupt)
	<- sc
}