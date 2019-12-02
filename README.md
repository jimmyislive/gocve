
# GoCVE

GoCVE is a command line client that provides CVE info (queried from a local database). GoCVE provides simple commands to download and populate a DB(postgres or sqlite) which you can then use to list, search or get CVE info from.

GoCVE is a single binary that was tested on linux.

# Usage

## Configure the GoCVE tool

The configs you set will be written out to a config file at `~/.gocve/gocve.yaml`

If you would like to use a different file for the configs, use the `--config` option to point to a different file. `--config` is a global flag and can be used in any of the following commands to point to a different source of config.


### postgres
`gocve config set-db --dbType postgres --dbHost pg-docker --dbPort 5432 --dbUser postgres --tableName cve`

Remember to set the env var GOCVE_PASSWORD to your DB password 

### sqlite: 
`gocve config set-db --dbType sqlite --dbName cvedb.sqlite`

The configs you set will be written out to a config file at `~/.gocve/gocve.yaml`

## Show the GoCVE config

`gocve config show`

### postgres 
```
dbtype:  postgres
dbhost:  pg-docker
dbname:  cvedb
dbport:  5432
dbuser:  postgres
tablename:  cve
password:  xxxxx
```
### sqlite
```
dbtype:  sqlite
dbhost:  localhost
dbname:  cvedb.sqlite
dbport:  0
dbuser:  
tablename:  cve
```

## Download CVE data

`gocve db download`

If you have used the defauls, a file `allitems.csv.gz` will be downloaded to your local. You can unzip it by doing `gunzip allitems.csv.gz`.

See `gocve help db download` for more details.

## Populate the DB

After downloading the data, you need to import it into a database. 

Your DB probably has UTF-8 encoding. To change to UTF-8 do:

`iconv -f ISO-8859-14 -t UTF-8 allitems.csv > allitems.utf8.csv`

### postgres 
We will assume that your postgres instance has a DB called `cvedb` created. (If not connect to your postgres instance and run `create database cvedb;`)

To load `allitems.utf8.csv` into the DB, do:

`gocve db populate --fileName allitems.utf8.csv`

*NOTE:* This will take a few minutes. The above command programatically inserts the info into the DB. It does not use an COPY/LOAD utility.

### sqlite

`gocve db populate --fileName allitems.utf8.csv`

*NOTE:* This may take a while in sqlite! (We don't use the normal `.import` of sqlite as that results in a lot of parsing errors)

You are now ready to use GoCVE !

## List all CVEs

`gocve list | more`

```
Using config file: /home/gouser/.gocve/gocve.yaml
CVE-1999-0001 	 ip_input.c in BSD-derived TCP/IP implementations allows remote attackers to cause a denial of servic
CVE-1999-0002 	 Buffer overflow in NFS mountd gives root access to remote attackers, mostly in Linux systems.
CVE-1999-0003 	 Execute commands as root via buffer overflow in Tooltalk database server (rpc.ttdbserverd).
CVE-1999-0004 	 MIME buffer overflow in email clients, e.g. Solaris mailtool and Outlook.
CVE-1999-0005 	 Arbitrary command execution via IMAP buffer overflow in authenticate command.
...
...
```

## Get details of a CVE

`gocve get CVE-2005-2266`

```
CVE-2005-2266
=============
Status: Candidate

Description: Firefox before 1.0.5 and Mozilla before 1.7.9 allows a child frame to call top.focus and other methods in a parent frame, even when the parent is in a different domain, which violates the same origin policy and allows remote attackers to steal sensitive information such as cookies and passwords from web sites whose child frames do not verify that they are in the same domain as their parents.

...
...
```

## Search for a CVE

`gocve search CVE-2005-22`

```
Using config file: /home/gouser/.gocve/gocve.yaml
CVE-2005-2200
=============
Multiple unknown vulnerabilities in the MicroServer Web Server for Xerox WorkCentre Pro Color 2128, 2636, and 3545, version 0.001.04.044 through 0.001.04.504, allow attackers to bypass authentication.

CVE-2005-2201
=============
Unknown vulnerability in the MicroServer Web Server for Xerox WorkCentre Pro Color 2128, 2636, and 3545, version 0.001.04.044 through 0.001.04.504, allow attackers to cause a denial of service or access files via crafted HTTP requests.

...
...
```

# Development

All dev work happens in a container. First build the container:

`make docker-build`

Next, exec into a shell to get your dev env:

`make docker-shell`

You can now build your go code:

`make go-build`


# TODO
* Complete unit tests
