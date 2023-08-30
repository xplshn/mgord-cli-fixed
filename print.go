// Print everything in pretty-colored, valid YAML.
// Why YAML? No practical reason, I just think it looks nice and I guess if you
// really wanted to it would be easier to parse than unformatted text for
// something like a nushell plugin

package cli

import (
	"fmt"
	"strings"
	"path/filepath"

	check   "github.com/mgord9518/aisap/spooky"
	clr     "github.com/gookit/color"
	xdg     "github.com/adrg/xdg"
	permissions "github.com/mgord9518/aisap/permissions"
)

func makeDevPretty(str string) string {
	str = filepath.Clean(str)

	if len(str) > 5 && str[0:5] == "/dev/" {
		str = strings.Replace(str, "/dev/", "", 1)
	}

	return str
}

func ListPerms(p *permissions.AppImagePerms) {
	for i, val := range p.Devices {
		p.Devices[i] = makeDevPretty(val)
	}
	clr.Println("<yellow>permissions</>:")
	if p.Level >= 1 {
		List("level", p.Level, 11)
		prettyListFiles("filesystem", p.Files, 11)
		List("devices", p.Devices, 11)
		prettyListSockets("sockets", p.Sockets, 11)
	} else {
		clr.Println("  level:      <lightYellow>0</>")
		clr.Println("  filesystem: <lightYellow>ALL</>")
		clr.Println("  devices:    <lightYellow>ALL</>")
		clr.Println("  sockets:    <lightYellow>ALL</>")
	}
}

func List(a ...interface{}) {
	for i := range(a) {
		if i == 0 {
			clr.Printf("  %s:", a[0])
			continue
		} else if i == len(a)-1 {
			break	
		}

		// pad with spaces until the requested lengh is reached
		n := a[len(a)-1].(int)
		str := a[0].(string)
		for i := len(str); i < n; i++ {
			fmt.Print(" ")
		}

		switch v := a[i].(type) {
		default:
			panic("invalid type!")
		case string:
			clr.Printf(" <green>%s</>\n", a[i])
		case []string:
			clr.Gray.Print("[")
			for i := range(v) {
				if i > 0 {
					clr.Gray.Print(", ")
				}
				clr.Green.Print(v[i])
			}
			clr.Gray.Println("]")
		case int:
			clr.Green.Printf(" %d\n", a[i])
		}
	}

}

// Like `prettyList` but highlights spooky files in orange
func prettyListFiles(a ...interface{}) {
	for i := range(a) {
		if i == 0 {
			fmt.Printf("  %s:", a[0])
			continue
		} else if i == len(a)-1 {
			break	
		}

		// pad with spaces until the requested lengh is reached
		n := a[len(a)-1].(int)
		str := a[0].(string)
		for i := len(str); i < n; i++ {
			fmt.Print(" ")
		}

		switch v := a[i].(type) {
		default:
			panic("invalid type!")
		case []string:
			clr.Gray.Print("[")
			for i := range(v) {
				if i > 0 {
					clr.Gray.Print(", ")
				}
				v[i] = strings.Replace(v[i], xdg.Home, "~", 1)

				if check.IsSpooky(v[i]) {
					clr.Printf("<lightYellow>%s</>", v[i])
				} else {
					clr.Printf("<green>%s</>", v[i])
				}
			}
			clr.Gray.Println("]")
		}
	}

}

// Like `prettyList` but highlights dangerous sockets in orange
func prettyListSockets(name string, a []permissions.Socket, n int) {
	fmt.Printf("  %s:", name)

	for i := len(name); i < n; i++ {
		fmt.Print(" ")
	}

	clr.Gray.Print("[")
	for i := range(a) {
		socketString := string(a[i])

		if i > 0 {
			clr.Gray.Print(", ")
		}
		socketString = strings.Replace(socketString, xdg.Home, "~", 1)

		if a[i] == permissions.Session || a[i] == permissions.X11 {
			clr.Printf("<lightYellow>%s</>", a[i])
		} else {
			clr.Printf("<green>%s</>", a[i])
		}
	}

	clr.Gray.Println("]")
}
