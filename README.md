# company

## Setup

The following setup procedures assume you are on a relatively fresh machine,
with even the repository not necessary installed; this will occur within the
setup steps at the appropriate time.

We'll first install several tools, utilities and much more, and then configure
them one-by-one to get into a properly setup state. Some steps will require
help from your Tech Lead to provide you external credentials to third party
systems.

### Installations

#### Golang

Note that these same commands below can be used to upgrade; just update the
version numbers and re-run them.

```bash
wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
rm go1.19.4.linux-amd64.tar.gz
```

#### Postgres

```bash
# prepare packages.
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'

# install packages.
sudo apt update
sudo apt -y install postgresql-15 postgresql-client-15
```

One manual step that is required is to go into
`/etc/postgresql/15/main/pg_hba.conf` and modify the first non-commented line
to look like this instead, so that you can access your local database server
easily:

```
local   all             all                               trust
```

Save the file, and then restart your server:

```bash
sudo systemctl restart postgresql@15-main.service
```

### Configuration

#### Golang

Add these to your `~/.bashrc` to be able to properly use Golang.

```bash
export GO111MODULE=on
export GOPATH=~/go
export GOROOT=/usr/local/go
```

#### Paths

The following `$PATH` additions will come in handy, so add them to your
`~/.profile` and restart the terminal.

```bash
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```

#### `company` Repository

Setup your SSH public/private key pair locally, after replacing `<email>` with
your email address (you can ignore entering any password, just hit `Enter` for
all prompts):

```
ssh-keygen -t ed25519 -C "<email>"
```

Follow the instructions at
https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account
for getting your SSH public key onto Github.

Clone the repository (and all submodules):

```bash
git clone git@github.com:fiazbilal/company.git --recurse-submodules
```

#### Go Modules

Install all Golang modules used throughout the codebase.

```bash
go mod download
```

#### Postgres

Create your base user & database which will contain all state & data. The
username, password & database will all be "`company`":

```bash
psql -U postgres -c "CREATE USER company PASSWORD 'company' SUPERUSER"
psql -U postgres -c "CREATE DATABASE company OWNER company"
```

Create table as,

```bash
psql -U company
```

```bash
CREATE TABLE companies (
  id UUID PRIMARY KEY,
  name VARCHAR(15) UNIQUE NOT NULL,
  description VARCHAR(3000),
  employees INT NOT NULL,
  registered BOOLEAN NOT NULL,
  type VARCHAR(50) NOT NULL
);
```


## Start

Here are a list of commands you can run, each in a separate terminal, to ensure
all parts of the project are properly running:

```bash
psql -U company
```

```bash
go run ./server/api/cmd/main.go
```

