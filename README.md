# echo_path
Simple CLI program to print out $PATH variable as a list with ability to sort

# Synopsis

```shell
echo_path # To list out $PATH variables unsorted
echo_path -s # To list out $PATH variables sorted
```

# To install

## From source code

1. Make sure you have golang installed. [Link](https://go.dev/learn/)
2. Clone the repo and `cd` to it

```shell
git clone git@github.com:2Cheetah/echo_path.git
cd echo_path
```

3. Install

```shell
make build
```

4. Add to PATH

Get absolute path to your executable:

```shell
pwd
```

Add line to your shell config file, e.g. `~/.zshrc`, `~/.bashrc` or other:

`export PATH="$PATH:/[output from pwd]"`

e.g.:

`export PATH="$PATH:/workspaces/echo_path"`

Then, source the config file:

```shell
source ~/.zshrc
```
