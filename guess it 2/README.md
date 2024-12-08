## Guess it 2


### Overview
This program implements a predictive algorithm that attempts to guess ranges for a sequence of numbers. The key feature is its stability - rather than trying to predict exact values, it maintains a stable "level" that adjusts gradually to changes in the input values.


### How It Works
#### Core Concepts

1. Current Level

- The algorithm maintains a "current level" which serves as the center point for predictions
- This level moves slowly towards new values to maintain stability
- It only adjusts 15% of the way towards each new value


2. Prediction Ranges

- Each prediction is a range of exactly 12 units (±6 from the current level)
- Example: if current level is 100, the prediction range will be 94 to 106


3. Gradual Adjustment

- When a new value arrives, the algorithm calculates:

    - The difference between the new value and current level
    - Moves 15% of that difference


- Example:

    - Current level: 100
    - New value: 200
    - Difference: 100
    - Adjustment: 15 (15% of 100)
    - New level: 115

### Key Features

1. Stability

- Resistant to sudden spikes or drops
- Maintains consistent predictions
- Prevents overreaction to outliers


2. Fixed Range Width

- Always uses exactly 12 units (±6)
- Provides consistent coverage


3. Simple State

- Only needs to track current level
- No complex calculations or historical data required


4. Gradual Adaptation

- 15% adjustment rate prevents wild swings
- Still responds to genuine trends over time

### Test the program 


1. Check if `script.sh` has executable permissions after cloning. Sometimes these permissions can be changed when cloned if the environment is different than the one when the file was pushed to repository. 
    If the file has executable permissions, procceed to next steps. Otherwise give the permission using:
```
chmod +x script.sh
```

2. Run the test by downloading this [zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) containing the tester. You should place the student/ folder in the root directory of the items provided.

3. Verify that Docker Desktop is running.

4. These commands should be ran (on the root directory of files downloaded) to have the dependencies needed and to start the webpage on the port 3000:
```
docker compose up --build
```

5. After opening your browser of preference in the port
[3000](http://localhost:3000/), if you try clicking on any of the `Test Data`
buttons, you will notice that in the Dev Tool/ Console there is a message which
tells you that you need another guesser besides the student.

6. Adding a guesser is simple. You need to add in the URL a guesser, in other
words, the name of one of the files present in the `ai/` folder:
```
?guesser=<name_of_guesser>
```

For example:
```
?guesser=big-range
```

7. After that, choose which of the random data set to test. After that you can
wait for the program to test all of the values (boooooring), or you can click
`Quick` to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays
clicking on the `Clean` button after each test.

You will need to test `linear-regr` , `big-range` and `correlation-coef`.
For each of these you will need to test `Test4` and `Test5` 3 times each.
If student program wins at least 2/3 every one of the test, they pass the audit.


### Study

- [Linear Regression](https://en.wikipedia.org/wiki/Linear_regression)

- [Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)

----------

#### Authors 
- _ychaniot (Ypatios Chaniotakos)_

