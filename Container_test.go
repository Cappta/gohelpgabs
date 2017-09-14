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
				Convey("When popping Value", func() {
					path := "Value"
					value := container.PopPath(path)
					Convey("Then value should no longer exist in container", func() {
						So(container.ExistsP(path), ShouldBeFalse)
					})
					Convey("Then value's data should equal struct's value", func() {
						So(value.Data(), ShouldEqual, sampleStruct.Value)
					})
				})
				Convey("When popping NilValue", func() {
					path := "NilValue"
					value := container.PopPath(path)
					Convey("Then value should no longer exist in container", func() {
						So(container.ExistsP(path), ShouldBeFalse)
					})
					Convey("Then NilValue's data value should be nil", func() {
						So(value.Data(), ShouldBeNil)
					})
				})
				Convey("When popping MissingValue", func() {
					path := "MissingValue"
					value := container.PopPath(path)
					Convey("Then MissingValue should not exist in container", func() {
						So(container.ExistsP(path), ShouldBeFalse)
					})
					Convey("Then MissingValue should be nil", func() {
						So(value, ShouldBeNil)
					})
				})
				Convey("When searching Value", func() {
					path := "Value"
					value := container.Search(path)
					Convey("Then value's data should equal struct's value", func() {
						So(value.Data(), ShouldEqual, sampleStruct.Value)
					})
				})
				Convey("When searching NilValue", func() {
					path := "NilValue"
					value := container.Search(path)
					Convey("Then NilValue's data should be nil", func() {
						So(value.Data(), ShouldBeNil)
					})
				})
				Convey("When searching MissingValue", func() {
					path := "MissingValue"
					value := container.Search(path)
					Convey("Then MissingValue should be nil", func() {
						So(value, ShouldBeNil)
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
	Convey("When creating a new container", t, func() {
		container := New()
		Convey("Then container should not be nil", func() {
			So(container, ShouldNotBeNil)
		})
	})
}
