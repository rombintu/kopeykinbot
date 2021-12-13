package bot

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Default struct {
	Token      string
	DebugLevel string
}

type Api struct {
	Debug         bool
	TimeoutUpdate int
}

type Config struct {
	Default Default
	Api     Api
}

// Struct BOT
type Bot struct {
	Api    *tgbotapi.BotAPI
	Config *Config
	Logger *logrus.Logger
}

// Return configuration
func GetConfig() *Config {
	confFile, err := os.ReadFile("./config.toml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf Config

	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}

// Return new bot
func NewBot() *Bot {
	return &Bot{
		Config: GetConfig(),
		Logger: logrus.New(),
	}
}

func (bot *Bot) configureLogger() error {
	level, err := logrus.ParseLevel(bot.Config.Default.DebugLevel)
	if err != nil {
		return err
	}

	bot.Logger.SetLevel(level)

	return nil
}

func (bot *Bot) InitApi() error {
	var token string
	if bot.Config.Default.Token != "" {
		token = bot.Config.Default.Token
	} else {
		token = os.Getenv("TOKEN")
	}
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot.Api = api

	if err := bot.configureLogger(); err != nil {
		return err
	}
	return nil
}
