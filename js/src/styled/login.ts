import { css } from "@emotion/css";

export const loginFormStyled = {
  wrapper: css({
    maxWidth: 320,
    margin: "auto",
    color: "#4b4b4b",
    h1: {
      fontSize: 24,
      padding: 0,
      margin: 0,
    },
    p: {
      fontSize: 12,
    },
  }),
  fields: css({
    display: "flex",
    flexDirection: "column",
    gap: 12,
  }),
  header: css({
    marginBlock: 24,
    marginBottom: 10,
    textAlign: "center",
    h1: {
      fontWeight: 600,
      fontSize: 24,
    },
  }),
  buttons: css({
    display: "flex",
    alignItems: "center",
    flexDirection: "column",
    boxSizing: "border-box",
    gap: 10,
    button: {
      width: "100%",
    },
  }),
  separator: css({
    marginBlock: 8,
    fontWeight: "bold",
    color: "#afafaf",
  }),
  actionButtons: css({
    display: "flex",
    width: "100%",
    gap: 10,
    alignItems: "center",
    justifyContent: "space-between",
    flex: 1,
  }),
};
