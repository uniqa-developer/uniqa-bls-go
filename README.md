uniqa-bls-go
-----------

Go module for [uniqa-bls-snark-rs](https://github.com/uniqa-developer/uniqa-bls-snark-rs/).

## Release process

* Create a new branch
* Delete the old libs
* Create a PR
* The CI will now build all the libs and store them as tar.gz. artifact. Use it to update all the libs and commit them.
* After the PR is merged, tag the version.

If needed, you can remove an old tag using `./scripts/remove_tag.sh VERSION`.
