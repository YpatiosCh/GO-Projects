## Linear-Stats

### Description

This project calculates the Linear Regression Line and Pearson Correlation Coefficient from input data files containing numerical values.

Requirements:

- Read numerical data from a text file
- Calculate Linear Regression Line
- Calculate Pearson Correlation Coefficient
- Print results in a specific format

### Test the program

1. Download the file provided in this [link](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip).

2. Extract the files.

3. Copy bin folder, run.sh script and the Dockerfile provided.

4. Paste in root directory of user program.

5. To test the user program, the auditor will have to run the command:
```
./bin/linear-stats
```
This command will run the script provided. It will generate a data.txt file and the result from the tester program that the auditor will need to compare with the user's program.

6. Now that the result of the tester program and data.txt file are created, it's time to test the user's program. To do so, the the auditor will have to run the command:
```
go run main.go data.txt
```
or
```
go run . data.txt
```

7. Compare the result from the user program with the result of the tester program. If they match, the project succeeds.


### Learning Objectives

- Statistical calculations
- Probability analysis
- Data processing
- Linear regression techniques


### Study

- [Linear Regression](https://en.wikipedia.org/wiki/Linear_regression)

- [Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)

----------

#### Authors 
- _ychaniot (Ypatios Chaniotakos)_

