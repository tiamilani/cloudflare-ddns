package config

import (
	"github.com/favonia/cloudflare-ddns/internal/notifier"
	"github.com/favonia/cloudflare-ddns/internal/file"
	"github.com/favonia/cloudflare-ddns/internal/pp"
)

// ReadAndAppendShoutrrrURL reads the URLs separated by the newline.
func ReadAndAppendShoutrrrURL(ppfmt pp.PP, key string, field *notifier.Notifier) bool {
	vals := GetenvAsList(key, "\n")
	if len(vals) == 0 {
		return true
	}

	ppfmt.InfoOncef(pp.MessageExperimentalShoutrrr, pp.EmojiHint,
		"You are using the experimental shoutrrr support added in version 1.12.0")

	s, ok := notifier.NewShoutrrr(ppfmt, vals)
	if !ok {
		return false
	}

	// Append the new monitor to the existing list
	*field = notifier.NewComposed(*field, s)
	return true
}

func ReadAndAppendShoutrrrURLFromFile(ppfmt pp.PP, key string, field *notifier.Notifier) bool {
	shoutrrrFile := Getenv(key)
	if shoutrrrFile == "" {
		return "", true
	}

	vals, ok := file.ReadString(ppfmt, shoutrrrFile)
	if !ok {
		return "", false
	}

	if vals == "" {
		ppfmt.Noticef(pp.EmojiUserError, "The file [%s] specified by %s is empty", shoutrrrFile, key)
		return false
	}

	valsList := SplitAndTrim(vals, "\n")
	
	ppfmt.InfoOncef(pp.MessageExperimentalShoutrrr, pp.EmojiHint,
		"You are using the experimental shoutrrr_file support!!!")

	s, ok := notifier.NewShoutrrr(ppfmt, valsList)
	if !ok {
		return false
	}

	// Append the new monitor to the existing list
	*field = notifier.NewComposed(*field, s)
	return true
}
