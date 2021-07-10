
# CLI tool to unzip and open files in it

- Uses Golang 1.16 to use CLI tool to unzip and open zipped folder in VS Code IDE
- The tool helps unzip and open all files in that zip folder in a single command

## Usage

- Go to the folder where you have zipped file on your CLI
  - Example: In your CLI execute command: **cd <file_path>**
- Use the application name provided in the "Use" part in root.go file
  - Example: I use **cli-uzo-golang** followed by space and type **code**
  - This gives you all usage and helpful commands you can use
- The actual command to unzip and open folder in VS Code is
  - **cli-uzo-golang code <file_name.zip>**