import Guess from "../Game/Guess";
import { exampleContainerStyle } from "./ExampleGuesses.styles";

function ExampleGuesses() {
  const exampleGuesses = [
    {
      cardId: "63288573",
      cardName: "Sky Striker Ace - Kagari",
      closenessText: "3 card(s) away!",
      closenessValue: 3,
    },
    {
      cardId: "63166095",
      cardName: "Sky Striker Mobilize - Engage!",
      closenessText: "10 card(s) away!",
      closenessValue: 10,
    },
    {
      cardId: "32807846",
      cardName: "Reinforcement of the Army",
      closenessText: "Warm",
      closenessValue: 0.68,
    },
    {
      cardId: "89997728",
      cardName: "Toon Table of Contents",
      closenessText: "Warm",
      closenessValue: 0.67,
    },
    {
      cardId: "73628505",
      cardName: "Terraforming",
      closenessText: "Warm",
      closenessValue: 0.62,
    },
    {
      cardId: "51452091",
      cardName: "Royal Decree",
      closenessText: "Tepid",
      closenessValue: 0.58,
    },
    {
      cardId: "1475311",
      cardName: "Allure of Darkness",
      closenessText: "Tepid",
      closenessValue: 0.51,
    },
    {
      cardId: "23434538",
      cardName: 'Maxx "C"',
      closenessText: "Cold",
      closenessValue: 0.36,
    },
    {
      cardId: "14558127",
      cardName: "Ash Blossom & Joyous Spring",
      closenessText: "Cold",
      closenessValue: 0.35,
    },
  ];

  return (
    <div style={exampleContainerStyle}>
      {exampleGuesses.map((guess, index) => (
        <Guess
          key={index}
          cardId={guess.cardId}
          cardName={guess.cardName}
          closenessText={guess.closenessText}
          closenessValue={guess.closenessValue}
        />
      ))}
    </div>
  );
}

export default ExampleGuesses;
