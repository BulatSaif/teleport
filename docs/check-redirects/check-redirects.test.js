import { Volume, createFsFromVolume } from 'memfs';
import { RedirectChecker } from './check-redirects.js';

describe('check files for links to missing Teleport docs', () => {
  const files = {
    '/blog/content1.mdx': `---
title: "Sample Page 1"
---

This is a link to a [documentation page](https://goteleport.com/docs/page1).

This is a link to the [index page](https://goteleport.com/docs/).

`,
    '/blog/content2.mdx': `---
title: "Sample Page 2"
---

This is a link to a [documentation page](https://goteleport.com/docs/subdirectory/page2).

Here is a link to a [missing page](https://goteleport.com/docs/page3).

`,
    '/docs/content/1.x/docs/pages/page1.mdx': `---
title: "Sample Page 1"
---
`,
    '/docs/content/1.x/docs/pages/subdirectory/page2.mdx': `---
title: "Sample Page 2"
---
`,
    '/docs/content/1.x/docs/pages/index.mdx': `---
title: "Index page"
---
`,
  };

  test(`throws an error if there is no redirect for a missing docs page`, () => {
    const vol = Volume.fromJSON(files);
    const fs = createFsFromVolume(vol);
    const checker = new RedirectChecker(fs, '/blog', '/docs/content/1.x', []);
    expect(() => {
      checker.check();
    }).toThrow('https://goteleport.com/docs/page3');
  });

  test(`allows missing docs pages if there is a redirect`, () => {
    const vol = Volume.fromJSON(files);
    const fs = createFsFromVolume(vol);
    const checker = new RedirectChecker(fs, '/blog', '/docs/content/1.x', [
      {
        source: '/docs/page3/',
        destination: '/docs/page1',
        permanent: true,
      },
    ]);
    expect(() => {
      checker.check();
    }).not.toThrow('https://goteleport.com/docs/page3');
  });
});
