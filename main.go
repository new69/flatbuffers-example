package main

import (
	"fmt"
	"test/School"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
    builder := flatbuffers.NewBuilder(1024)

    courseName := builder.CreateString("Go Programming")
    courseDescription := builder.CreateString("Learn Go Programming")

    School.CourseStart(builder)
    School.CourseAddId(builder, 1)
    School.CourseAddName(builder, courseName)
    School.CourseAddDescription(builder, courseDescription)
    course := School.CourseEnd(builder)
    builder.Finish(course)
    
    buf := builder.FinishedBytes()
    c := School.GetRootAsCourse(buf, 0)
    fmt.Println(c.Id())
    fmt.Println(string(c.Name()))
    fmt.Println(string(c.Description()))
}