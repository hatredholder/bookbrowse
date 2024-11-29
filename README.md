# bookbrowse

Find information about your favorite book without leaving the comfort of your terminal. Powered by [hardcover.app](https://hardcover.app/)

![example](docs/example.gif)

## Installation

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

- `-f`, `--fulldesc`: display full description
- `-h`, `--help`: help for bookbrowse
- `-m`, `--markdown`: format output into markdown
- `-v`, `--version`: display current version

### Advanced Usage

bookbrowse also supports formatting the book details into Markdown.

Here, for example, I wanted to have the `Fear and Loathing in Las Vegas` book's information saved neatly as a Markdown file, so I'm doing something like this:

```
bookbrowse --fulldesc --markdown Fear and Loathing in Las Vegas > "Fear and Loathing in Las Vegas".md
```

## Contributing

If you encounter any problems, get any ideas on what should this project have or feel like giving me some critique, you're free to open an issue.
