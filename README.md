# memory-db
Simple in-memory database

##Requirements
You need Golang 1.5 or newer [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/code.html)  

##Installation
`go get github.com/alyoshka/memory-db`  

##Run
`memory-db`

##Usage
memory-db is simple in-memory database.
Receives commands via stdin, and writes responses to stdout.
Also supports interactive input.

### Data Commands
Database accepts the following commands:
SET name value – Set the variable name to the value value. Neither variable names nor values will contain spaces.
GET name – Print out the value of the variable name, or NULL if that variable is not set.
UNSET name – Unset the variable name, making it just like that variable was never set.
NUMEQUALTO value – Print out the number of variables that are currently set to value. If no variables equal that value, print 0.
END – Exit the program.
Commands must be fed to program one at a time, with each command on its own line.

### Transaction Commands
memory-db also supports database transactions by also implementing these commands:
BEGIN – Open a new transaction block. Transaction blocks can be nested; a BEGIN can be issued inside of an existing block.
ROLLBACK – Undo all of the commands issued in the most recent transaction block, and close the block. Prints nothing if successful, or prints NO TRANSACTION if no transaction is in progress.
COMMIT – Closes all open transaction blocks, permanently applying the changes made in them. Prints nothing if successful, or print NO TRANSACTION if no transaction is in progress.
