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
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"github.com/sirupsen/logrus"
	"mzd.org.uk/hrafn/common"
	"os"
	"path/filepath"
)

func GenerateReport() {

	sOutput := common.GetStringFromConfig("files.output")
	fCsv, err := os.Create(sOutput)
	if err != nil {
		common.LogError("Can't open report file.", logrus.Fields{"file": sOutput, "error": err})
	}

	defer fCsv.Close()

	eSslv2, err := xpath.Compile("count(//document/results/target/sslv2/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "sslv2", "error": err})
	}

	eSslv3, err := xpath.Compile("count(//document/results/target/sslv3/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "sslv3", "error": err})
	}

	eTlsv10, err := xpath.Compile("count(//document/results/target/tlsv1/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "tlsv10", "error": err})
	}

	eTlsv11, err := xpath.Compile("count(//document/results/target/tlsv1_1/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "tlsv11", "error": err})
	}

	eTlsv12, err := xpath.Compile("count(//document/results/target/tlsv1_2/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "tlsv12", "error": err})
	}

	eTlsv13, err := xpath.Compile("count(//document/results/target/tlsv1_3/acceptedCipherSuites/cipherSuite)")
	if err != nil {
		common.LogError("Error in creating xpath expression.", logrus.Fields{"expression": "tlsv13", "error": err})
	}

	scanFiles, err := filepath.Glob("./*_sslyze_current.xml")
	if err != nil {
		common.LogError("Error when trying to find scan results.", logrus.Fields{"error": err})
	}

	headLine := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", "host", "ip", "fingerprint", "serial", "notAfter", "sslv2", "sslv3", "tlsv10", "tlsv11", "tlsv12", "tlsv13", "heartBleed", "ccs", "sessionReneg", "robot")
	fCsv.WriteString(headLine)

	for _, sFile := range scanFiles {

		f, err := os.Open(sFile)
		if err != nil {
			common.LogError("Can't open result file.", logrus.Fields{"file": sFile, "error": err})
		}

		doc, err := xmlquery.Parse(f)
		if err != nil {
			common.LogError("Can't parse result file.", logrus.Fields{"file": sFile, "error": err})
		}

		host := xmlquery.FindOne(doc, "//document/results/target/@host").InnerText()
		ip := xmlquery.FindOne(doc, "//document/results/target/@ip").InnerText()

		leafCert := xmlquery.FindOne(doc, "//document/results/target/certinfo/receivedCertificateChain/certificate[1]")

		fingerprint := xmlquery.FindOne(leafCert, "@sha1Fingerprint").InnerText()
		serial := xmlquery.FindOne(leafCert, "serialNumber").InnerText()
		notAfter := xmlquery.FindOne(leafCert, "notAfter").InnerText()

		sslv2 := eSslv2.Evaluate(xmlquery.CreateXPathNavigator(doc))
		sslv3 := eSslv3.Evaluate(xmlquery.CreateXPathNavigator(doc))

		tlsv10 := eTlsv10.Evaluate(xmlquery.CreateXPathNavigator(doc))
		tlsv11 := eTlsv11.Evaluate(xmlquery.CreateXPathNavigator(doc))
		tlsv12 := eTlsv12.Evaluate(xmlquery.CreateXPathNavigator(doc))
		tlsv13 := eTlsv13.Evaluate(xmlquery.CreateXPathNavigator(doc))

		heartBleed := xmlquery.FindOne(doc, "//document/results/target/heartbleed/openSslHeartbleed/@isVulnerable").InnerText()
		ccs := xmlquery.FindOne(doc, "//document/results/target/openssl_ccs/openSslCcsInjection/@isVulnerable").InnerText()
		sessionReneg := xmlquery.FindOne(doc, "//document/results/target/reneg/sessionRenegotiation/@isSecure").InnerText()
		robot := xmlquery.FindOne(doc, "//document/results/target/robot/robotAttack/@resultEnum").InnerText()

		// TODO decide if needed
		/*
			sCC := ""
			certChain := xmlquery.Find(doc, "//document/results/target/certinfo/receivedCertificateChain/certificate")
			for i, cert := range certChain {

				sCC += xmlquery.FindOne(cert, "issuer").InnerText() + ":"
				sCC += xmlquery.FindOne(cert, "subject").InnerText() + ":"
				sCC += xmlquery.FindOne(cert, "serialNumber").InnerText() + ":"
				sCC += xmlquery.FindOne(cert, "notAfter").InnerText() + ":"
				sCC += xmlquery.FindOne(cert, "signatureAlgorithm").InnerText() + ":"

				if i < (len(certChain) - 1) {
					sCC += "\n"
				}
			}

			sCC = "\"" + sCC + "\""
		*/

		sLine := fmt.Sprintf("%v,%v,%v,\"%v\",\"%v\",%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", host, ip, fingerprint, serial, notAfter, sslv2, sslv3, tlsv10, tlsv11, tlsv12, tlsv13, heartBleed, ccs, sessionReneg, robot)

		fCsv.WriteString(sLine)
	}
}
