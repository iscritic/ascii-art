## ascii-art

### Objectives

Ascii-art is a program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the `string` received using ASCII characters, as you can see in the example below:


- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.

### Instructions



- Some **banner** files with a specific graphical template representation using ASCII can be used. 

  - [shadow](shadow.txt)
  - [standard](standard.txt)
  - [thinkertoy](thinkertoy.txt)


### Usage

```console
$student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $

$student$ go run cmd/app/main.go "Hello" shadow | cat -e
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $

                                 
$student$ go run cmd/app/main.go "Hello" thinkertoy | cat -e
                 $
o  o     o o     $
|  |     | |     $
O--O o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```
### Testing

```
$chmod +x script.sh
$./script.sh
```

### Author

- [arthkim](https://01.alem.school/git/arthkim)