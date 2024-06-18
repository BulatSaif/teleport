import { Volume, createFsFromVolume } from 'memfs';
import { RedirectChecker } from './check-redirects.js';

describe('check files for links to missing Teleport docs', () => {
  const fileTree = {
    '/page1.mdx': `---
title: "Sample Page 1"
---

This is a link to a [documentation page](https://goteleport.com/docs/page1).

`,
    '/page2.mdx': `---
title: "Sample Page 2"
---

This is a link to a [documentation page](https://goteleport.com/docs/page2).

`,
  };

  test(`throws an error if there is no redirect for a missing docs page`, () => {
    // TODO: Feed a filesystem with current docs pages to the checker
    const vol = Volume.fromJSON(fileTree);
    const fs = createFsFromVolume(vol);
    expect(() => {
      const frag = new RedirectChecker(fs, '/docs');
      // TODO: Add a function to run the check.
      // TODO: Add an expected error substring
    }).toThrow('');
  });
});
