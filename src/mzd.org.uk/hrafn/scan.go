package main

/*-----------------------------------------------------------------------------
 **
 ** - hrafn -
 **
 ** Copyright (c) 2015-18 by Krogoth of
 ** Ministry of Zombie Defense - http://www.mzd.org.uk/
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** krogoth @ Ministry of Zombie Defense
 **
-----------------------------------------------------------------------------*/

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"mzd.org.uk/hrafn/common"
	"mzd.org.uk/hrafn/data"
	"os"
	"os/exec"
	"path/filepath"
)

func Scan(rec interface{}) error {

	record := rec.(data.Record)

	common.LogInfo("Scanning domain.", logrus.Fields{"domain": record.Domain})

	if record.ScanTLS {

		common.LogDebug("Scanning TLS.", logrus.Fields{"domain": record.Domain})

		err := scanTLS(record.Domain)
		if err != nil {

			common.LogError("Scanning TLS returned an error.", logrus.Fields{"domain": record.Domain, "error": err})
			return err
		}
	}

	if record.ScanPorts {

		common.LogDebug("Scanning ports.", logrus.Fields{"domain": record.Domain})

		err := scanPorts(record.Domain)
		if err != nil {

			common.LogError("Scanning ports returned an error.", logrus.Fields{"domain": record.Domain, "error": err})
			return err
		}
	}

	return nil
}

func scanTLS(domain string) error {

	sslyze := common.GetStringFromConfig("tools.sslyze")
	if len(sslyze) <= 0 {

		common.LogError("Could not find executable in config.", logrus.Fields{"exec": "sslyze"})
		return exec.ErrNotFound
	}

	outFile := domain + "_sslyze_current.xml"

	opt := "--regular"
	sni := "--sni=" + domain
	out := "--xml_out=" + outFile

	sslyze = "python"

	ex, err := os.Executable()
	if err != nil {

		common.LogError("Cant seem to find myself?", logrus.Fields{"error": err})
		return exec.ErrNotFound
	}

	cmd := exec.Command(sslyze, "-m", "sslyze", opt, sni, out, domain)

	exPath := filepath.Dir(ex)
	cmd.Dir = exPath

	devMode := common.GetLogLevel() == "debug"

	var stdout io.ReadCloser
	var stderr io.ReadCloser

	if devMode {
		// todo, redirect output only, if in dev mode
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			common.LogError("Stdout Pipe returned an error", logrus.Fields{"error": err})
		}
		stderr, err = cmd.StderrPipe()
		if err != nil {
			common.LogError("Sterr Pipe returned an error", logrus.Fields{"error": err})
		}
	}

	err = cmd.Start()

	if err != nil {

		common.LogError("Command start returned an error", logrus.Fields{"error": err})
		return errors.New("Command start failed")
	}

	if devMode {
		go copyOutput(stdout)
		go copyOutput(stderr)
	}

	cmd.Wait()

	return nil
}

func scanPorts(domain string) error {

	nmap := common.GetStringFromConfig("tools.nmap")
	if len(nmap) <= 0 {

		common.LogError("Could not find executable in config.", logrus.Fields{"exec": "nmap"})
		return exec.ErrNotFound
	}

	outFile := domain + "_nmap_current.xml"

	s := "-sV"
	v := "-v"
	p := "-Pn"
	port := "-p http,https"
	script := "--script=http-headers"
	o := "-oX=" + outFile

	ex, err := os.Executable()
	if err != nil {
		common.LogError("Cant seem to find myself?", logrus.Fields{"error": err})

	}

	cmd := exec.Command(nmap, s, v, p, port, script, domain, o)

	exPath := filepath.Dir(ex)
	cmd.Dir = exPath

	devMode := common.GetLogLevel() == "debug"

	var stdout io.ReadCloser
	var stderr io.ReadCloser

	if devMode {

		stdout, err = cmd.StdoutPipe()
		if err != nil {
			common.LogError("Stdout Pipe returned an error", logrus.Fields{"error": err})
		}
		stderr, err = cmd.StderrPipe()
		if err != nil {
			common.LogError("Sterr Pipe returned an error", logrus.Fields{"error": err})
		}
	}

	err = cmd.Start()
	if err != nil {

		common.LogError("Command start returned an error", logrus.Fields{"error": err})
		return errors.New("Command start failed")
	}

	if devMode {
		go copyOutput(stdout)
		go copyOutput(stderr)
	}

	cmd.Wait()

	return nil
}

func copyOutput(r io.Reader) {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		common.LogDebug(scanner.Text(), nil)
	}
}
