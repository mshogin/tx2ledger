# tx2ledger

The purpose of this application is to parse bank monthly statements
to the ledger format.

## Current Features

* https://www.deutsche-bank.de/

## Planned Features
* N26
* PayPal

## Quick start

### Clone the repository
git https://github.com/mshogin/tx2ledger.git

### Run the application
go run main.go -tx /path/to/the/file/with/transaction.csv -o /path/to/output/file.ledger -c /path/to/categories/and/patterns/file


## Installation

### Install golang
```bash
GOVERSION=1.16

curl --fail --silent --show-error --retry 10 --retry-connrefused https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz --output go${GOVERSION}.tar.gz

tar xvzf go${GOVERSION}.tar.gz 1>/dev/null

rm -rf /usr/local/go

mv go /usr/local/

cp -f /usr/local/go/bin/go /usr/bin/

/usr/bin/go version

```

## Donations
If this application helped you in any way, or you would like to support me working on it, please donate:
* Etherium <img width="16px" height="16px" src="https://ethereum.org/static/a183661dd70e0e5c70689a0ec95ef0ba/31987/eth-diamond-purple.png"/>: 0x14e2FC230271d1359A3f7bd0C33A05F06Ab0Fc92
* Toncoin <img width="16px" height="16px" src="https://raw.githubusercontent.com/mshogin/assets/master/crypto/logo/toncoin.svg"/>: EQBx9bFDNCnFcis-LDRuHVPd4RY-s9K2__9YMQBeTj7kQPxp

## Contribution

Please feel free to submit any pull requests or suggest any desired features to be added.

## Contributor List

### A very special thank you to all who have contributed to this program:

|User|
|--|--|
| [mshogin](https://github.com/mshogin) |
