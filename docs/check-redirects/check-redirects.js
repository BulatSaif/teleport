const yaml = require('yaml');
const path = require('path');

const docsPrefix = 'https://goteleport.com/docs';
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

    // Assemble a map of redirects for faster lookup
    redirects.forEach(r => {
      this.redirects[r.source] = true;
    });
  }

  check() {
    const results = checkDir(this.otherRepoRoot);
    if (results.length > 0) {
      const output = '- ' + results.join('\n - ');
      throw new Error(
        `Found docs URLs in ${this.otherRepoRoot} with no corresponding docs path or redirect:
${output}`
      );
    }
  }

  // checkDir recursively checks for docs URLs with missing docs paths or
  // redirects at dirPath. It returns an array of missing URLs.
  checkDir(dirPath) {
    const files = this.fs.readDirSync(dirPath, 'utf8');
    let result = [];
    files.forEach(f => {
      const fullPath = join(dirPath, f);
      const info = this.fs.statSync(fullPath);
      if (!info.isDirectory()) {
        result = concat(result, checkFile(fullPath));
      }
      result = concat(result, checkDir(fullPath));
    });
  }

  // checkFile determines whether docs URLs found in the file
  // match either an actual docs file path or a redirect source.
  // Returns an array of URLs with missing files or redirects.
  checkFile(filePath) {
    const docsPattern = new RegExp(
      'https://goteleport.com/docs/[w/._#-]+',
      'gm'
    );
    const text = this.fs.readFileSync(filePath);
    const docsURLs = docsPattern.exec(text);
    if (docsURLs == null) {
      return;
    }
    let result = [];
    docsURLs.forEach(url => {
      const docsPath = urlToDocsPath(url);
      const entry = this.fs.statSync(dp, {
        throwIfNoEntry: false,
      });
      if (entry != undefined) {
        return;
      }
      const pathPart = docsPath.slice(docsPrefix.length);
      if (this.redirects[pathPart] == undefined) {
        result.push(url);
      }
    });
  }

  urlToDocsPath(url) {
    const rest = url.slice(docsPrefix.length, url.length);
    return path.join(this.docsRoot, 'docs', 'pages', rest + '.mdx');
  }
}

module.exports.RedirectChecker = RedirectChecker;
