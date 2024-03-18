# delta-shadeshifter

A convenience wrapper around [`delta`](https://dandavison.github.io/delta/) which automatically detects OS light/dark mode (on Windows...).

Works both in Windows and under WSL, but no other OS-es are implemented.

To install, run
```
go install github.com/tomasaschan/delta-shadeshifter@main
```
and configure according to [instructions for `delta`](https://dandavison.github.io/delta/configuration.html), except use `delta-shadeshifter` everywhere the `delta` executable is referenced in that config.
