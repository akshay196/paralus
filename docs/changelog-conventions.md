# How to generate CHANGELOG.md

```sh
conventional-changelog -r 2 -o AUTO_CHANGELOG.md -p conventionalcommits
```

# Commit Message

Commit messages must be written as per conventional commits conventions.

## Type

Must be one of the following:

- **fix**: A bug fix.
- **feat**: A new feature.
- **chore**: Miscellaneous changes.
- **build**: Changes that affect the build system or external dependencies (example scopes: go, npm) 
- **ci**: Changes to CI configuration files and scripts.
- **docs**: Documentation changes.
- **perf**: A code that improves performance.
- **refactor**: A code change that neither fixes a bug nor adds a feature.
- **style**: Changes that do not affect the meaning of the code (white-space, formatting etc)
- **test**: Adding missing tests or correcting existing tests

## Scope

TBD

## Subject

The subject contains a succinct description of the change:

  - use the imperative, present tense: "change" not "changed" nor "changes"
  - don't capitalize the first letter
  - no dot (.) at the end
  
## Body

Just as in the subject, use the imperative, present tense: "change"
not "changed" nor "changes". The body should include the motivation
for the change and contrast this with previous behavior.

## Footer

Breaking Changes should start with the word BREAKING CHANGE: with a
space or two newlines. The rest of the commit message is then used for
this.
