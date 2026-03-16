import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import styles from './index.module.css';

function Hero() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx(styles.hero)}>
      <div className="container">
        <div className={styles.heroInner}>
          <h1 className={styles.heroTitle}>{siteConfig.title}</h1>
          <p className={styles.heroSubtitle}>{siteConfig.tagline}</p>
          <div className={styles.heroCtas}>
            <Link className={clsx('button button--primary', styles.cta)} to="/docs/getting-started/installation">
              Get started
            </Link>
            <Link className={clsx('button button--secondary', styles.ctaSecondary)} to="/docs/intro">
              Read the docs
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
}

function CardGrid({title, items}) {
  return (
    <section className={styles.section}>
      <div className="container">
        <h2 className={styles.sectionTitle}>{title}</h2>
        <div className={styles.grid}>
          {items.map((item) => (
            <Link key={item.title} to={item.to} className={styles.card}>
              <div className={styles.cardHeader}>
                <div className={styles.cardTitle}>{item.title}</div>
                <div className={styles.cardArrow} aria-hidden="true">
                  →
                </div>
              </div>
              <div className={styles.cardBody}>{item.description}</div>
            </Link>
          ))}
        </div>
      </div>
    </section>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  const startHere = [
    {
      title: 'Get started',
      description: 'Install Actio and scaffold your first sidecar.',
      to: '/docs/getting-started/installation',
    },
    {
      title: 'Quick start',
      description: 'Create a project, validate, and initialize an existing repo.',
      to: '/docs/getting-started/quick-start',
    },
    {
      title: 'Concepts',
      description: 'Understand the sidecar, routing, domains, rules, and tasks.',
      to: '/docs/concepts/sidecar',
    },
    {
      title: 'CLI reference',
      description: 'Commands: create, init, validate, doctor, mcp, version.',
      to: '/docs/cli/create',
    },
  ];
  const whatYouCanDo = [
    {
      title: 'Deterministic routing',
      description: 'Use `actio/router.yaml` to route agents to the right context.',
      to: '/docs/concepts/index-yaml',
    },
    {
      title: 'Schema validation',
      description: 'Catch missing files, bad YAML, and broken references early.',
      to: '/docs/guides/schema-validation',
    },
    {
      title: 'Plugins',
      description: 'Add extra checks with simple YAML validation plugins.',
      to: '/docs/guides/plugins',
    },
    {
      title: 'MCP integration',
      description: 'Expose Actio context to AI tools over stdio via `actio mcp`.',
      to: '/docs/guides/mcp-integration',
    },
  ];
  return (
    <Layout title={siteConfig.title} description={siteConfig.tagline}>
      <Hero />
      <main>
        <CardGrid title="Start here" items={startHere} />
        <CardGrid title="What you can do with Actio" items={whatYouCanDo} />
      </main>
    </Layout>
  );
}
