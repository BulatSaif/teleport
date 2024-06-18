const yaml = require('yaml');
const path = require('path');

// RedirectChecker checks for Teleport docs site domains and paths within a
// given file tree and determines whether a given docs configuration requires
// redirects.
// @param fs - The filesystem to use. Either memfs or the NodeJS fs package.
// @param {string} root - file path in fs in which to check for required
// redirects.
// @param {object} redirects - array of objects with keys "source",
// "destination", and "permanent".
class RedirectChecker {
  constructor(fs, root, redirects) {
    this.fs = fs;
    this.root = root;
    this.redirects = redirects;
  }
}

module.exports.RedirectChecker = RedirectChecker;
