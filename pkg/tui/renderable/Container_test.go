package renderable

import (
	"os"
	"strings"
	"testing"
)

func Test_container_rendering_simple_text(t *testing.T) {
	cont := NewContainer()

	text := NewText("OK", 3, 4)
	cont.AddItem(&text)

	var rendered = strings.ReplaceAll(cont.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "|||____OK" {
		t.Errorf("Container render method failed: expected %s, got %s", "|||____OK", cont.Render())
	}
}

func Test_container_rendering_multiple_text_not_overlapped(t *testing.T) {
	cont := NewContainer()

	text2 := NewText("OK2", 0, 0)
	text := NewText("OK", 0, 0)
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

	text := NewText("OK", 0, 0)
	text2 := NewText("OK2", 0, 0)
	text3 := NewText("ABC", 0, 0)
	text4 := NewText("XXX", 0, 0)
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

	text := NewText("OK", 0, 0)
	text2 := NewText("OK2", 0, 0)
	text3 := NewText("ABC", 0, 0)
	text4 := NewText("XXX", 0, 0)
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

	tempText := NewText("CCC", 0, 0)
	cont2.AddItem(&tempText)

	textSpaces := NewText("   ", 0, 0)
	textSpaces.SetPosition(6, 2)

	cont2.AddItem(&textSpaces)

	var rendered = strings.ReplaceAll(cont2.Render(), " ", "_")
	rendered = strings.ReplaceAll(rendered, "\n", "|")
	if rendered != "CCC|||_|_|_|_ABXXXOK2" {
		t.Errorf("Container render method failed: expected %s, got %s", "CCC|||_|_|_|_ABXXXOK2", rendered)
	}

	print(cont2.Render())
}

func Test_container_rendering_fixed_size(t *testing.T) {
	cont := NewContainer()

	text := NewText("OK", 0, 0)
	text2 := NewText("OK2", 0, 0)
	text3 := NewText("ABC", 0, 0)
	text4 := NewText("XXX", 0, 0)
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

	tempText := NewText("CCC", 0, 0)
	cont2.AddItem(&tempText)

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

func Test_container_rendering_border(t *testing.T) {
	box := make([]byte, 1000)
	file, err := os.Open("boxWithBorder.txt")
	if err != nil {
		t.Errorf("Could not find read boxWithBorder.txt")
	}

	file.Read(box)

	c := NewContainer()

	c.SetFixedSize(12, 5)
	if !c.SetBorder('|', '+', '-', '+', '|', '+', '-', '+') {
		t.Errorf("SetBorder returned false")
	}

	tempText := NewText("My text is", 0, 0)
	c.AddItem(&tempText)
	txt1 := Text{Text: "here."}
	txt1.SetPosition(1, 0)
	c.AddItem(&txt1)
	txt2 := Text{Text: "In the box"}
	txt2.SetPosition(2, 0)
	c.AddItem(&txt2)

	rendered := []byte(c.Render())

	for idx, char := range rendered {
		if char != box[idx] {
			t.Errorf("At position %d, found %c but expected %c", idx, char, box[idx])
		}
	}
	print(c.Render())
}

func Test_container_centralization_horizontal(t *testing.T) {
	text := NewText("Text", 0, 0)
	cont := NewContainer()
	cont.AddItem(&text)

	cont.SetCentralization(true, false)
	cont.SetFixedSize(10, 2)

	expected := "___Text___\n__________"
	rendered := strings.ReplaceAll(cont.Render(), " ", "_")
	if rendered != expected {
		t.Errorf("Container centralization failed: expected %s, got %s", expected, rendered)
	}
}
