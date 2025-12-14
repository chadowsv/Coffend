import React from 'react';

type HomeProps = React.PropsWithChildren<{
  id: string;
  title: string
}>;

export default function HomeSection({ id, title, children }: HomeProps) {
  return (
    <section id={id} className="sections_home">
      <h2 className="section_title">{title}</h2>
      <p className="section_text">
        {children}
      </p>
    </section>
  );
}