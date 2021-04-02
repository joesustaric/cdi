# Continuous DisIntegration
<todo some semi witty intro>

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
#TODO
```

### Dependency management

We're using Go 1.16+ and [Go Modules](https://github.com/golang/go/wiki/Modules)
# TODO later

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
