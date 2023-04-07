# I Don't Know ~ A simple CLI that knows for you.

[![Go Report Card](https://goreportcard.com/badge/github.com/abhisekmazumdar/idk)](https://goreportcard.com/report/github.com/abhisekmazumdar/idk)

IDK (I Don't Know) is a CLI that simplifies local development with ddev and lando by automatically detecting which one to use for each project. It can also recognize composer and drush commands in a similar manner.

If you are someone who uses `ddev` & `lando` in your day to day local development workflow and you get confused or don't remember which project you used for which, this CLI will know for you.

Note: Identifying `composer` & `drush` commands are very limited right now.

## Install

### 1. [Download the latest binary](https://github.com/abhisekmazumdar/idk/releases/latest)

### 2. Build binary

- Install Homebrew: https://brew.sh/

- Install go: `brew install go`

- Clone this repo in `$HOME` i.e. `/Users/[user-name]`:

    ```bash
    gh repo clone abhisekmazumdar/idk

    cd idk
    # Run go build to create the binary file
    go build
    ```

### 3. Go
```bash
go install github.com/abhisekmazumdar/idk@latest

# The above command will create the binary file in ~/go/bin
```

Then create an alias:

```bash
alias idk="~/go/bin/idk"

```

If you have never set up an alias before, follow these steps: [Supercharge Your Terminal Workflow: Step-by-Step Guide to Creating Aliases in Bash and Zsh](https://www.notion.so/abhisekmazumdar/Supercharge-Your-Terminal-Workflow-Step-by-Step-Guide-to-Creating-Aliases-in-Bash-and-Zsh-23b0f7db876e4ceda4f14ae6f2d616b3?pvs=4)

## Usage

```bash
cd path/to/a/drupal/project

# Use the newly created alias in place of ddev or lando.
idk start

# Use it the same way you would have ran composer commands.
idk composer instal

# Use it the same way you would have used other commands.
idk ssh

# also, this will run ddev/lando composer install
idk install

# this will run ddev/lando drush site-install
idk site-install

# To know more about the available commands
idk --help

```

Each time you run `idk` like the above said examples, it will show you whether it is using `ddev` or `lando`.

## Known Issues / Missing Features

Its still a Work In Progress tool. I see many corner case yet to cover like flags sometimes doesn't works. If you find any other issues/bugs kindly open an issue in this repo.

## Contributing

Contributions are welcome! To contribute to this project, fork the repository, make your changes, and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
