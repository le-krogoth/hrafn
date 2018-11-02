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
	"github.com/panjf2000/ants"
	"mzd.org.uk/hrafn/common"
	"mzd.org.uk/hrafn/data"
	"sync"
)

func main() {

	//
	common.InitConfig()

	start()

	GenerateReport()
}

func start() {

	defer ants.Release()

	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(100, func(rec interface{}) error {
		Scan(rec)
		wg.Done()
		return nil
	})

	defer p.Release()

	data.LoadDomains()

	// submit tasks
	for i := 0; i < data.GetDomainCount(); i++ {

		// todo handle error
		record, err := data.GetDomainRecord(i)
		if err != nil {

		}

		wg.Add(1)
		p.Serve(record)
	}
	wg.Wait()
	// fmt.Printf("running goroutines: %d\n", p.Running())
}
