import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import styles from './index.module.css';

function Hero() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <div className={styles.heroContent}>
          <h1 className="hero__title">{siteConfig.title}</h1>
          <p className="hero__subtitle">{siteConfig.tagline}</p>
          <div className={styles.buttons}>
            <Link className="button button--secondary button--lg" to="/docs/intro">
              Introduction
            </Link>
            <Link className="button button--outline button--lg" to="/docs/getting-started/installation">
              Get Started
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
}

function Features() {
  const features = [
    {
      title: 'Deterministic context',
      description: 'Agents read actio/index.yaml first. No more random file scanning—routing is explicit and versioned.',
      to: '/docs/concepts/index-yaml',
    },
    {
      title: 'Sidecar layout',
      description: 'actio/ holds architecture, interfaces, patterns, rules, and tasks. One source of truth for AI and humans.',
      to: '/docs/concepts/sidecar',
    },
    {
      title: 'CLI + validation',
      description: 'actio create, init, validate, doctor. Schema and referential checks keep the sidecar consistent.',
      to: '/docs/cli/create',
    },
    {
      title: 'Plugins & MCP',
      description: 'Extend validation with YAML plugins. Expose Actio context to Cursor and other tools via actio mcp.',
      to: '/docs/guides/plugins',
    },
  ];
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {features.map(({ title, description, to }) => (
            <div key={title} className="col col--6 margin-bottom--lg">
              <div className={styles.featureCard}>
                <h3><Link to={to}>{title}</Link></h3>
                <p>{description}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

function QuickStart() {
  return (
    <section className={styles.quickStart}>
      <div className="container">
        <h2 className="text--center">Quick start</h2>
        <div className={styles.codeBlock}>
          <pre>
            <code>
{`# Create a new Actio-enabled project
actio create my-app
cd my-app

# Validate the sidecar
actio validate

# Add Actio to an existing repo
actio init`}
            </code>
          </pre>
        </div>
        <p className="text--center">
          <Link to="/docs/getting-started/installation">Install the CLI</Link>
          {' · '}
          <Link to="/docs/getting-started/quick-start">Quick start guide</Link>
        </p>
      </div>
    </section>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout title={siteConfig.title} description={siteConfig.tagline}>
      <Hero />
      <main>
        <Features />
        <QuickStart />
      </main>
    </Layout>
  );
}
