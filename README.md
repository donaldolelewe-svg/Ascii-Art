# ASCII Art Generator

## Description

ASCII Art Generator is a Go program that converts input text into ASCII art using a banner font file.
The constant banner font file set is (standard.txt), however there are three files containing different fonts. The banner file can be changed to either of the files.

The program reads each character from the input, finds its matching ASCII representation in the banner file, and prints the formatted output in the terminal.

## Features

- Converts normal text into ASCII art
- Supports multiple characters
- Supports newline input (`\n`)
- Uses a customizable banner file
- Runs directly from the terminal