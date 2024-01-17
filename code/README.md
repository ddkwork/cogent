![alt tag](logo/gide_icon.png)

**Gide** is a flexible IDE (integrated development environment) framework in pure Go, using the [GoGi](https://github.com/goki/gi) GUI (for which it serves as a continuous testing platform :) and the [GoPi](https://github.com/goki/pi) interactive parser for syntax highlighting and more advanced IDE code processing.

See the [Wiki](https://github.com/goki/gide/v2/wiki) for more docs,   [Install](https://github.com/goki/gide/v2/wiki/Install) instructions (`go install github.com/goki/gide/v2/cmd/gide@latest` should work if GoGi system libraries are in place), and [Google Groups goki-gi](https://groups.google.com/forum/#!forum/goki-gi) emailing list.

[![Go Report Card](https://goreportcard.com/badge/github.com/goki/gide/v2)](https://goreportcard.com/report/github.com/goki/gide/v2)
[![GoDoc](https://godoc.org/github.com/goki/gide/v2?status.svg)](https://godoc.org/github.com/goki/gide/v2)
[![Travis](https://travis-ci.com/goki/gide.svg?branch=master)](https://travis-ci.com/goki/gide)

There are many existing, excellent choices for text editors and IDEs, but *Gide* is possibly the best pure *Go* option available.  The Go language represents a special combination of simplicity, elegance, and power, and is a joy to program in, and is currently the main language fully-supported by Gide.  Our ambition is to capture some of those Go attributes in an IDE.

Some of the main features of *Gide* include:

* Designed to function as both a general-purpose text editor *and* an IDE.  It comes configured with command support for LaTeX, Markdown, and Makefiles, in addition to Go,a and the command system is fully extensible to support any command-line tools.

* Provides a tree-based file browser on the left, with builtin support for version control (git, svn, etc) and standard file management functionality through drag-and-drop, etc.  You can look at git logs, diffs, etc through this interface.

* Command actions show output on a tabbed output display on the right, along with other special interfaces such as Find / Replace, Symbols, Debugger, etc.  Thus, the overall design is extensible and new interfaces can be easily added to supply new functionality.  You don't spend any time fiddling around with lots of different panels all over the place, and you always know where to look to find something.  Maybe the result is less fancy and "bespoke" for a given use-case (e.g., Debugging), but our "giding" principle is to use a simple framework that does most things well, much like the Go language itself.

* Strongly keyboard focused, inspired by Emacs -- existing Emacs users should be immediately productive.  However, other common keyboard bindings are also supported, and key bindings are entirely customizable.  If you're considering actually using it, we strongly recommend reading the [Wiki](https://github.com/goki/gide/v2/wiki) tips to get the most out of it, and understand the key design principles (e.g., why there are no tabs for open files!).

# Install

* Wiki instructions: [Install](https://github.com/goki/gide/v2/wiki/Install) -- for building directly from source.

* See Releases on this github page for pre-built OS-specific app packages that install the compiled binaries.

* See `install` directory for OS-specific Makefiles to install apps and build packages.

# Current Status

As of April 2020, it is feature complete as a Go IDE, including type-comprehension-based completion, and an integrated GUI debugger, running on top of [delve](https://github.com/go-delve/delve).  It is in daily use by the primary developers, and very stable at this point, with the initial 1.0 release now available.

In addition to Issues shown on github, some important features to be added longer-term include:

* More coding automation, refactoring, etc.  We don't want to go too crazy here, preferring the more general-purpose and simple approach, but probably some more could be done.

* Full support for Python, including extending the [GoPi](https://github.com/goki/pi) interactive parser to handle Python, so it will have been worth writing instead of just using the native Go parser.

# Screenshots

![Screenshot](screenshot.png?raw=true "Screenshot")

![Screenshot, darker](screenshot_dark.png?raw=true "Screenshot, dark mode")

# TODO

* symbolsview icons not updating immediately -- fix from filenodedid not fix here

* lineno too conservative in not rendering bottom
* don't render top text, lineno if out of range
* always display cursor when typing!
* drag-n-drop tableview

* color highlighting for diff output in commandshell!
* outbuf use textview markup in addition to link formatting and other formatting.  tried but failed

* more helpers for URI api
* filter function for chooser for URI case

* Editable Chooser doesn't work with shift-tab -- requires focus prev fix to work properly!

* Find not selecting first item (sometimes?)

* fileview global priority key shortcuts
* editor rendering overflow
* filetree --get rid of empty upper level or not?
* sliceview / tableview should activate select and focus on selectidx item in readonly chooser mode -- select is working, but focus is not -- cannot move selection via keyboard


# DONE:



