<h1 align="center">decr</h1>

# About
Desktop Entry Creator, or `decr` for short, is a CLI program for adding `.desktop` files for executables.

# Installation
Before building from source, make sure you have a Go compiler on your system. (preferably 1.21 or newer but I can't
really stop you from editing `mod.go` now, can I?)

```
git clone https://github.com/local-interloper/decr
cd decr
go build
```
There should now be a file named `decr` in your project's root directory. Move it wherever you deem right. Or don't. Hopefully it finds the right`$PATH`.

# Usage
You can learn how to use the program by passing the `--help` flag:
```bash
decr --help
```

Practical example:
```bash
decr /opt/blender/blender
```

This will create `/home/<USER>/.local/share/applications/<PROGRAM-NAME>.desktop` file for use with an `xdg` compliant desktop environment.

# TODO
- Add Icon support