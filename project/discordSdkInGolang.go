package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	ImageURL     string   `json:"imageUrl"`
	Tags         []string `json:"tags"`
	IsSoftDelete bool     `json:"isSoftDelete"`
	GitHubURL    string   `json:"github_url"`
	CreateAt     int64    `json:"created_at"`
	IsActive     bool     `json:"isActive"`
	Owner        struct {
		Name           string `json:"name"`
		PersonalGithub string `json:"personal_github"`
		Role           string `json:"role"`
	} `json:"owner"`
}

const (
	token     = "put discord bot token here"
	channelID = "put channel id here"
)

func main() {

	discord, err := discordgo.New("Bot " + token)
	// Create a new Discord session using the provided bot token.
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages
	// Register the messageCreate func as a callback for MessageCreate events.

	/* handler*/
	// discord.AddHandler(messageCreate)
	getMessages(discord, channelID)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func getMessages(dg *discordgo.Session, channelID string) {
	messages, err := dg.ChannelMessages(channelID, 50, "", "", "")
	if err != nil {
		log.Fatalf("Error fetching messages: %v", err)
	}
	for _, message := range messages {
		fmt.Printf("User: %s\nMessage: %s\n\n", message.Author.Username, message.Content)
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != channelID { // 将频道 ID 替换为你想要接收消息的频道 ID
		return
	}
	if m.Content == "" {
		fmt.Println("Message content is empty")
		return
	}
	// 打印消息内容到控制台
	fmt.Printf("%s: %s\n", m.Author.Username, m.Content)
	PrintJSON(m.Content)

	fmt.Printf("%T\n", m.Content)

	result := GetMarshaledData(m.Content)
	PrintJSON(result)

}

/* utilities */

func GetMarshaledData(jsonData string) *Project {
	// Create an instance of the BlogPost struct
	var project *Project

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(jsonData), &project)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return project
}

func PrintJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}
	fmt.Println(string(json))
}
