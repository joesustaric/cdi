# Continuous DisIntegration

This tool is (will be) designed to assist with identifying (and hopefully help rectify) actions that people and teams do to code bases which slow down Continuous Integration (Continuous DisIntegration).

Currently all this tool can do is count how many branches exist for a single GitHub repository.

# Coming Soon Features
In no particular order..  
- Given a GitHub Organization for all repositories, count the branches
- Break down counts to branches that can be deleted (already merged) and ones with unmerged work
- Functionality to delete branches that have been merged
- Functionality to delete branches that live older than x days (configurable).
- Analysis on if there is _any_feedback on PR's (comment counts and distributions) 

# Usage
<todo>

## Development
These instructions assume you are running a recent macOS. You may need to adapt slightly for Linux and Windows.

```bash
# Make sure you have go 1.+ installed.
brew install go

# Download the code somewhere, no GOPATH required
git clone git@github.com:joesustaric/cdi.git
cd cdi

# Install dependencies
#TODO

# Run tests
make test
```

### Dependency management

We're using Go 1.16+ and [Go Modules](https://github.com/golang/go/wiki/Modules)

<TODO later>

If you introduce a new package:

```bash
go get github.com/my/new/package
```

Then you can write that package to the `vendor/` with:

```bash
go mod vendor
```

## Contributing

1. Fork it
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes (`git commit -am 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new Pull Request

## Contributors
add ppl
