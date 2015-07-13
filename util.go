package main

const WHALE = "\xF0\x9f\x90\xb3"

type displayUtil struct{ color bool }

func (p *displayUtil) coloredString(text string, colorCode string) string {
	if p.color {
		return "\x1b[" + colorCode + "m" + text + "\x1b[39m"
	} else {
		return text
	}
}

func (p *displayUtil) headerString(text string) string {
	if p.color {
		return p.coloredString("\n#"+"\n#"+" "+WHALE+"  "+text+"\n#\n", "33")
	} else {
		return p.coloredString("\n#"+"\n#"+" "+text+"\n#\n", "33")
	}
}
