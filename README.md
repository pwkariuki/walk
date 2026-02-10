# walk

This is a simple command-line tool called `walk`, which crawls into file system directories looking for specific files. When the tool finds the files it’s looking for, it can list, archive, or delete them.

It is a learning project and **not production ready!**

You can use this tool manually, or even better, you can schedule it to run automatically by using a background job scheduler such as `cron`.

**Note:** Be careful when trying this tool on your system. The files will be deleted without any prompt or user confirmation.
Never run this tool as a privileged user such as root or Administrator because it can cause irreversible damage to your system.