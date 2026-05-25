import { css } from "@emotion/css";

const block = css({ display: "block" });
const inlineBlock = css({ display: "inline-block" });
const inline = css({ display: "inline" });
const hidden = css({ display: "none" });

const wFull = css({ width: "100%" });
const hFull = css({ height: "100%" });
const wScreen = css({ width: "100vw" });
const hScreen = css({ height: "100vh" });

const layout = {
  block,
  inlineBlock,
  inline,
  hidden,
  wFull,
  hFull,
  wScreen,
  hScreen
};

export default layout;