import { useState } from "react";
import { cardImageStyle, cardImageHoverStyle } from "./Card.styles";

function Card({ cardId, cardName }) {
  const [isHovered, setIsHovered] = useState(false);

  return (
    <img
      src={`https://card.yugioh-api.com/${cardId}/image`}
      alt={cardName}
      style={isHovered ? cardImageHoverStyle : cardImageStyle}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    />
  );
}

export default Card;
