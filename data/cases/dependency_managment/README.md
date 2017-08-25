# Dependency Management

It always hard to manage all your dependencies, and upgrade never is easy as `2 + 2`. In
order to keep it under control it worth to spend some time and setup you project properly

## Ruby

For `ruby` the standard way to go is start new project with `bundler`. Getting started
with bundler is easy! Open a terminal window and run this command:

```sh
$ gem install bundler
```

Specify your dependencies in a `Gemfile` in your project's root:

```ruby
source 'https://rubygems.org'

gem 'nokogiri'
gem 'rack', '~> 2.0.1'
gem 'rspec'

# ...
```

Install all of the required gems from your specified sources:

```sh
$ bundle install
$ git add Gemfile Gemfile.lock
```

### Links
* http://bundler.io/v1.12/#getting-started
* https://www.cloudcity.io/blog/2015/07/10/how-bundler-works-a-history-of-ruby-dependency-management/
* http://bundler.io/v1.5/gemfile.html

## Go

For `Golang` there is no single right solution, and in many cases people just vendor all source and keep it right in the repo. But for now we will take a look to most similar to `bundler` solution and it's `gide`

TO get Glide just open terminal and run the following command:

```sh
$ curl https://glide.sh/get | sh
```

The script puts it with your Go binaries ($GOPATH/bin or $GOBIN).
Scan a codebase and create a `glide.yaml` file containing the dependencies with one single command:

```sh
$ glide init
```

Optionally, edit the `glide.yaml` file to add versions and other information. Install the
dependencies and revisions listed in the lock file into the vendor directory. If no lock
file exists an update is run.

```sh
$ glide install
```

Add a new dependency to the `glide.yaml`, install the dependency, and re-resolve the
dependency tree. Optionally, put a version after an anchor.

```sh
$ glide get github.com/foo/bar
```

Or

```sh
$ glide get github.com/foo/bar#^1.2.3
```

### Links
* https://glide.readthedocs.io/en/latest/
