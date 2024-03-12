package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// var hamtaros = [2]string{
// 	"https://i.pinimg.com/736x/6f/df/bc/6fdfbc41d6a8e26d4b9073bc1afd899f.jpg",
// 	"https://i.pinimg.com/736x/df/a9/03/dfa9037c75441e76ee8e4df2e75eb02d.jpg",
// }

// Example from https://www.youtube.com/watch?v=XuFq7NW3ii4

type photo struct {
	title string
	url   string
}

func newPhoto(title string, url string) photo {
	p := photo{title: title, url: url}
	return p
}

func main() {
	// Load env file, defaults to .env if not specified
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	appID := os.Getenv("APP_ID")
	sess, initErr := discordgo.New("Bot " + token)
	if initErr != nil {
		log.Fatal(initErr)
	}

	// Register slash commands to app, please note that if guildID is set to
	// empty string it is ignored - this is only used for dev servers
	commands, overrideErr := sess.ApplicationCommandBulkOverwrite(appID, "", []*discordgo.ApplicationCommand{
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
		{
			Name:        "d20",
			Description: "Rolls a d20",
		},
		{
			Name:        "randtaro",
			Description: "Random picture of Hatmaro",
		},
	})
	if overrideErr != nil {
		// Handle the error
		log.Fatal(overrideErr)
	}

	// Register handler for slash commander interactions
	sess.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				log.Fatal(err)
			}
		case "d20":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: rollDice(20),
					},
				},
			)
			if err != nil {
				log.Fatal(err)
			}

		case "randtaro":
			var hamtaros []photo
			hamtaros = append(hamtaros,
				newPhoto("Sleepytaro with friends", "https://github.com/akam/hamtaro/assets/19315796/7098538e-09e8-49fe-b68f-92dd200f9e68"),
				newPhoto("Sleepy boxtaro", "https://github.com/akam/hamtaro/assets/19315796/50741eac-8dfb-4ff3-a16e-21e2ffb0ca6b"),
				newPhoto("Close up taro", "https://github.com/akam/hamtaro/assets/19315796/c054ab54-a5bf-437c-8d32-cffa2323e517"),
				newPhoto("Partytaro", "https://github.com/akam/hamtaro/assets/19315796/ab574974-2a9e-4927-ba0b-5e823dddda47"),
				newPhoto("Cartaro", "https://github.com/akam/hamtaro/assets/19315796/3cdd5a82-c78e-47fb-b77f-76bda1cf8c7f"),
				newPhoto("Stalk taro", "https://github.com/akam/hamtaro/assets/19315796/6611a7f4-6db4-4e53-a57f-c50edef63957"),
				newPhoto("Chickentaro", "https://github.com/akam/hamtaro/assets/19315796/edaa2e3a-362b-4b9b-88b8-85471993f7db"),
				newPhoto("Cooltaro", "https://github.com/akam/hamtaro/assets/19315796/ef2c0e7b-1049-475b-af09-7dcadb701ed6"),
			)
			var hamtaro = hamtaros[rand.Intn(len(hamtaros))]
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: hamtaro.title,
								// Description: "picture of cat",
								Image: &discordgo.MessageEmbedImage{

									URL: hamtaro.url,
								},
							},
						},
					},
				},
			)
			if err != nil {
				log.Fatal(err)
			}
		}

	})

	// Register handler for message listener
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			// If the message is from the session user, ignore processing
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	openErr := sess.Open()

	if openErr != nil {
		log.Fatal(openErr)
	}

	// Defer is used to ensure that sess.Close() is run regardless of what it exits
	defer sess.Close()

	fmt.Println("bot is online!")

	// This part is used to help gracefully handle operating errors to allow for a cleaner shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// On shutdown, remove all added commands for given bot (iterate through commands and delete each ID added)
	for _, v := range commands {
		err := sess.ApplicationCommandDelete(sess.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
	fmt.Println("Commands deleted, shutting down bot")
}

func rollDice(number int) string {
	roll := rand.Intn(number) + 1
	return fmt.Sprintf("You rolled a: %v", roll)
}
