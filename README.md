hrafn
=====

**hrafn** (c) 2014-15 by [Krogoth](https://twitter.com/le_krogoth) of [Ministry of Zombie Defense](http://www.mzd.org.uk/)

## Introduction ##

**hrafn** consists of a set of scripts to regularly scan and report on your own hosts for policy violations, mostly in SSL/TLS.
At least TLS is what **hrafn** was written for, but you could add your own checks easily.

And just in case you wonder, **hrafn** means raven in Old Norse (https://en.wikipedia.org/wiki/Hrafn).


## Prerequisites ##

Right now, **hrafn** needs an installation of *sslyze* and of *nmap* for its scans. As well as an installation of nodejs to run the scripts.

* Get sslyze from here: https://github.com/nabla-c0d3/sslyze/releases
* Get nodejs and nmap through the package manager of your distribution / OS.


## Installation ##

The installation is quite straightforward:

* Clone this git repository to your local machine with:

```
git clone https://github.com/le-krogoth/hrafn.git
```

* Run the following command in the root directory to install all dependencies.

```
npm install
```

## Configuration ##

* Change the settings in the config.js file to your liking.
* Add your domains to the domains.js file. Please make sure that you are allowed to scan these domains.
* Configure ciphers.js to your liking. This file contains all the ciphers which your policy allows. See report section for details.

## Run ##

* Run the scan process like this. If no parameter is given, the scan script takes the configuration from the config file.

```
./scan.js

// or when using command line configurations

./scan.js --out foldertostoreresultsin --sslyze pathtosslyze --nmap pathtonmap

// run scan.js with the help flag to learn more about the commandline

./scan.js --help
```
* Run the report job to generate a CSV file.

```
./report.js

// run again with the help flag for details regarding the commandline. Again, when no
// parameter is given, the script takes the configuration from the config file.
```

If you want to run the scan as well as generate the report regularly, you could add these lines to your crontab file
as root. *Just don't forget to change the hrafnuser and your path accordingly*.

```
07 8    * * *   hrafnuser   cd /path/to/hrafn && scan.js
53 8    * * *   hrafnuser   cd /path/to/hrafn && report.js
```


## Report ##

The generated report is in CSV format (to be imported in some tool like, say, Splunk) and contains these fields:

* **host**: Scanned host
* **ip**: IP address of scanned host
* **fingerprint**: Fingerprint of certificate found on host
* **sslv2**: Amount of supported ciphers with this protocol version
* **sslv3**: Amount of supported ciphers with this protocol version
* **tlsv1_0**: Amount of supported ciphers with this protocol version
* **tlsv1_1**: Amount of supported ciphers with this protocol version
* **tlsv1_2**: Amount of supported ciphers with this protocol version
* **policy_violation_tlsv1_0**: This field contains all ciphers which are not in your ciphers.js but were supported on this protocol on the server.
* **policy_violation_tlsv1_1**: This field contains all ciphers which are not in your ciphers.js but were supported on this protocol on the server.
* **policy_violation_tlsv1_2**: This field contains all ciphers which are not in your ciphers.js but were supported on this protocol on the server.

