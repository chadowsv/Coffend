import React from 'react';

type HomeProps = React.PropsWithChildren<{
  id: string;
  title: string
}>;

export default function HomeSection({ id, title, children }: HomeProps) {
  return (
    <section id={id} className="home-section">
      <h2 className="home-section-title">{title}</h2>
      <p className="home-section-text">
        {children}
      </p>
    </section>
  );
}