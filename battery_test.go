// battery
// Copyright (C) 2016 Karol 'Kenji Takahashi' Woźniak
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package battery

import (
	"fmt"

	"github.com/distatus/battery"
)

func ExampleGetAll() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get batteries info")
		return
	}

	for i, battery := range batteries {
		fmt.Printf("Bat%d: ", i)
		fmt.Printf("state: %f, ", battery.State)
		fmt.Printf("current capacity: %f mWh, ", battery.Current)
		fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		fmt.Printf("design capacity: %f mWh, ", battery.Design)
		fmt.Printf("charge rate: %f mW\n", battery.ChargeRate)
	}
}

func ExampleGetAll_errors() {
	batteries, err := battery.GetAll()
	if err == nil {
		fmt.Println("Got batteries info")
		return
	}

	switch perr := err.(type) {
	case ErrFatal:
		fmt.Println("Fatal error! No info retrieved")
	case Errors:
		for i, err := range perr {
			if err != nil {
				fmt.Printf("Could not get battery info for `%d`\n", i)
				continue
			}

			fmt.Printf("Got battery info for `%d`\n", i)
		}
	}
}

func ExampleGet() {
	battery, err := battery.Get(0)
	if err != nil {
		fmt.Println("Could not get battery info")
		return
	}

	fmt.Printf("Bat%d: ", i)
	fmt.Printf("state: %f, ", battery.State)
	fmt.Printf("current capacity: %f mWh, ", battery.Current)
	fmt.Printf("last full capacity: %f mWh, ", battery.Full)
	fmt.Printf("design capacity: %f mWh, ", battery.Design)
	fmt.Printf("charge rate: %f mW\n", battery.ChargeRate)
}

func ExampleGet_errors() {
	battery, err := battery.Get(0)
	if err == nil {
		fmt.Println("Got battery info")
		return
	}

	switch perr := err.(type) {
	case ErrFatal:
		fmt.Println("Fatal error! No info retrieved")
	case ErrPartial:
		if perr.Current != nil {
			fmt.Println("Could not get current battery capacity")
			return
		}

		fmt.Println("Got current battery capacity")
	}
}
