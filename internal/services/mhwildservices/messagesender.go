package mhwildservices

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const discordMaxMessageLen = 50

// SendChunkedMessageSlowly splits msg on blank lines and sends as few messages as possible,
// pausing `delay` between each send to avoid rate limits. Pass delay=0 to use a safe default.
func SendChunkedMessageSlowly(s *discordgo.Session, channelID string, msg string) error {
	delay := 5 * time.Second

	// If it fits in one message, send directly
	if len(msg) <= discordMaxMessageLen {
		_, err := s.ChannelMessageSend(channelID, msg)
		return err
	}

	// Split into blocks separated by blank lines
	blocks := strings.Split(msg, "\n\n")
	var chunks []string
	var cur strings.Builder

	for i, b := range blocks {
		part := b
		if i < len(blocks)-1 {
			part += "\n\n" // restore the blank line we split on
		}
		if cur.Len()+len(part) > discordMaxMessageLen {
			chunks = append(chunks, cur.String())
			cur.Reset()
		}
		cur.WriteString(part)
	}
	if cur.Len() > 0 {
		chunks = append(chunks, cur.String())
	}

	// Send with a small pause between chunks
	for i, c := range chunks {
		// show typing to signal progress
		_ = s.ChannelTyping(channelID)

		if _, err := s.ChannelMessageSend(channelID, c); err != nil {
			return fmt.Errorf("failed to send chunk %d/%d: %w", i+1, len(chunks), err)
		}
		time.Sleep(delay)
	}

	return nil
}
