/**
 * Documentation sidebar. Order defines the order of docs.
 */
const sidebars = {
  docsSidebar: [
    'intro',
    {
      type: 'category',
      label: 'Getting Started',
      link: { type: 'doc', id: 'getting-started/installation' },
      items: [
        'getting-started/installation',
        'getting-started/quick-start',
        'getting-started/create-project',
      ],
    },
    {
      type: 'category',
      label: 'Core Concepts',
      link: { type: 'doc', id: 'concepts/sidecar' },
      items: [
        'concepts/sidecar',
        'concepts/index-yaml',
        'concepts/domains',
        'concepts/rules-and-tasks',
      ],
    },
    {
      type: 'category',
      label: 'CLI Reference',
      link: { type: 'doc', id: 'cli/create' },
      items: [
        'cli/create',
        'cli/init',
        'cli/validate',
        'cli/doctor',
        'cli/mcp',
        'cli/version',
      ],
    },
    {
      type: 'category',
      label: 'Guides',
      link: { type: 'doc', id: 'guides/plugins' },
      items: [
        'guides/plugins',
        'guides/mcp-integration',
        'guides/schema-validation',
        'guides/use-cases',
      ],
    },
  ],
};

module.exports = sidebars;
