# Markdown Docs Generator

This is the docs-helper to automatically render the documentation for the project.

See [Makefile](Makefile) for the available commands.

## Workflow

**Pre-requisites:** Install the github CLI tool and authenticate it with your GitHub account.

```bash
    $> make git-clone
    $> make all
    $> make create-docs-pullrequest
```

This step will clone the opencloud and opencloud docs repository. 
Generate all markdown files and create a pull request with the changes towards the docs repository.