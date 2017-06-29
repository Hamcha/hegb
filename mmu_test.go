package hegb

import (
	"fmt"
	"os"
	"testing"
)

// Test all instructions to check that they are all handled
func TestIORegisterPresence(t *testing.T) {
	handled := 0
	unhandled := 0
	jsrows := "<table class=\"reg\"><tr><th>Address</th><th>Register name</th><th>Read</th><th>Write</th></tr>"
	for i := 0; i <= 0x80; i++ {
		isok := true
		regid := ioregister(MIOJoypad + ioregister(i))
		// Skip unused registers
		if regid.String() == "<unused IO register>" || regid.String() == "<invalid IO register>" {
			continue
		}
		jsrows += fmt.Sprintf("<tr><td>%04X</td><td>%s</td>", uint16(regid), regid)
		if regfn, ok := ioreadhandlers[regid]; !ok {
			isok = false
			fmt.Fprintf(os.Stderr, "IOReg R/%02X | %s is MISSING!\n", uint16(regid), regid)
			jsrows += "<td class=\"regno\">✕</td>"
		} else {
			if regfn == nil {
				jsrows += "<td class=\"invalid\"></td>"
			} else {
				jsrows += "<td class=\"regok\">✓</td>"
			}
		}
		if regfn, ok := iowritehandlers[regid]; !ok {
			isok = false
			fmt.Fprintf(os.Stderr, "IOReg W/%02X | %s is MISSING!\n", uint16(regid), regid)
			jsrows += "<td class=\"regno\">✕</td>"
		} else {
			if regfn == nil {
				jsrows += "<td class=\"invalid\">✓</td>"
			} else {
				jsrows += "<td class=\"regok\">✓</td>"

			}
		}
		jsrows += "</tr>"
		if isok {
			handled++
		} else {
			unhandled++
		}
	}
	jsrows += "</table>"

	fmt.Fprintf(os.Stderr, "Summary: %d handled, %d missing (%.2f%% total)\n", handled, unhandled, (float32(handled) / float32(handled+unhandled) * 100))
	fmt.Fprintf(os.Stderr, "JS table code: %s\n", jsrows)
	if unhandled > 0 {
		t.Fail()
	}
}
