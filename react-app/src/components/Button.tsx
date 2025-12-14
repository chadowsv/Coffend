import React from "react";

type ButtonProps = {
  type: "submit" | "reset" | "button";
  text: string;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
  className?: string;
};

const Button = ({ type, text, onClick, className = '' }: ButtonProps) => {
  return (
    <button 
      type={type}
      onClick={onClick}
      className={`btn ${className}`.trim()}
    >
      {text}
    </button>
  );
};

export default Button;
