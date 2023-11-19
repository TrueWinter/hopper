# Hopper

Hopper is a command-line tool to extract Minecraft resources for use in resource packs.

You should be familiar with the structure of resource packs and running command-line tools before using Hopper.

# Usage

Run `hopper.exe -help` for a list of flags and their purpose.

Running the following command will copy all assets with names containing either `blaze` or `magma` from 1.20.1 and asset index 5 to a folder called `test`.

```
build\hopper.exe -index 5 -version 1.20.1 -pattern blaze,magma -output test
```

This will only work if you have downloaded and opened 1.20.1 at least once.

# Disclaimer

Hopper is not an official Minecraft product and is not approved by or associated with Mojang or Microsoft.