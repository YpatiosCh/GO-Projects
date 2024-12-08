
# Guess it 1

## Instructions

You must build a program that given a number as standard input, prints out a range in which the next number provided should be.

The data received by the program, as always, will be presented as the following example:
```
189
113
121
114
145
110
...
```

## Testing

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

You will need to test `Median` , `Big-Range` and `Average`.
For each of these you will need to test `Test1`, `Test2` and `Test3` 3 times each.
If student program wins at least 2/3 every one of the test, they pass the audit.

## Authors 

- ychaniot (Ypatios Chaniotakos)
