package main;

import "fmt";
import (
	"github.com/twinj/uuid"
	"os"
);

func main() { os.Exit(mainReturnWithCode()) }


func mainReturnWithCode() (returnCode int){
	fmt.Printf("hello, world\n");
	var err error;

	var teachers []unsharableResource;
	teachers, err = getTeacherInfo();

	fmt.Printf("teachers: %v", teachers);



	if (err!=nil) {return -1;};

	return 0;

	//Brute force plan
	//load the course alternatives
	//load the student preferences
	//load  (hard code) the structural restrictions (parallel courses, non preferred slots, etc.)
	//generate (viable) schedules  - viable includes teacher in the same period, may include rules like "core" courses not in conflict
	//score schedules against student preferences,
	//report preferred courses


}

func getTeacherInfo() ([]unsharableResource, error) {
	var err error;
	var teachers []unsharableResource = make([]unsharableResource, 0);
	var teacher unsharableResource;

	teacher.resourceUuid, err =uuid.Parse("e5d19622-b81c-4cc5-bf66-f4a041d5ffbd");
	teacher.displayName="Alice";
	teachers = append(teachers, teacher);
	teacher.resourceUuid, err = uuid.Parse ("a4209c52-839f-4e04-ad1c-8a4d98de1ae4");
	teacher.displayName="Bob";
	teachers = append(teachers, teacher);

	if( err != nil) {return nil, err;}  //not really checking all the errors, just the last one....
	return teachers, nil;

}



//Some Tasks
//   - Model course, preference, restrictions, schedule
//   - firm up permutations of schedules based on assumed sizes



//Valid Schedule:

//1) Does not require a teacher to be in two places at once
//2) Does not require a class to be occupying the same space (and some "extra" spaces may be less desirable)
//3) Meets course hour requirements per cycle, i.e Does not exceed the number of slices available for each cycle
// 			(e.g. if a course requires a full day each 5 day week, then there are only 5 independent slots)

//A unique valid schedule:

// 1) A Valid Schedule where all the isomophic schedules are only counted as one
// 		Example: if there are only two slots A and B and courses C, D and E:  C in A and E in B is isomorphic with:
// 																			  E in A and C in B


