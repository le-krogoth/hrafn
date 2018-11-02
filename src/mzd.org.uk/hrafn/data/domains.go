package data

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
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"io"
	"mzd.org.uk/hrafn/common"
	"os"
	"strconv"
)

// File content: domain,scan_tls,scan_ports
type Record struct {
	Domain    string // -> key
	ScanTLS   bool
	ScanPorts bool
}

var domains []Record

//var domains map[string]record

func LoadDomains() {

	//domains = make(map[string]record)

	sFile := common.GetStringFromConfig("files.domains")

	file, err := os.Open(sFile)
	if err != nil {
		common.LogFatal("Error reading domains file", logrus.Fields{"file": sFile, "error": err})
		return
	}

	defer file.Close()

	r := csv.NewReader(file)
	// todo get from config
	r.Comma = ','
	r.Comment = '#'

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			common.LogFatal("Failing when reading domains file.", logrus.Fields{"file": sFile, "error": err})
			return
		}

		scanTLS, err := strconv.ParseBool(line[1])
		if err != nil {
			common.LogError("There is an error in your domains list.", logrus.Fields{"file": sFile, "error": err})
			scanTLS = false
		}

		scanPorts, err := strconv.ParseBool(line[2])
		if err != nil {
			common.LogError("There is an error in your domains list.", logrus.Fields{"file": sFile, "error": err})
			scanTLS = false
		}

		rec := Record{
			Domain:    line[0],
			ScanTLS:   scanTLS,
			ScanPorts: scanPorts,
		}

		domains = append(domains, rec)
	}
}

func GetDomainCount() int {

	return len(domains)
}

func GetDomainRecord(id int) (Record, error) {

	// check size, return error
	return domains[id], nil
}
