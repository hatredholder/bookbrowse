# bookbrowse

![example](docs/example.gif)

## Installation

Clone the GitHub repository:

```
git clone https://github.com/hatredholder/bookbrowse
```

cd into the cloned repository:

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
