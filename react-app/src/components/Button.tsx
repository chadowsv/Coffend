import React from "react";

type ButtonProps = {
  type: "submit" | "reset" | "button";
  text: string;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
};

const Button = ({ type, text, onClick }: ButtonProps) => {
  return (
    <button type={type} onClick={onClick}>
      {text}
    </button>
  );
};

export default Button;
