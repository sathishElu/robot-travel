# robot-travel

### Run the below commands to build, test and run

**To build**
  - go build

**To test**
  - go test -v -tags=unitTest -count=1 ./...

**To run**
  - go run main.go

### INPUT to the application:
  - The first line of input is the top-right coordinates of the rectangular plan (M, N),
  - Next two lines of input are about the robot
    - first line gives the robot's starting position, X,Y coordinates and a letter, all 3 separated by spaces
    - second line is a series of commands for the robot

### OUTPUT of the application:
  Robot's final coordinates and direction.

### Sample Input and Output

Test Input <br />
4 4 <br />
0 0 N <br />
MMMRMMLM <br />

Test Output <br />
2 4 N <br />
