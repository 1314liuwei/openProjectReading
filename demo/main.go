package main

import "github.com/gogf/gf/v2/os/gres"

type Demo struct {
	Name string
	Age  int
}

func main() {

	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GB4zRkawoAE+Bg4GdLz9XLzU/T0IXRoCCsD4+rnQklTJvtOm20g0HL9o2aEGvff5sbcdROEOa5wsC785tFkaFkr923Pz1oLfemEwPjyUM8zGWL8rY2S/urdrQsW2c1VstHVc0qMlZve983CuLssoSxhtU6Yp+VTlW+ffH/+/fJ27o/nQs7/ti65uLFns36Yn7cW1zEu8Z7uC8/5P8ldcnVPuP56xeUfm5wCKnl++LU/aEr9/PV7975FVxfeu7Fj6UnDbhUngw/pR1XfrvsmM8n7nPTKHRrZoZPLVliKu3nE8rYzs8+9e6Lqi0R3aMBCDftryxpu8t5+fdwvhP/Vjp0meht/F+d2bV3odOGT4MR5NoYX1wrc3LNUfb++pX1BqqHVe9PqNr1Ee8d7jZ8cn5s92rJc4vMbz1MXYz2umQl2f1fq/3zG6vmG6j1fl1rzWytm/M4/pDr387dNDyz+r7bTY5w39+RjWct7DXLLUg8ZKJ7evfj5C6Mj1l9/yike5W7hkUs7UbiYTzJth+mZ/SY1Cmknff6zP1lRbeGwh29a0pY3pw7VFa0Xj+RQfPIvrGtOdEDuct/cX+eX/6pzuHXvP2eT+Z6pQdNCPX2K15WubNsX+iHBoe312aZFvA+/5DWFB51w81Ka5Te15kW2bY1h4Uyd0sbcqm8loUdX3fPhiOsO2xY0aZNteOrS813X77Yz2vvVMzMw/P8f4M3OcXzviuWPGBkY9rAwMMCSCQODUVAESjJhgycTcPJwjBZPAmlGVhLgzcgkwoxIZcgGg1IZDGxrBJE40hzCEOxugAABhv+O6kwMmC5iZQNJMzEwMTQxMDCkMIF4gAAAAP//bDPbUPwCAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
	gres.Dump()

}
