package logchain

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	t.Log(now.UTC().Format(RFC339micro))
}

func TestMakeString(t *testing.T) {
	msg := NewSyslogMessage()
	msg.Message = "Kilroy was here!ÅÄÖ"
	text := msg.MakeString()
	t.Log(text)
}
