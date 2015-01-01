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

// this file contains all domains which are checked during one scan.
//
// IMPORTANT:
// Please make sure that you are allowed to scan all of the listed domains!

var domains = {
    domains: [
        'www.domain1.com',
        'www.domain2.com'
    ]
};

module.exports = {
    domains: domains.domains
};

