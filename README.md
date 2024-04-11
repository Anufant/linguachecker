# Name

Linguachecker is a linter for Golang that checks the language of comments in your code.

## Usage

After installation, you can run the linter using the following command:

```bash
linguachecker ./...
```

This command will check all files in the current directory and its subdirectories.

## Configuration

You can configure the linter by creating a `.comment-lang-linter.yaml` file in the root directory of your project. Here is an example configuration:

```yaml
language: en
```

In this example, the linter will check that all comments are written in English.

## TODO

- [ ] Add actions
- [ ] Add testdata filtration and config to filter out some files
- [ ] Sanitize multiline comments
- [ ] Check for code comments
- [ ] Add linting
- [ ] Detect code comments

## Contribution

We welcome any contribution to this project. If you found a bug or want to add a new feature, feel free to create an issue or pull request.

## License

This project is licensed under the MIT license. Details can be found in the `LICENSE` file.