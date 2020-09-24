# Sports-parser

This web-application can displays news from the sports.ru site using a tag that helps you to find relevant information about the club, player, or league.
It can be easier to recognize the last news without using this site to search for interesting information. Or it cant be :)

Anyway, I think this work will be very useful for me: I assume understanding work with code more closely. And current web-application is useful for me :)


## Future improvements:
- add search for several tags;
- add a log for good or bad work;
- add a log for search history;
- do some refactoring;
- realize search not only on one page.

## Current versions:

### 24.09.2020:
Add:
- log.txt.

### 14.09.2020:
Add:
- history.txt that saves search history;
- fmt.Scanln for pause before exit;
- work with files.

Also redesigned the structure of the application.

### 12.09.2020:
Add:
- information about good or bad results of the search;
- foolproofs for input tag;
- fixed the issue related to len of the input tag on windows.

### 11.09.2020:
Web-application can:
- read the tag from standard input;
- search relevant news from the site using input tag;
- write to the standard output date, news, and link to the news.

Issues:
- web-application work without problems on Linux, but on windows, there is a problem: the length of the input tag and length of the site's tag don't match.
