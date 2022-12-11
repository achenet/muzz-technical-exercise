# Muzz Technical Exercise

## Running tests
On a Linux machine with `bash` and `docker`, the `test.sh` file can be run to execute tests.

Docker must have permission to run in userspace for this file to work properly.
This can be done with the `sudo usermod -aG docker $USER` command. The api will look for the database at the IP specified in the `DB_IP` environment varialbe. This variable is
automatically set when running the test script.

## Assumptions made
For Part 1.i , I assumed that 'generating a random user' would mean generating a truly random user.
Just in case, I did add another handler that generates an actual user from a JSON payload, which can be accessed using the `POST /user/create?random=false` query parameter.

No assumption was made as to the amount of genders existing.
If one makes the assumption that there are only a finite number genders,
it is possible to use an ENUM instead of a STRING in the database.


## Decisions made
I used the `math/rand` package instead of the `crypto/rand` package because the former has a simpler API than the latter, and the use case was not security-sensitive.

The function `generateID` returns the number of nanoseconds elapsed since Jan 1 1970. This was chosen as a relatively simple way of generating unique ids that would still be integers.
There are probably better options however - if the constraint that ID is an integer is relaxed, it is possible to generate a UUID for each user and match.

The database IP was stored in an environment variable, because access is simpler than a file, if more configuration is required however, it would be better to use a simple `config.json` file.

