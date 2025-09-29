// src/components/Header.js
import React, { useState } from "react";
import { headerBackground } from "./Header.styles";
import HeaderLink from "./HeaderLink";

const Header = () => {
  return (
    <nav style={headerBackground}>
      <HeaderLink href={"/"}>Home</HeaderLink>
      <HeaderLink href={"/how-to-play"}>How to play</HeaderLink>
    </nav>
  );
};

export default Header;
