# Troubleshooting

Solutions to common CGTerm issues.

## Installation Issues

### Build Errors

**Problem:** `go build .` fails with compilation errors

**Solutions:**
1. Verify Go is installed correctly:
   ```bash
   go version
   ```
2. Ensure you're in the CGTerm directory:
   ```bash
   pwd
   cd CGTerm
   ```
3. Clean build cache:
   ```bash
   go clean -cache
   go build .
   ```

### Permission Denied

**Problem:** `sudo mv cgterm /usr/bin` fails with permission error

**Solutions:**
1. Check file exists:
   ```bash
   ls -la cgterm
   ```
2. Ensure you have sudo access
3. Try with explicit permissions:
   ```bash
   sudo install -m 755 cgterm /usr/bin/cgterm
   ```

### Command Not Found

**Problem:** `cgterm` command not found after installation

**Solutions:**
1. Verify installation location:
   ```bash
   which cgterm
   ```
2. Check if binary exists:
   ```bash
   ls -la /usr/bin/cgterm
   ```
3. Add to PATH if needed:
   ```bash
   export PATH=$PATH:/usr/bin
   ```

## Runtime Issues

### Help Command Not Working

**Problem:** `help` command produces large errors

**Status:** ✅ **Fixed** in recent versions

**If still occurring:**
1. Update to the latest version:
   ```bash
   git pull origin main
   go build .
   ```

### Clear Command Displays Rogue Characters

**Problem:** `clear` command shows `[` character

**Status:** ✅ **Fixed**

**If still occurring:**
1. Rebuild the project
2. Report the issue on GitHub

### Text Editors Not Working

**Problem:** nano, vi, vim, nvim don't work properly

**Status:** ✅ **Known issue - Fixed**

**Cause:** stdin handling in original version

**Solution:** Update to the latest version where this is fixed.

### Sudo Infinite Password Spam

**Problem:** `sudo` command prompts infinitely for password

**Status:** ✅ **Fixed**

**Solution:** Update to the latest version.

### Sheh Blocking Termination

**Problem:** Using `sheh` command prevents proper exit from CGTerm

**Status:** ✅ **Fixed**

**Solution:** Update to the latest version.

## Command Issues

### External Command Not Found

**Problem:** `fastfetch` or `sheh` commands not working

**Solutions:**
1. Verify the package is installed:
   ```bash
   # For fastfetch
   fastfetch --version
   
   # For sheh
   sheh --version
   ```
2. Install if missing:
   ```bash
   # fastfetch - see https://github.com/fastfetch-cli/fastfetch
   # sheh - see https://github.com/waxodium/sheh
   ```

### List Commands (ls*) Show Nothing

**Problem:** `lsa`, `lsd`, `lsf`, `lse` return empty or error

**Solutions:**
1. Verify you're in a readable directory:
   ```bash
   > cd ~
   > lsa
   ```
2. Check permissions:
   ```bash
   ls -la
   ```
3. Try changing directory:
   ```bash
   > cd /tmp
   > lsa
   ```

### Battery Status Not Available

**Problem:** `batt` command fails

**Solutions:**
1. Command may not be available on your system
2. Check if battery is present:
   ```bash
   ls /sys/class/power_supply/
   ```
3. Try as regular user (not sudo)

## Performance Issues

### CGTerm Running Slowly

**Solutions:**
1. Check system resources:
   ```bash
   free -h
   htop
   ```
2. Ensure you have 5MB+ free storage
3. Try restarting CGTerm
4. Check for background processes

## Getting More Help

### Debug Information

When reporting issues, collect:
1. CGTerm version (check GitHub releases)
2. Operating system and version
3. Go version: `go version`
4. Full error message
5. Steps to reproduce

### Report an Issue

1. Visit [GitHub Issues](https://github.com/MasterArd/CGTerm/issues)
2. Check if similar issue exists
3. Create new issue with debug information
4. Be descriptive and include examples

### Additional Resources

- [Project Website](https://msad.online/projects/)
- [GitHub Repository](https://github.com/MasterArd/CGTerm)
- See the [Custom Commands](Custom-Commands) guide if customizing

---

**Still stuck?** Feel free to open an issue on GitHub!
