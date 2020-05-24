# Gostuff

A collection of simple golang command-line utility tools.

For extra ease of use of installed commands, add the following settings:

```bash
cd
mkdir gobin  # directory for go installed files for easy access
```

In `.bash_profile`:

```bash
export GOBIN="/Users/[path]/gobin"
alias gocmds="ls $GOBIN"
```

Now, can access all installed go utilities via `gocmds`.
