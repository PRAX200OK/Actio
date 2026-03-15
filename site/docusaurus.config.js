// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
const {themes} = require('prism-react-renderer');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Actio Framework',
  tagline: 'AI-sidecar framework for deterministic context. Reduce hallucinations, enforce architecture.',
  favicon: 'img/favicon.svg',

  url: 'https://actio.dev',
  baseUrl: '/',

  organizationName: 'actio',
  projectName: 'actio',

  onBrokenLinks: 'throw',
  markdown: {
    hooks: { onBrokenMarkdownLinks: 'warn' },
  },

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: 'docs',
          sidebarPath: './sidebars.js',
          editUrl: undefined,
          showLastUpdateTime: false,
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Actio',
        logo: {
          alt: 'ACT Framework',
          src: 'img/logo.svg',
        },
        items: [
          { to: '/docs/intro', label: 'Docs', position: 'left' },
          { to: '/docs/getting-started/installation', label: 'Getting Started', position: 'left' },
          { to: '/docs/cli/create', label: 'CLI Reference', position: 'left' },
          {
            href: 'https://github.com/act-framework/act',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              { label: 'Introduction', to: '/docs/intro' },
              { label: 'Getting Started', to: '/docs/getting-started/installation' },
              { label: 'Concepts', to: '/docs/concepts/sidecar' },
              { label: 'CLI Reference', to: '/docs/cli/create' },
            ],
          },
          {
            title: 'Community',
            items: [
              { label: 'GitHub', href: 'https://github.com/act-framework/act' },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} ACT Framework. Built with Docusaurus.`,
      },
      prism: {
        theme: themes.github,
        darkTheme: themes.dracula,
        additionalLanguages: ['bash', 'yaml', 'json'],
      },
      colorMode: {
        defaultMode: 'light',
        respectPrefersColorScheme: true,
      },
    }),
};

module.exports = config;
