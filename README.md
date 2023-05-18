# company

## Problem Statement
#### Technical requirements 
Build a microservice to handle companies. It should provide the following operations:
• Create 
• Patch 
• Delete
• Get One
Each company is defined by the following attributes:
• ID (uuid) required 
• Name (15 characters) required - unique 
• Description (3000 characters) optional 
• Amount of Employees (int) required 
• Registered (boolean) required 
• Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) required 

#### Expectations: 
As a deliverable, we expect a GitHub repository (or any other git based repo) with the source code. We would like the solution to contain clear instructions to set up and execute the project. 
We expect the solution to be production ready. 

#### Will be considered a plus: 
On each mutating operation, an event should be produced. 
Dockerize the application to be ready for building the production docker image 
Use docker for setting up the external services such as the database 
REST is suggested, but GRPC is also an option 
JWT for authentication 
Kafka for events 
DB is up to you 
Integration tests are highly appreciated
Configuration file 

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

#### Kafka

Set up a Kafka using the guidline [Kafka site](https://kafka.apache.org/quickstart).
or follow bellow bash instructions,

```bash 
cd Downloads
```

```bash 
wget https://dlcdn.apache.org/kafka/3.4.0/kafka_2.13-3.4.0.tgz
```

```bash
tar -xzf kafka_2.13-3.4.0.tgz
cd kafka_2.13-3.4.0/
```

### Configuration

#### Git

```bash
git init
```

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

#### `Company` Repository

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
psql -d company
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


```bash
go run ./server/kafka-consumer/cmd/main.go
```

Go to kafka folder(home/Downloads/kafka_2.13-3.4.0/) and run these commands in sperate terminals,

```bash
bin/zookeeper-server-start.sh config/zookeeper.properties
```

```bash
bin/kafka-server-start.sh config/server.properties
```

optional for testing the kafka queue that it's working or not?
```bash
bin/kafka-console-consumer.sh --topic <TOPIC_NAME> --from-beginning --bootstrap-server localhost:9092
```

```bash
bin/kafka-console-producer.sh --topic <TOPIC_NAME> --bootstrap-server localhost:9092
```

To load env
```bash
source ./server/api/cmd/.env
```

## Testing Using CURL Command From Terminal
#### Create Company

```bash
curl -X POST 'localhost:8000/api/v1/company/create' -H 'Content-Type: application/json' -d '{
    "name": "Test Inc.",
    "description": "This is a test company",
    "employees": 100,
    "registered": true,
    "type": "Sole"
}'
```

#### Get Company Info

```bash
curl -X GET 'localhost:8000/api/v1/company/info?uuid=<COMPANY_ID>'
```

#### Update Company Info

```bash
curl -X PATCH 'localhost:8000/api/v1/company/update' -H 'Content-Type: application/json' -d '{"id":"<COMPANY_ID>",
    "name": "Updated Inc.",
    "employees": 5000,
    "type": "NonProfit"
}'
```

#### Delete Company

```bash
curl -X DELETE 'localhost:8000/api/v1/company/delete?uuid=<COMPANY_ID>'
```