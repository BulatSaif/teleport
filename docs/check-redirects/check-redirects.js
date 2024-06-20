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

  check() {
    // TODO: Fill this in
  }

  checkDir(dirPath) {
    const files = this.fs.readDirSync(dirPath, 'utf8');
    files.forEach(f => {
      const fullPath = join(dirPath, f);
      const info = this.fs.statSync(fullPath);
      if (!info.isDirectory()) {
        checkFile(fullPath);
      }
      checkDir(fullPath);
    });
  }

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
    const docsPaths = docsURLs.map(url=>{
    	return urlToDocsPath(url);
    });

    docsPaths.forEach(dp=>{
    	const entry = this.fs.statSync(dp, {
	    throwIfNoEntry: false
	});
	if (entry != undefined) {
	    return
	}
	// TODO: Look up redirect and throw an error if it doesn't exist.
    });
  }

  urlToDocsPath(url){
      const prefix = "https://goteleport.com/docs";
      const rest = url.slice(prefix.length, url.length)
      return path.join(this.docsRoot, "docs", "pages", rest+".mdx")
  }
}

module.exports.RedirectChecker = RedirectChecker;
