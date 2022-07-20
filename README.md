# Git with Filecoin

An attempt to maintain Git repos in Filecoin

## Requirements:

 - Git
 - Web3.Storage api token

## Building

```shell
go install
```

## Running

```shell
git clone your-repo.git
cd your-repo

# make some changes

git_with_filecoin -commit -token TOKEN

# switch to your another computer
git_with_filecoin -apply -token TOKEN -cid CID
```
