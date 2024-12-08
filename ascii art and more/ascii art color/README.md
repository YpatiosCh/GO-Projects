# ASCII Art Color Application

This project is an ASCII art generator that supports color decoration through various formats, including named colors, RGB, HSL, and HEX codes. The application allows users to print ASCII art with customizable color options for terminal display and file output.

## Features

- Generate ASCII art from text input.
- Support for multiple color formats:
  - Named colors (e.g., `red`, `green`, `blue`)
  - HEX format (e.g., `#ff0000`)
  - RGB format (e.g., `rgb(255, 0, 0)`)
  - HSL format (e.g., `hsl(0, 100%, 50%)`)
- Option to specify primary and secondary colors for substrings.
- Output to terminal and optionally to a file.

## Getting Started

### Prerequisites

- Go 1.16 or later installed on your machine.
- Terminal that supports ANSI escape codes.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/ascii-art-color.git
   cd ascii-art-color

2. Execute
    ```bash
    go run . --color=<color1[,color2]> <substring> <fullInput> <font> [outputFile]

## Parameters
- `--color=<color1[,color2]>: Specify one or more colors in various formats (e.g., red, #ff0000, rgb(255, 0, 0), hsl(0, 100%, 50%)).`
- `<substring>: The substring to highlight in the ASCII art.`
- `<fullInput>: The complete string to convert into ASCII art.`
- `<font>: The ASCII art font to use (e.g., standard, shadow, thinkertoy).`
- `[outputFile] (optional): Specify a file name to save the output.`

## Available Colors

The application supports a wide range of named colors, represented below:

| Color Name  | Color Preview |
|-------------|---------------|
| Red         | <img src="https://www.colorhexa.com/ff0000.png" alt="red" width="30" height="30"/>         |
| Green       | <img src="https://www.colorhexa.com/00ff00.png" alt="green" width="30" height="30"/>       |
| Blue        | <img src="https://www.colorhexa.com/0000ff.png" alt="blue" width="30" height="30"/>        |
| Yellow      | <img src="https://www.colorhexa.com/ffff00.png" alt="yellow" width="30" height="30"/>      |
| Magenta     | <img src="https://www.colorhexa.com/ff00ff.png" alt="magenta" width="30" height="30"/>     |
| Cyan        | <img src="https://www.colorhexa.com/00ffff.png" alt="cyan" width="30" height="30"/>        |
| Black       | <img src="https://www.colorhexa.com/000000.png" alt="black" width="30" height="30"/>       |
| White       | <img src="https://www.colorhexa.com/ffffff.png" alt="white" width="30" height="30"/>       |
| Gray        | <img src="https://www.colorhexa.com/808080.png" alt="gray" width="30" height="30"/>        |
| Orange      | <img src="https://www.colorhexa.com/ffa500.png" alt="orange" width="30" height="30"/>      |
| Purple      | <img src="https://www.colorhexa.com/800080.png" alt="purple" width="30" height="30"/>      |
| Pink        | <img src="https://www.colorhexa.com/ffc0cb.png" alt="pink" width="30" height="30"/>        |
| Brown       | <img src="https://www.colorhexa.com/a52a2a.png" alt="brown" width="30" height="30"/>       |
| Navy        | <img src="https://www.colorhexa.com/000080.png" alt="navy" width="30" height="30"/>        |
| Teal        | <img src="https://www.colorhexa.com/008080.png" alt="teal" width="30" height="30"/>        |
| Lime        | <img src="https://www.colorhexa.com/00ff00.png" alt="lime" width="30" height="30"/>        |
| Olive       | <img src="https://www.colorhexa.com/808000.png" alt="olive" width="30" height="30"/>       |
| Gold        | <img src="https://www.colorhexa.com/ffd700.png" alt="gold" width="30" height="30"/>        |
| Coral       | <img src="https://www.colorhexa.com/ff7f50.png" alt="coral" width="30" height="30"/>       |
     |


