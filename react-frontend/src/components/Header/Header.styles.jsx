export const headerBackground = {
  background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
  padding: "20px",
  display: "flex",
  justifyContent: "center",
  boxShadow: "0 8px 32px rgba(0, 0, 0, 0.3)",
  width: "100%",
  position: "fixed",
  top: 0,
  left: 0,
  zIndex: 1000,
  backdropFilter: "blur(10px)",
};

export const linkStyle = {
  color: "#ffffff",
  textDecoration: "none",
  margin: "0 20px",
  fontSize: "1.1em",
  fontWeight: "500",
  padding: "8px 16px",
  borderRadius: "8px",
  transition: "all 0.3s ease",
};

export const linkStyleHover = {
  ...linkStyle,
  backgroundColor: "rgba(255, 255, 255, 0.2)",
  transform: "translateY(-2px)",
  boxShadow: "0 4px 12px rgba(255, 255, 255, 0.2)",
};
