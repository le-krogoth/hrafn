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
var spawn = require('child_process').spawn;
var fs = require('fs-extra');
var async = require('async');

console.log(chalk.red("Init..."));

program
    .version('0.0.1')
    .option('-o, --out <path>', 'Path to output scanresults to [config.folder_scans]', config.folder_scans)
    .option('-s, --sslyze <path>', 'Path to the sslyze tool [config.sslyze]', config.sslyze)
    .option('-n, --nmap <path>', 'Path to the nmap tool [config.nmap]', config.nmap)

program
    .parse(process.argv);


runScan();

function scanDomain(domain, fdate, callback)
{
    runSslyze(domain, fdate);
    runNmap(domain, fdate, callback);
}

function runSslyze(domain, fdate)
{
    console.log(chalk.green("Executing sslyze on domain %s"), domain);

    var outFile = program.out + "/" + domain + "_sslyze_current.xml";

    var scan = spawn(program.sslyze, ['--regular', domain, '--sni=' + domain, '--xml_out=' + outFile]);
    scan.stdout.on('data', function (data) {
        //console.log('stdout: ' + data);
    });

    scan.stderr.on('data', function (data) {
        console.log('stderr: ' + data);
    });

    scan.on('close', function (code) {
        console.log("sslyze exited for domain '%s' with code '%s'", domain, code);

        if(code == 0)
        {
            fs.copy(outFile, program.out + "/" + fdate + "/" + domain + "_sslyze_" + fdate + ".xml", function (err) {
                if (err) return console.error(err)
            });
        }
    });
}

function runNmap(domain, fdate, callback)
{
    console.log(chalk.green("Executing nmap on domain %s"), domain);

    var outFile = program.out + "/" + domain + "_nmap_current.xml";

    var scan = spawn(program.nmap, ['-sV', '-v', '-Pn', '-p http,https', '--script=http-headers', domain, '-oX=' + outFile]);
    scan.stdout.on('data', function (data) {
        //console.log('stdout: ' + data);
    });

    scan.stderr.on('data', function (data) {
        console.log('stderr: ' + data);
    });

    scan.on('close', function (code) {
        console.log("nmap exited for domain '%s' with code '%s'", domain, code);

        if(code == 0)
        {
            fs.copy(outFile, program.out + "/" + fdate + "/" + domain + "_nmap_" + fdate + ".xml", function (err) {
                if (err) return console.error(err)
            });
        }

        callback();
    });
}

function runScan()
{
    if (!program.out)
    {
        console.log("  error: option '-o, --out <path>' argument missing");
        return;
    }

    var glob = require("glob");

    var usc = require('underscore');
    var domains = require('./domains.js').domains;

    var moment = require('moment');
    var fdate = moment().format("YYYY-MM-DD-hh-mm");

    var outPath = program.out + "/" + fdate;
    fs.ensureDirSync(outPath, function(err) {
        console.log(err); // => null
        //dir has now been created, including the directory it is to be placed in
    })

    //var lenDomains = domains.length;
    //for (var countDomain = 0; countDomain < lenDomains; countDomain++)
    //{
    //    var domain = domains[countDomain];
    //    scanDomain(domain, fdate);
    //}

    async.forEachLimit(domains, config.job_limit, function(domain, callback) {
        scanDomain(domain, fdate, callback);
    }, function(err) {
        if (err) {
            console.log("Error in foreachlimit: %s", err);
        }
    });

}