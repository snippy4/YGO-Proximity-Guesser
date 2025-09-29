import ExampleGuesses from "./ExampleGuesses";
import { subTitleStyle } from "./HowToPlayBody.styles";
import { containerStyle, titleStyle, textStyle } from "./HowToPlayBody.styles";

function HowToPlayBody() {
  return (
    <div style={containerStyle}>
      <h1 style={titleStyle}>How to play</h1>
      <h2 style={subTitleStyle}>How to play</h2>
      <p style={textStyle}>
        Each day a new random card will be selected to be the daily card (the
        daily card will reset at 6:00 UTC). Your goal is to try and guess that
        card based on the information your previous geusses have given you. Each
        time you guess a score will be provided to you, this score is an
        adjusted weighting of how frequently the correct card is played with
        your guess. The adjustments made to the weighting will reduce the score
        of frequently played cards, or in short it stops Maxx "C" from being the
        closest card to every other card and should make the proximity more
        telling of the actual card (note - adjustments are still being made to
        how this value is calculated and at the moment archetypal cards seem to
        have a larger score so often the closest cards to some staples will be
        the archetypes they are played in, the best tell of this is looking at
        warm and better cards to see if they would fit into the archetypes that
        appear close). If you need a hint then pressing the hint button will
        provide a guess closer than your current closest guess (unless you have
        the closest card already).
      </p>
      <ExampleGuesses />
    </div>
  );
}

export default HowToPlayBody;
