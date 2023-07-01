package logchain

import "testing"

func TestStructuredData(t *testing.T) {
	sd := NewStructuredData("Olle@123")
	if sd == nil {
		t.Fatal("New returned nil")
	}

	emptyText := sd.MakeString()
	if emptyText != "" {
		t.Error("Not a empty string")
	}
	if sd.Name() != "Olle" {
		t.Error("Name error")
	}
	if sd.EnterpriseId() != "123" {
		t.Error("Invalid enterprise ID")
	}

	sd.ParameterList["olle"] = "123456"
	sd.ParameterList["pelle"] = "123"

	sdText := sd.MakeString()
	if sdText == "" {
		t.Error("Invalid SD test")
	}
	t.Log(sdText)
}
