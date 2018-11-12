package ns

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

// path stores information on the nmap location
var path string

//
// Scanner struct for some configuration
//
type Scanner struct {
	ForwardOutput bool
}

//
// Result ...
//
type Result struct {
	Scanner string  `xml:"scanner,attr"` // Information about the used scanner
	Args    string  `xml:"args,attr"`    // The whole command including arguments
	Start   uint32  `xml:"start,attr"`   // Starttime
	Hosts   []*Host `xml:"host"`         // Slice of hosts
}

func init() {
	var err error
	path, err = exec.LookPath("nmap")
	if err != nil {
		panic(err)
	}
}

//
// Query executes the nmap command with an additional -oX param for XML output
// After execution, the XML is parsed to a Result struct and removed.
//
func (s *Scanner) Query(args ...string) (result Result, err error) {

	xmlFile := fmt.Sprintf(".%d.xml", time.Now().UnixNano())
	args = append(args, "-oX", xmlFile)

	defer func() {
		if _, e := os.Stat(xmlFile); e == nil {
			e = os.Remove(xmlFile)
			if e != nil {
				log.Println(e)
			}
		}
	}()

	cmd := exec.Command(path, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	cmd.Start()

	if s.ForwardOutput {
		sc := bufio.NewScanner(stdout)
		sc.Split(bufio.ScanLines)
		for sc.Scan() {
			out := sc.Text()
			fmt.Println(out)
		}
	}

	se := bufio.NewScanner(stderr)
	se.Split(bufio.ScanLines)
	for se.Scan() {
		err = fmt.Errorf("%s", se.Text())
		return
	}

	cmd.Wait()

	return ParseXML(xmlFile)
}

//
// ParseXML reads the passed xml file and unmarshals the content to a Result struct
//
func ParseXML(xmlFile string) (result Result, err error) {
	dat, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return
	}

	err = xml.Unmarshal(dat, &result)
	return
}
