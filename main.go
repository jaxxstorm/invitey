package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	token         = kingpin.Flag("token", "Slack token").String()
	email         = kingpin.Flag("email", "User email").String()
	channelPrefix = kingpin.Flag("channel_prefix", "Channel prefix to invite to").String()
)

func main() {

	kingpin.Parse()
	kingpin.Version("0.0.1")

	if *token == "" {
		kingpin.FatalUsage("Must specify a slack bot token")
	}

	if *email == "" {
		kingpin.FatalUsage("Must specify a user to invite")
	}

	if *channelPrefix == "" {
		kingpin.FatalUsage("Must specify a channel prefix to search for")
	}

	api := slack.New(*token)

	// get all the channels in the workspace
	// FIXME: we need to paginate
	channels, _, err := api.GetConversations(&slack.GetConversationsParameters{
		ExcludeArchived: true,
	})
	if err != nil {
		kingpin.FatalIfError(err, "error getting conversations")
	}

	prompt := promptui.Prompt{
		Label:     "Would you like to add the user to the following found channels?",
		IsConfirm: true,
	}

	var channelsToInviteTo []slack.Channel

	// FIXME: We need to paginate here
	for _, channel := range channels {
		isInviteChannel := strings.HasPrefix(channel.Name, *channelPrefix)
		if isInviteChannel {
			fmt.Println(channel.Name)
			channelsToInviteTo = append(channelsToInviteTo, channel)
		}
	}

	_, err = prompt.Run()

	if err != nil {
		log.Fatal("User cancelled")
		return
	}

	// lookup the user by email
	user, err := api.GetUserByEmail(*email)
	if err != nil {
		kingpin.FatalIfError(err, "error getting users")
	}

	// set up logrus
	logger := log.WithFields(log.Fields{"user": user.Name})

	// loop through all channels
	for _, channel := range channelsToInviteTo {

		logger = logger.WithFields(log.Fields{"channel_name": channel.Name})
		logger.Info("Checking channel")

		isInviteChannel := strings.HasPrefix(channel.Name, *channelPrefix)

		if isInviteChannel {
			logger.Infof("Inviting user to channel")
			_, err := api.InviteUsersToConversation(channel.ID, user.ID)
			if err != nil {
				switch err.Error() {
				case "already_in_channel":
					logger.Info("User is already in channel")
				case "not_in_channel":
					logger.Warn("slack app is not present in channel, cannot invite user")
				default:
					logger.Errorf("unhandled error: %w", err.Error())
				}
			}

		}

		logger.Info("User successfully invited to channel")

	}

}
