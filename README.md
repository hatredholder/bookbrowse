# bookbrowse

Find information about your favorite book without leaving the comfort of your terminal. Powered by [hardcover.app](https://hardcover.app/)

![example](docs/example.gif)

## Installation

### Use a Package Manager

```bash
# Arch Linux
yay -S bookbrowse
```

### Build It Yourself

Clone the GitHub repository:

```
git clone https://github.com/hatredholder/bookbrowse
```

`cd` into the cloned repository:

```
cd bookbrowse
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

- `-h`, `--help`: help for bookbrowse
- `-t`, `--templates`: template for format output
- `-v`, `--version`: display current version

### Advanced Usage

bookbrowse supports [templates](https://pkg.go.dev/text/template). On install following directories and files are created inside user config directory:

```
bookbrowse
└── templates
    └── predefined 
        ├── default.tmpl
        └── markdown.tmpl
```

By default `default.tmpl` template is used (duh), to use other templates `-t | --template` flag has to be passed:

```sh
bookbrowse --template markdown "Fear and Loathing in Las Vegas" | glow
```

```sh
bookbrowse --template markdown "Fear and Loathing in Las Vegas" > "Fear and Loathing in Las Vegas".md
```

`.tmpl` extension must be ommited.

User defined templates are also supported:

```
bookbrowse
└── templates
    ├── predefined 
    │   └── ...
    ├── default.tmpl # will be picked instead of "predefined/default.tmpl"
    └── my-custom.tmpl
```

Functions:

- commify: turn slice array into comma-separated list
- truncate: truncate string, ex: `truncate {{ .long_string }} 100`
- format_rating: n/10 stars format

All functions could be found in `internal/utils/functions.go`

## Contributing

If you encounter any problems, get any ideas on what should this project have or feel like giving me some critique, you're free to open an issue.
