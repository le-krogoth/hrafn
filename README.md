hrafn
=====

**hrafn** (c) 2014-18 by [Krogoth](https://twitter.com/le_krogoth) of [Ministry of Zombie Defense](http://www.mzd.org.uk/)

## Introduction ##

**hrafn** scans your domains and reports policy violations as well as security problems with your SSL/TLS configuration.

At least TLS scans are what **hrafn** was written for, you could easily add your own checks to that.

And just in case you wonder, **hrafn** means raven in Old Norse (https://en.wikipedia.org/wiki/Hrafn).


## Prerequisites ##

**hrafn** used to need an installation of nodejs to run. Nodejs is not needed anymore. **hrafn** comes in binary form nowadays and can be run directly on your scan host.

**hrafn** needs an installation of *sslyze* and of *nmap* for its scans. 

* Get sslyze from here: https://github.com/nabla-c0d3/sslyze/releases or install it via pip (better) or through the package manager of your OS.
* Get nmap through the package manager of your distribution / OS.

## Installation ##

Get hrafn from our [release page](https://github.com/le-krogoth/hrafn/releases).

If you prefer to build your own copy, clone this git repository to your local machine with:

```
git clone https://github.com/le-krogoth/hrafn.git
```

You will need a go compiler to compile hrafn.


## Configuration ##

### Settings ###
Change the settings in the hrafn.config.js file to your liking. **hrafn** will generate a file for you if it does not detect one.

### Domains ###
There should be a domains.csv file. Add the IP addresses or domains to be scanned to this file. The format of the file is quite simple and consists of these elements:

```
domain,tls_scan,nmap_scan
```

- domain: The domain or IP to be scanned. Please make sure that you are allowed to scan these domains.
- tls_scan: 0 or 1 if the domain should be run through sslyze
- nmap_scan: 0 or 1 if the domain should be run through nmap


### Ciphers ###
Configure ciphers.csv to your liking. This file contains all the ciphers which your policy allows. See report section for details.

## Run ##

* Run the scan process like this. If no parameter is given, the scan takes the configuration from the config file.

```
hrafn scan
```

* Run the report job to generate a CSV file.

```
hrafn report
```

* If you want to run both jobs, use this:

```
hrafn full
```


If you want to run the scan as well as generate the report regularly, you could add this line to your crontab file as root. 

*Just don't forget to change the hrafnuser and your path accordingly*.

```
07 8    * * *   hrafnuser   cd /path/to/hrafn && hrafn full
```


## Report ##

The generated report is in CSV format (to be imported in some tool like, say, Splunk) and contains these fields:

* **host**: Scanned host
* **ip**: IP address of scanned host
* **fingerprint**: Fingerprint of certificate found on host
* **serial**: Serial number of the leaf certificate
* **notAfter**: Expiration date of the leaf certificate
* **sslv2**: Amount of supported ciphers with this protocol version
* **sslv3**: Amount of supported ciphers with this protocol version
* **tlsv1_0**: Amount of supported ciphers with this protocol version
* **tlsv1_1**: Amount of supported ciphers with this protocol version
* **tlsv1_2**: Amount of supported ciphers with this protocol version
* **tlsv13**: Amount of supported ciphers with this protocol version
* **heartBleed**: Is this installation vulnerable to Heartbleed?
* **ccs**: Is this installation vulnerable to the OpenSSL CCS Injection?
* **sessionReneg**: Is this installation vulnerable to Session Renegotiation?
* **robot**: Is this installation vulnerable to ROBOT attack?
* **policy_violation_tlsv1_0**: This field contains all ciphers which are not in your ciphers.csv but were supported on this protocol on the server.
* **policy_violation_tlsv1_1**: This field contains all ciphers which are not in your ciphers.csv but were supported on this protocol on the server.
* **policy_violation_tlsv1_2**: This field contains all ciphers which are not in your ciphers.csv but were supported on this protocol on the server.
