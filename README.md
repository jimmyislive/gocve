
# GoCVE

GoCVE is a command line client that provides CVE info (queried from a local database). GoCVE provides simple commands to download and populate a DB(postgres or sqlite) which you can then use to list, search or get CVE info from.

GoCVE is a single binary that works on mac and linux.

# Usage

## Configure the GoCVE tool

Let's assume you wish to use a postgres DB. Configure it like below:

`gocve config set-db --dbType postgres --dbHost pg-docker --dbPort 5432 --dbUser postgres --tablename cve`

Remember to set the env var GOCVE_PASSWORD to your DB password. 

The configs you set will be written out to a config file at `~/.gocve/gocve.yaml` e.g.

dbhost: pg-docker
dbname: cvedb
dbport: 5432
dbtype: postgres
dbuser: postgres
tablename: cve

If you would like to use a different file for the configs, use the `--config` option to point to a different file. `--config` is a global flag and can be used in any of the following commands to point to a different source of config.

## Show the GoCVE config

`gocve config show`

## Download CVE data

`gocve db download`

If you have used the defauls, a file `allitems.csv.gz` will be downloaded to your local. You can unzip it by doing `gunzip allitems.csv.gz`. 

See `gocve help db download` for more details.

## Populate the DB

After downloading the data, you need to import it into a database. We will assume that your postgres instance has a DB called `cvedb` created. (If not connect to your postgres instance and run `create database cvedb;`)

Before you can use the data inside `allitems.csv` you have to make sure the character encoding is the same. Find your current encoding:

`file allitems.csv `

Your DB probably has UTF-8 encoding. To change to UTF-8 do:

`iconv -f ISO-8859-14 -t UTF-8 allitems.csv > allitems.utf8.csv`

To load `allitems.utf8.csv` into the DB, do:

`gocve db populate --fileName allitems.utf8.csv`

This will take a few minutes. NOTE: The above command programatically inserts the info into the DB. It does not use an COPY/LOAD utility.

You are now ready to use GoCVE !

## List all CVEs

`gocve list | more`

## Get details of a CVE

`gocve get CVE-2005-2266`

## Search for a CVE

`gocve search CVE-2005-22`

# Development

All dev work happens in a container. First build the container:

`make docker-build`

Next, exec into a shell to get your dev env:

`make docker-shell`

You can now build your go code:

`make go-build`


# TODO
* Unit tests
