import React from 'react';

type CardProps = React.PropsWithChildren<{
  className?: string;
  style?: React.CSSProperties;
}>;

export default function Card({ children, className = '', style }: CardProps) {
  return (
    <div className={['card', className].filter(Boolean).join(' ')} style={style}>
      {children}
    </div>
  );
}