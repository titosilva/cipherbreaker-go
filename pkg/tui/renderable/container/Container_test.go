package container

import (
	"strings"
	"testing"

	. "github.com/titosilva/cipherbreaker-go/pkg/tui/renderable"
)

func Test_container_rendering_simple_text(t *testing.T) {
	cont := NewContainer()

	var text = Text{Text: "OK"}
	text.SetPosition(3, 4)
	cont.AddItem(&text)

	var rendered = strings.ReplaceAll(cont.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "|||____OK" {
		t.Errorf("Container render method failed: expected %s, got %s", "|||____OK", cont.Render())
	}
}

func Test_container_rendering_multiple_text_not_overlapped(t *testing.T) {
	cont := NewContainer()

	var text2 = Text{Text: "OK2"}
	var text = Text{Text: "OK"}
	text.SetPosition(3, 4)
	cont.AddItem(&text)
	cont.AddItem(&text2)

	var rendered = strings.ReplaceAll(cont.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "OK2|||____OK" {
		t.Errorf("Container render method failed: expected %s, got %s", "OK2||____OK", rendered)
	}
}

func Test_container_rendering_multiple_text_overlapped(t *testing.T) {
	cont := NewContainer()

	var text = Text{Text: "OK"}
	var text2 = Text{Text: "OK2"}
	var text3 = Text{Text: "ABC"}
	var text4 = Text{Text: "XXX"}
	text.SetPosition(3, 4)
	text2.SetPosition(3, 5)
	text3.SetPosition(3, 0)
	text4.SetPosition(3, 2)
	cont.AddItem(&text)
	cont.AddItem(&text2)
	cont.AddItem(&text3)
	cont.AddItem(&text4)

	var rendered = strings.ReplaceAll(cont.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "|||ABXXXOK2" {
		t.Errorf("Container render method failed: expected %s, got %s", "|||ABXXXOK2", rendered)
	}
}

func Test_container_rendering_container_inside_container(t *testing.T) {
	cont := NewContainer()

	var text = Text{Text: "OK"}
	var text2 = Text{Text: "OK2"}
	var text3 = Text{Text: "ABC"}
	var text4 = Text{Text: "XXX"}
	text.SetPosition(3, 4)
	text2.SetPosition(3, 5)
	text3.SetPosition(3, 0)
	text4.SetPosition(3, 2)
	cont.AddItem(&text)
	cont.AddItem(&text2)
	cont.AddItem(&text3)
	cont.AddItem(&text4)

	cont.SetPosition(3, 1)

	cont2 := NewContainer()
	cont2.AddItem(&cont)

	cont2.AddItem(&Text{Text: "CCC"})

	textSpaces := Text{Text: "   "}
	textSpaces.SetPosition(6, 2)

	cont2.AddItem(&textSpaces)

	var rendered = strings.ReplaceAll(cont2.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "CCC|||_|_|_|_ABXXXOK2" {
		t.Errorf("Container render method failed: expected %s, got %s", "CCC|||_|_|_|_ABXXXOK2", rendered)
	}
}

func Test_container_rendering_fixed_size(t *testing.T) {
	cont := NewContainer()

	var text = Text{Text: "OK"}
	var text2 = Text{Text: "OK2"}
	var text3 = Text{Text: "ABC"}
	var text4 = Text{Text: "XXX"}
	text.SetPosition(3, 4)
	text2.SetPosition(3, 5)
	text3.SetPosition(3, 0)
	text4.SetPosition(3, 2)
	cont.AddItem(&text)
	cont.AddItem(&text2)
	cont.AddItem(&text3)
	cont.AddItem(&text4)

	cont.SetPosition(3, 1)

	cont2 := NewContainer()
	cont2.AddItem(&cont)

	cont2.AddItem(&Text{Text: "CCC"})

	textSpaces := Text{Text: "   "}
	textSpaces.SetPosition(6, 2)

	cont2.AddItem(&textSpaces)

	cont2.SetFixedSize(3, 3)

	var rendered = strings.ReplaceAll(cont2.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "CCC|___|___" {
		t.Errorf("Container render method failed: expected %s, got %s", "CCC|___|___", rendered)
	}

	expected := "CCC_______|__________|__________|__________|__________|__________|_ABXXXOK2_|__________|__________|__________"
	cont2.SetFixedSize(10, 10)
	var rendered2 = strings.ReplaceAll(cont2.Render(), " ", "_")
	rendered2 = strings.ReplaceAll(rendered2, "\n", "|")
	if rendered2 != expected {
		t.Errorf("Container render method failed: expected %s, got %s", expected, rendered2)
	}
}
