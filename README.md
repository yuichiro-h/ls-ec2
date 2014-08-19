ls-ec2
======

## Usage

```bash
$ ls-ec2
```

with peco and ssh:

```bash
connect_to_ec2() { ssh $(ls-ec2 | peco | awk '{print "username@"$2}') }
alias s="connect_to_ec2"
```

```bash
$ s
```

## Install

```bash
$ brew tap yuichiro-h/ls-ec2
$ brew install ls-ec2
```

You need to set aws credentials environmental variable:

```bash
$ export AWS_ACCESS_KEY_ID="0123456789ABCDEFGHIJ"
$ export AWS_SECRET_ACCESS_KEY="0123456789ABCDEFGHIJ0123456789ABCDEFGHIJ"
```
