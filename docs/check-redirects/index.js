#!/user/bin/env node
const { RedirectChecker } = require('./check-redirects.js');
const yargs = require('yargs/yargs');
const { hideBin } = require('yargs/helpers');
const process = require('node:process');
const fs = require('node:fs');
const path = require('node:path');

const args = yargs(hideBin(process.argv))
  .option('in', {
    describe:
      'Comma-separated list of root directory paths in which to check for necessary redirects.',
  })
  .option('config', {
    describe: 'path to a docs configuration file with a "redirects" key',
  })
  .demandOption(['in', ' config'])
  .help()
  .parse();

const checkDir = (dirPath, command, redirects) => {
  const checker = new RedirectChecker(fs, dirPath, redirects);
  // TODO: Call some function to check redirects
};

const conf = fs.readFileSync(args.config);
const redirects = JSON.parse(conf).redirects;

args.in.split(',').forEach(p => {
  checkDir(p);
});
