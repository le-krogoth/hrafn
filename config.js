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

var config =
{
    // separator used in fields where there are multiple items, like the policy violation fields
    csv_separator: ":",

    // path (and name) of used tools
    sslyze: "../sslyze/sslyze.py",
    nmap: "nmap",

    // folder to put all scan results into
    folder_scans: "scans",

    // path and name of file to write report into
    file_output: "results_current.csv"
};

module.exports = {
    config: config
};