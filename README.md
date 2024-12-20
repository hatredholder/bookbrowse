# mediabrowse

Find information about your favorite book without leaving the comfort of your terminal. Powered by [hardcover.app](https://hardcover.app/)

![example](docs/example.gif)

## Installation

### Use a Package Manager

```sh
# Arch Linux
yay -S mediabrowse
```

### Build It Yourself

Clone the GitHub repository:

```
git clone https://github.com/hatredholder/mediabrowse
```

`cd` into the cloned repository:

```
cd mediabrowse
```

Install it with `go`:

```
go install .
```

## Setup

Before you can get your hands onto using this wonderful project, you must first get a [Hardcover](https://hardcover.app/) API key, which requires creating an account.

After you sign into your account, navigate to the [Hardcover API Access](https://hardcover.app/account/api) page, from where you copy your key (including the "Bearer" keyword).

Set the `HARDCOVER_API_KEY` environment variable.

## Usage

- `-h`, `--help`: help for mediabrowse
- `-t`, `--templates`: choose template for output format
- `-v`, `--version`: display current version

### Advanced Usage

#### Templates

mediabrowse supports [templates](https://pkg.go.dev/text/template). When mediabrowse is ran following directories and files are created inside `~/.config` directory:

```
mediabrowse
├── default.tmpl
└── markdown.tmpl
```

`default.tmpl` template is used by default, to use other templates `--template` flag has to be passed:

```sh
mediabrowse book "Fear and Loathing in Las Vegas" --template markdown > "Fear and Loathing in Las Vegas".md
```

`.tmpl` extension must be ommited when choosing a template.

##### Template Functions

- `commify`: turn slice array into comma-separated list
- `truncate`: truncate string elliptically
- `formatRating`: n/10 stars format

All functions can be found in `internal/templates/functions.go`

## Contributing

If you encounter any problems, get any ideas on what should this project have or feel like giving me some critique, you're free to open an issue.
