import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/Actio/__docusaurus/debug',
    component: ComponentCreator('/Actio/__docusaurus/debug', 'efd'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/config',
    component: ComponentCreator('/Actio/__docusaurus/debug/config', '628'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/content',
    component: ComponentCreator('/Actio/__docusaurus/debug/content', 'dc7'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/globalData',
    component: ComponentCreator('/Actio/__docusaurus/debug/globalData', 'aae'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/metadata',
    component: ComponentCreator('/Actio/__docusaurus/debug/metadata', '8ed'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/registry',
    component: ComponentCreator('/Actio/__docusaurus/debug/registry', '337'),
    exact: true
  },
  {
    path: '/Actio/__docusaurus/debug/routes',
    component: ComponentCreator('/Actio/__docusaurus/debug/routes', '458'),
    exact: true
  },
  {
    path: '/Actio/docs',
    component: ComponentCreator('/Actio/docs', '8f5'),
    routes: [
      {
        path: '/Actio/docs',
        component: ComponentCreator('/Actio/docs', '04c'),
        routes: [
          {
            path: '/Actio/docs',
            component: ComponentCreator('/Actio/docs', '54b'),
            routes: [
              {
                path: '/Actio/docs/cli/create',
                component: ComponentCreator('/Actio/docs/cli/create', '2a9'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/cli/doctor',
                component: ComponentCreator('/Actio/docs/cli/doctor', '5c2'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/cli/init',
                component: ComponentCreator('/Actio/docs/cli/init', '5eb'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/cli/mcp',
                component: ComponentCreator('/Actio/docs/cli/mcp', 'a06'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/cli/validate',
                component: ComponentCreator('/Actio/docs/cli/validate', '08e'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/cli/version',
                component: ComponentCreator('/Actio/docs/cli/version', '845'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/concepts/domains',
                component: ComponentCreator('/Actio/docs/concepts/domains', 'b05'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/concepts/index-yaml',
                component: ComponentCreator('/Actio/docs/concepts/index-yaml', 'dbf'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/concepts/rules-and-tasks',
                component: ComponentCreator('/Actio/docs/concepts/rules-and-tasks', 'df6'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/concepts/sidecar',
                component: ComponentCreator('/Actio/docs/concepts/sidecar', 'e08'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/getting-started/create-project',
                component: ComponentCreator('/Actio/docs/getting-started/create-project', '8e8'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/getting-started/installation',
                component: ComponentCreator('/Actio/docs/getting-started/installation', '95b'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/getting-started/quick-start',
                component: ComponentCreator('/Actio/docs/getting-started/quick-start', '51a'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/guides/mcp-integration',
                component: ComponentCreator('/Actio/docs/guides/mcp-integration', '60f'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/guides/plugins',
                component: ComponentCreator('/Actio/docs/guides/plugins', '757'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/guides/schema-validation',
                component: ComponentCreator('/Actio/docs/guides/schema-validation', '7ae'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/guides/use-cases',
                component: ComponentCreator('/Actio/docs/guides/use-cases', '7e6'),
                exact: true,
                sidebar: "docsSidebar"
              },
              {
                path: '/Actio/docs/intro',
                component: ComponentCreator('/Actio/docs/intro', '6f9'),
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
    path: '/Actio/',
    component: ComponentCreator('/Actio/', 'b06'),
    exact: true
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
