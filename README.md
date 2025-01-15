# close-apps

This is a simple Go script that closes all running apps on your Mac. It assumes you're using iTerm2 as your terminal but can be easily modified to work with any terminal.

## Installation

1. **Create a Scripts Folder**: If you don't already have a `scripts` folder in your home directory, create one by running:

   ```bash
   mkdir ~/scripts
   ```

2. **Add Scripts Folder to PATH**: Add the `scripts` folder to your PATH by adding this line to your `.zshrc` file:

   Open file:

   ```bash
   code ~/.zshrc
   ```

   Add the following line to the file:

   ```bash
   export PATH="$PATH:$HOME/scripts"
   ```

   Reload your shell configuration:

   ```bash
   source ~/.zshrc
   ```

3. **Build the Script**: First, you need to build the Go script using Taskfile. Ensure you have Task installed on your machine. You can install it by running:

   ```bash
   brew install go-task/tap/go-task
   ```

   After installing Task, navigate to the root directory of the project and run:

   ```bash
   task build
   ```

4. **Make the Script Executable**: After building the script, you need to make it executable by running:

   ```bash
   chmod +x ~/scripts/close
   ```

   You only need to do this the first time. After that, every time you do a new build it will immediately work.
