## My Own `wc` Tool

This is my first completion of the challenge from [coding challenges](https://codingchallenges.fyi)

You can view [here](https://codingchallenges.fyi/challenges/challenge-wc/) for more details on the challenge stages and requirements.

These are the features of the `wc` tool:

1. It takes the flag `-c` to calculate the number of bytes of a file
```shell
> ./ccwc -c test.txt
  342190 test.txt
```
2. `-l` outputs the number of lines in a file
```shell
> ./ccwc -l test.txt
    7145 test.txt
```
3. `-w` outputs the number of words in a file
```shell
> ./ccwc -w test.txt
   58164 test.txt
```
4. `-m` outputs the number of characters in a file
```shell
> ./ccwc -m test.txt
  339292 test.txt
```
5. the default option - i.e. no options are provided, which is the equivalent to the `-c`, `-l` and `-w` options
```shell
> ./ccwc test.txt
    7145   58164  342190 test.txt
```
6. reads from standard input if no filename is specified
```shell
> cat test.txt | ./ccwc -l
    7145
```

> A text file has been added `test.txt` to test the `wc` tool.

If for whatever reason you want to use my clone instead of the provided `wc` tool in your shell 😂 or want to make changes, you can clone by:

```shell
git clone https://github.com/zub-bee/build-wc.git
cd build-wc
```