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

// add your ciphers here. Ciphers which are found on the server but not in this list are
// reported as policy violations
var ciphers = {
    ciphers: [
        'DHE-RSA-AES128-SHA256',
        'DHE-RSA-AES256-SHA256',
        'DHE-RSA-AES128-GCM-SHA256',
        'DHE-RSA-AES256-GCM-SHA384',
        'DHE-DSS-AES128-SHA256',
        'DHE-DSS-AES256-SHA256',
        'DHE-DSS-AES128-GCM-SHA256',
        'DHE-DSS-AES256-GCM-SHA384',
        'ECDH-RSA-AES128-SHA256',
        'ECDH-RSA-AES256-SHA384',
        'ECDH-RSA-AES128-GCM-SHA256',
        'ECDH-RSA-AES256-GCM-SHA384',
        'ECDH-ECDSA-AES128-SHA256',
        'ECDH-ECDSA-AES256-SHA384',
        'ECDH-ECDSA-AES128-GCM-SHA256',
        'ECDH-ECDSA-AES256-GCM-SHA384',
        'ECDHE-RSA-AES128-SHA256',
        'ECDHE-RSA-AES256-SHA384',
        'ECDHE-RSA-AES128-GCM-SHA256',
        'ECDHE-RSA-AES256-GCM-SHA384',
        'ECDHE-ECDSA-AES128-SHA256',
        'ECDHE-ECDSA-AES256-SHA384',
        'ECDHE-ECDSA-AES128-GCM-SHA256',
        'ECDHE-ECDSA-AES256-GCM-SHA384'
    ]
};

module.exports = {
    ciphers: ciphers.ciphers
};

