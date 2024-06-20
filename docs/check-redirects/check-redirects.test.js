import { Volume, createFsFromVolume } from 'memfs';
import { RedirectChecker } from './check-redirects.js';

describe('check files for links to missing Teleport docs', () => {
  const files = {
    '/blog/content1.mdx': `---
title: "Sample Page 1"
---

This is a link to a [documentation page](https://goteleport.com/docs/page1).

`,
    '/blog/content2.mdx': `---
title: "Sample Page 2"
---

This is a link to a [documentation page](https://goteleport.com/docs/subdirectory/page2).

Here is a link to a [missing page](https://goteleport.com/docs/page3.mdx).

`,
    '/docs/content/1.x/docs/pages/page1.mdx': `---
title: "Sample Page 1"
---

This is a link to a [documentation page](https://goteleport.com/docs/page1).

`,
    '/docs/content/1.x/docs/pages/subdirectory/page2.mdx': `---
title: "Sample Page 2"
---

This is a link to a [documentation page](https://goteleport.com/docs/page2).

`,
  }

  test(`throws an error if there is no redirect for a missing docs page`, () => {
    const vol = Volume.fromJSON(files);
    const fs = createFsFromVolume(vol);
    const checker = new RedirectChecker(fs, '/blog', '/docs/content/1.x', []);
    expect(() => {
    	checker.check();
    }).toThrow(/content2\.mdx.*page2.mdx/);
  });

    // TODO: Same test case as above, but with a redirect
});
