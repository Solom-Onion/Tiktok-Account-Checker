# Tiktok Account Checker
The tool used to check Tiktok session cookies on the Solomon service

### TODO:
- [x] Upload source that was implemented into Solomon 
- [ ] Add Tor session requests for the chrome-driver
- [ ] Improve speeds
- [ ] Make fully request based

## Getting Started
### Prerequisites
Before you can build and run the file, you'll need to have the following:
- Go
- Tiktok Session Cookies from Solomon
### Step 1: Clone the repository
Open a terminal and navigate to the directory where you want to clone the repository. Then, run the following command to clone the repository:
```bash
git clone https://github.com/Solom-Onion/Tiktok-Account-Checker
```
### Step 2: Edit the cookies file
In the root directory of the cloned repository, you'll find a folder with `./data/cookies.json`.
Each session is its own list obj in the json file.
The format will be as follows:
```json
[
  [
    {
      "domain": "www.tiktok.com",
      "name": "Session 1s cookies",
      "value": "SESSION VALUE",
      "path": "/",
      "expires": 1701388800
    }
  ],
    [
    {
      "domain": "www.tiktok.com",
      "name": "Session 2s cookies",
      "value": "SESSION VALUE",
      "path": "/",
      "expires": 1701388800
    }
  ]
]
```
### Step 3: Build the file
Navigate to the root directory of the cloned repository in the terminal. Then, run the following command to build the file:
```bash
go build
```
### Step 4: Run the file
```
./bin
```
_______________

## Screenshot
<img src="https://files.catbox.moe/t797im.png">
