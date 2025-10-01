// src/components/Header.js
import React, { useState } from "react";
import { headerBackground, linkStyle, linkStyleHover } from "./Header.styles";

const HeaderLink = ({ href, children }) => {
  const [isHovered, setIsHovered] = useState(false);

  return (
    <a
      href={href}
      style={isHovered ? linkStyleHover : linkStyle}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {children}
    </a>
  );
};

export default HeaderLink;
