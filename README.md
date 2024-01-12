# Hopper

Hopper is a command-line tool to extract Minecraft resources for use in resource packs.

You should know how to create resource packs and run command-line tools before using Hopper.

# Usage

Run `hopper.exe -help` for a list of flags and their purpose.

Running the following command will copy all assets with names containing either `blaze` or `magma` from 1.20.1/asset index 5 to an `assets` folder in the `nether` folder. Filtered language files (for file names containing `en`) will also be outputted.

```
hopper.exe -index 5 -version 1.20.1 -pattern blaze,magma -output nether -lang en
```

This will only work if you have downloaded and opened 1.20.1 at least once, and if the `nether` folder exists.

# Disclaimer

Hopper is not an official Minecraft product and is not approved by or associated with Mojang or Microsoft.