package cmd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOverThreshold(t *testing.T) {
	var num1 int64
	var num2 int64
	var alarm bool

	Convey("When the threshold is exceeded", t, func() {
		num1 = 20
		num2 = 2
		alarm = overThreshold(num1, num2)

		Convey("20 should be greater than 2 and the alarm should be true", func() {
			So(alarm, ShouldBeTrue)
		})
	})

	Convey("When the threshold is not exceeded", t, func() {
		num1 = 2
		num2 = 20
		alarm = overThreshold(num1, num2)

		Convey("2 should be less than 20 and the alarm should not be true", func() {
			So(alarm, ShouldBeFalse)
		})
	})
}

func TestCreateMap(t *testing.T) {

	Convey("When creating a map", t, func() {
		var m = make(map[string]int64)

		Convey("When reading /proc/meminfo", func() {

			Convey("The value of 'AnonPages' should not be empty", func() {
				So(m["AnonPages"], ShouldBeGreaterThanOrEqualTo, 0)
			})

			Convey("The value of 'MemFree' should not be empty", func() {
				So(m["Memfree"], ShouldBeGreaterThanOrEqualTo, 0)
			})

			Convey("The value of 'Slab' should not be empty", func() {
				So(m["Slab"], ShouldBeGreaterThanOrEqualTo, 0)
			})

			Convey("The value of 'AnonPages' should not be ", func() {
				So(m["Active(file)"], ShouldBeGreaterThanOrEqualTo, 0)
			})
		})
	})
}
