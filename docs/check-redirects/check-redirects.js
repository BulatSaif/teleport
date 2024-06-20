const yaml = require('yaml');
const path = require('path');

// RedirectChecker checks for Teleport docs site domains and paths within a
// given file tree and determines whether a given docs configuration requires
// redirects.
// @param fs - The filesystem to use. Either memfs or the NodeJS fs package.
// @param {string} otherRepoRoot - directory path in fs in which to check for
// required redirects.
// @param {string} docsRoot - directory path in fs in which to check for present
// or missing docs files based on URL paths found in the directory tree at 
// otherRepoRoot.
// @param {object} redirects - array of objects with keys "source",
// "destination", and "permanent".
class RedirectChecker {
  constructor(fs, otherRepoRoot, docsRoot, redirects) {
    this.fs = fs;
    this.otherRepoRoot = otherRepoRoot;
    this.docsRoot = docsRoot;
    this.redirects = redirects;
  }
}

module.exports.RedirectChecker = RedirectChecker;
