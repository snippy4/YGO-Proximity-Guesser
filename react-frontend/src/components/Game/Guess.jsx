import React from "react";
import Card from "./Card";
import ClosnessBubble from "./ClosnessBubble";
import {
  guessContainerStyle,
  cardInfoStyle,
  cardNameStyle,
} from "./Guess.styles";

const Guess = React.memo(function Guess({
  cardId,
  cardName,
  closenessText,
  closenessValue,
}) {
  return (
    <div style={guessContainerStyle}>
      <Card cardId={cardId} cardName={cardName} />
      <div style={cardInfoStyle}>
        <div style={cardNameStyle}>
          {cardName} <br /> {closenessText}
        </div>
        <ClosnessBubble value={closenessValue} />
      </div>
    </div>
  );
});

export default Guess;
