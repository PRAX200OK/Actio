import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/docs',
    component: ComponentCreator('/docs', 'a25'),
    routes: [
      {
        path: '/docs',
        component: ComponentCreator('/docs', 'bea'),
        routes: [
          {
            path: '/docs',
            component: ComponentCreator('/docs', '989'),
            routes: [
              {
                path: '/docs/cli/create',
                component: ComponentCreator('/docs/cli/create', '7a2'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/cli/doctor',
                component: ComponentCreator('/docs/cli/doctor', 'e0c'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/cli/init',
                component: ComponentCreator('/docs/cli/init', '7e9'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/cli/mcp',
                component: ComponentCreator('/docs/cli/mcp', '879'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/cli/validate',
                component: ComponentCreator('/docs/cli/validate', 'cc0'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/cli/version',
                component: ComponentCreator('/docs/cli/version', '6e5'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/concepts/domains',
                component: ComponentCreator('/docs/concepts/domains', '813'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/concepts/index-yaml',
                component: ComponentCreator('/docs/concepts/index-yaml', '365'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/concepts/rules-and-tasks',
                component: ComponentCreator('/docs/concepts/rules-and-tasks', 'b32'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/concepts/sidecar',
                component: ComponentCreator('/docs/concepts/sidecar', '4cb'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/getting-started/create-project',
                component: ComponentCreator('/docs/getting-started/create-project', '550'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/getting-started/installation',
                component: ComponentCreator('/docs/getting-started/installation', 'f1f'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/getting-started/quick-start',
                component: ComponentCreator('/docs/getting-started/quick-start', '835'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/guides/mcp-integration',
                component: ComponentCreator('/docs/guides/mcp-integration', 'a42'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/guides/plugins',
                component: ComponentCreator('/docs/guides/plugins', '86a'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/guides/schema-validation',
                component: ComponentCreator('/docs/guides/schema-validation', '99c'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/guides/use-cases',
                component: ComponentCreator('/docs/guides/use-cases', '1d0'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/docs/intro',
                component: ComponentCreator('/docs/intro', '058'),
                exact: true,
                sidebar: "docsSidebar"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '/',
    component: ComponentCreator('/', '070'),
    exact: true
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
