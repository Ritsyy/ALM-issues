# ALMighty

###Fetching A Specific Trello Board and displaying selected list's cards
(In this case, Fetching "Epic Backlog" list from "atomicopenshift-roadmap" Trello Board)

----
###How it works
1. Add yourself to [AtomicOpenShift Roadmap card] (https://trello.com/b/nlLwlKoz/atomicopenshift-roadmap).
2. Get the API key and token from (https://developers.trello.com/get-started/start-building).
3. Do
  1. **go build trello-alm.go**
  2. **go install**
  3. **trello-alm -apikey=//trello_API_key -token=//trello_token -username=//trello_username -BoardName=//By_default AtomicOpenShift Roadmap -ListName=//By_default Epic Backlog**
----
