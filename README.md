# NAME
   `setzer` -- manipulate feeds and update data

# SYNOPSIS
   `setzer <command> [<args>]`
   `setzer <command> --help`

# COMMANDS

  | command    |      description                                           |
  |------------|------------------------------------------------------------|
  |`price`     |      show ETH/USD price from `<source>`                    |


# INSTALLATION

Install dependencies with Nix:

```
nix-channel --add https://nix.dapphub.com/pkgs/dapphub
nix-channel --update
nix-env -iA dapphub.{seth,jshon}
```
   |                |                                        |
   |----------------|----------------------------------------|
   |`make link`     |  install setzer(1) into `/usr/local`   |
   |`make uninstall`|  uninstall setzer(1) from `/usr/local` |
