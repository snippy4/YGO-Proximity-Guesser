export const cardImageStyle = {
  width: "50px",
  height: "70px",
  borderRadius: "5px",
  border: "1px solid #444",
  marginRight: "15px",
  transition: "transform 0.3s ease",
  cursor: "pointer",
};

export const cardImageHoverStyle = {
  ...cardImageStyle,
  transform: "scale(3)",
  animation: "wiggle 0.5s ease-in-out",
};
