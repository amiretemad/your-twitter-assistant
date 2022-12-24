Sets of command to fetch your existence twitter followers, new and who just unfollowed you! I use this commands just to track my
followers on Twitter and have some plans to add some more functionality to it!

## Build it!

```
go build cmd/console/main.go
```

## Define required ENV variables

```
# Database configuration
DATABASE_NAME=tweak_twitter

# Twitter credentials
CONSUMER_KEY=NXXXXXXXXXXXXXXXXXF
CONSUMER_SECRET=ZXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXR
TOKEN=1XXXXXX6-Z1XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX4X
TOKEN_SECRET=UrXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXOy
```

## How to Execute it

```
./main fetch_followers --username=someusername
```

## Available commands

You need to provider --username option in commands to fetch followers of provided user.

```
COMMANDS:
   fetch_followers    Fetch followers for provided username.
   compare_followers  Compare your followers with previous iterations
```
