import React from 'react';

type CardProps = React.PropsWithChildren<{
  className?: string;
  style?: React.CSSProperties;
  name?: string;
  description?: string;
  price?: number;
}>;

export default function Card({ 
  children, 
  className = '', 
  style,
  name,
  description,
  price
}: CardProps) {
  return (
    <div className={['card', className].filter(Boolean).join(' ')} style={style}>
      {children || (
        <div className="card-content">
          {name && <h3 className="card-title">{name}</h3>}
          {description && <p className="card-description">{description}</p>}
          {price !== undefined && <p className="card-price">${price.toFixed(2)}</p>}
        </div>
      )}
    </div>
  );
}