package logchain

import (
	"fmt"
	"strings"
)

type StructuredData struct {
	Identity      string            // SD name consist of <name>@<enterprise number>
	ParameterList map[string]string // SD parameter name/value
}

func NewStructuredData(identity string) *StructuredData {
	sd := StructuredData{identity, make(map[string]string)}
	return &sd
}

func (sd *StructuredData) Name() string {
	sdList := strings.Split(sd.Identity, "@")
	if len(sdList) > 0 {
		return sdList[0]
	}
	return ""
}

func (sd *StructuredData) EnterpriseId() string {
	sdList := strings.Split(sd.Identity, "@")
	if len(sdList) > 1 {
		return sdList[1]
	}
	return ""
}

func (sd *StructuredData) MakeString() string {
	if len(sd.ParameterList) <= 0 {
		return ""
	}
	var temp strings.Builder
	if sd.Identity != "" {
		temp.WriteString(fmt.Sprintf("[%s ", sd.Identity))
	}
	count := 0
	for key, value := range sd.ParameterList {
		if key == "" {
			continue
		}
		if count > 0 {
			temp.WriteString(" ")
		}
		temp.WriteString(fmt.Sprintf("%s=\"%s\"", key, value))
		count++
	}
	temp.WriteString("]")
	return temp.String()
}
