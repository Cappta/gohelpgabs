package gohelpgabs

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type SampleStruct struct {
	NilValue *string
	Value    string
}

func TestContainer(t *testing.T) {
	Convey("Given a sample struct", t, func() {
		sampleStruct := &SampleStruct{Value: "Value"}
		sampleStructJSON := "{\"NilValue\":null,\"Value\":\"Value\"}"

		Convey("When marshalling to JSON", func() {
			jsonData, err := json.Marshal(sampleStruct)
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then json should not be nil", func() {
				So(jsonData, ShouldNotBeNil)
			})
			Convey("When converting json data to string", func() {
				jsonString := string(jsonData)
				Convey("Then json should equal expected output", func() {
					So(jsonString, ShouldEqual, sampleStructJSON)
				})
			})
			Convey("When parsing nil", func() {
				container, err := ParseJSON(nil)
				Convey("Then error should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
				Convey("Then container should be nil", func() {
					So(container, ShouldBeNil)
				})
			})
			Convey("When parsing container", func() {
				container, err := ParseJSON(jsonData)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
				Convey("When marshalling container to json string", func() {
					containerString := container.String()
					Convey("Then json string should equal expected output", func() {
						So(containerString, ShouldEqual, sampleStructJSON)
					})
				})
				Convey("When getting missing paths", func() {
					missingPaths := container.GetMissingPaths("Value", "NilValue", "MissingPath")
					Convey("Then missing paths should resemble expected output", func() {
						expectedOutput := []string{"MissingPath"}
						So(missingPaths, ShouldResemble, expectedOutput)
					})
				})
				Convey("When setting Value if path exists", func() {
					path := "Value"
					value := "NewValue"
					container.SetValueIfPathExists(path, value)
					Convey("Then path's value should equal set value", func() {
						So(container.Path(path).Data(), ShouldEqual, value)
					})
				})
				Convey("When setting NilValue if path exists", func() {
					path := "NilValue"
					value := "NewValue"
					container.SetValueIfPathExists(path, value)
					Convey("Then path's value should equal set value", func() {
						So(container.Path(path).Data(), ShouldEqual, value)
					})
				})
				Convey("When setting MissingValue if path exists", func() {
					path := "MissingValue"
					value := "NewValue"
					container.SetValueIfPathExists(path, value)
					Convey("Then path's value should be nil", func() {
						So(container.Path(path), ShouldBeNil)
					})
				})
			})
		})
	})
}
