# ALMighty

###Fetching A Specific Trello Board and displaying selected list's cards
(In this case, Fetching "Epic Backlog" list from "atomicopenshift-roadmap" Trello Board)

----
###How it works
1. Add yourself to [AtomicOpenShift Roadmap card] (https://trello.com/b/nlLwlKoz/atomicopenshift-roadmap).
2. Get the API key and token from (https://developers.trello.com/get-started/start-building).
3. go build trello-alm.go
4. go install
5. trello-alm -apikey=your_api_key -token=your_token -username=trello_username -BoardName=board_name  -ListName=list_name
6. default BoardName is AtomicOpenShift Roadmap and list name is Epic Backlog

----

###Fetching Issues from Github
(In this case, Fetching the query: is:open is:issue user:arquillian author:aslakknutsen)

----

###How it works
1. go install
2. github-alm

----

###Fetching Issues from Github and Trello

----

###How it works
1. go install
3. for trello
  * goalm -tool=trello -apikey=your_api_key -token=your_token -username=trello_username -BoardName=board_name -ListName=list_name
  * default BoardName is AtomicOpenShift Roadmap and list name is Epic Backlog
4. for github
  * goalm -tool=github -query=your_search_query
  * default query is "query: is:open is:issue user:arquillian author:aslakknutsen"
5. for open trello boards you can access as:
  * goalm -tool=trello -username=trello_username -BoardName=board_name

----
