#!/usr/bin/env node

/*
====================================================================

hrafn

--------------------------------------------------------------------
Copyright (c) 2015 by Krogoth of
Ministry of Zombie Defense - http://www.mzd.org.uk/

This file is part of hrafn.

hrafn is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
 ====================================================================
*/

var config = require('./config').config;
var program = require('commander');
var chalk = require('chalk');

console.log(chalk.red("Init..."));

program
    .version('0.0.1')
    .option('-i, --scanfiles <path>', 'Path to folder with scanfiles [config.folder_scans]', config.folder_scans)
    .option('-o, --outfile <file>', 'Filename to output csv to [config.file_output]', config.file_output)
    .option('-s, --separator <char>', 'Separator to be used for list fields in CSV [config.csv_separator]', config.csv_separator)

program
    .parse(process.argv);


generateReport();

function genCsv(err, csv)
{
    if (err)
    {
        throw err;
    }

    console.log(chalk.red("Writing files"));

    var fs = require('fs');
    fs.writeFile(program.outfile, csv, 'utf8', function(err) {
        if(err) {
            console.log(chalk.red("%s"), err);
        } else {
            console.log(chalk.green("CSV file '%s' was created"), program.outfile);
        }
    });
};

function checkAllowedCiphers(usc, ciphersAllowed, ciphersFound)
{
    var ciphLen = ciphersFound.length;
    var strResult = "";
    for (var ciphCount = 0; ciphCount < ciphLen; ciphCount++)
    {
        var cipherName = ciphersFound[ciphCount].value;
        if (!usc.contains(ciphersAllowed.ciphers, cipherName))
        {
            if(strResult.length > 0)
            {
                strResult += program.separator;
            }
            strResult += cipherName;
        }
    }

    return strResult;
}


function generateReport()
{
    if (!program.scanfiles || !program.outfile)
    {
        console.log("  error: option '-s, --scanfiles <path>' argument missing");
        console.log("  error: option '-o, --outfile <path>' argument missing");
        return;
    }

    var glob = require("glob");

    var fs = require('fs');

    var xpath = require('xpath');
    var dom = require('xmldom').DOMParser;

    var converter = require('json-2-csv');

    var usc = require('underscore');
    var ciphers = require('./ciphers.js').ciphers;

    console.log(chalk.red("Starting processing export"));

    glob(program.scanfiles + "/*sslyze_current.xml", function (err, files) {

        var lines = [];

        var length = files.length;
        for (var i = 0; i < length; i++)
        {
            console.log(chalk.green("Processing: %s"), files[i]);

            var file = files[i];

            try {
                var xml = fs.readFileSync(file, 'utf8');
                var doc = new dom().parseFromString(xml);

                //
                var host = xpath.select1("//document/results/target/@host", doc).value;
                var ip = xpath.select1("//document/results/target/@ip", doc).value;
                var fingerprint = xpath.select1("//document/results/target/certinfo_basic/receivedCertificateChain/certificate[@position='leaf']/@sha1Fingerprint", doc).value;

                var sslv2 = xpath.select1("count(//document/results/target/sslv2/acceptedCipherSuites/cipherSuite)", doc);
                var sslv3 = xpath.select1("count(//document/results/target/sslv3/acceptedCipherSuites/cipherSuite)", doc);

                var ciphtlsv10 = xpath.select("//document/results/target/tlsv1/acceptedCipherSuites/cipherSuite/@name", doc);
                var tlsv10 = ciphtlsv10.length;

                var ciphtlsv11 = xpath.select("//document/results/target/tlsv1_1/acceptedCipherSuites/cipherSuite/@name", doc);
                var tlsv11 = ciphtlsv11.length;

                var ciphtlsv12 = xpath.select("//document/results/target/tlsv1_2/acceptedCipherSuites/cipherSuite/@name", doc);
                var tlsv12 = ciphtlsv12.length;

                var arrCertChain = xpath.select("//document/results/target/certinfo_basic/receivedCertificateChain/certificate", doc);
                var iCCLength = arrCertChain.length;

                var strCC = "";

                for(var icc = 0; icc < iCCLength; icc++)
                {
                    var cert = arrCertChain[icc];
                    strCC += xpath.select1("@position", cert).value + ":";
                    strCC += xpath.select1("issuer/commonName/text()", cert) + ":";
                    strCC += xpath.select1("subject/commonName/text()", cert) + ":";
                    strCC += xpath.select1("serialNumber/text()", cert) + ":";
                    strCC += xpath.select1("validity/notAfter/text()", cert) + ":";
                    strCC += xpath.select1("signatureAlgorithm/text()", cert);

                    if(icc < iCCLength -1)
                    {
                        strCC += ";\n";
                    }
                }

                strCC = "\"" + strCC + "\"";

                lines.push({
                    host: host,
                    ip: ip,
                    fingerprint: fingerprint,
                    certchain: strCC,
                    sslv2: sslv2,
                    sslv3: sslv3,
                    tlsv1_0: tlsv10,
                    tlsv1_1: tlsv11,
                    tlsv1_2: tlsv12,
                    policy_violation_tlsv1_0: checkAllowedCiphers(usc, ciphers, ciphtlsv10),
                    policy_violation_tlsv1_1: checkAllowedCiphers(usc, ciphers, ciphtlsv11),
                    policy_violation_tlsv1_2: checkAllowedCiphers(usc, ciphers, ciphtlsv12)
                });

                //console.log("File '" + file + "/ was successfully read.");
            } catch (ex) {
                console.log("Unable to read file '%s'.", file);
                //console.log(ex);
            }
        }

        console.log(chalk.red("Formatting CSV"));
        converter.json2csv(lines, genCsv);
    });

}