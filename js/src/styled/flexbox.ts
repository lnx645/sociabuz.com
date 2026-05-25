import { css } from "@emotion/css";

const flex = css({ display: "flex" });
const inlineFlex = css({ display: "inline-flex" });

const flexCol = css({ flexDirection: "column" });
const flexRow = css({ flexDirection: "row" });

const alignItemsCenter = css({ alignItems: "center" });
const alignItemsStart = css({ alignItems: "flex-start" });
const alignItemsEnd = css({ alignItems: "flex-end" });

const justifyCenter = css({ justifyContent: "center" });
const justifyBetween = css({ justifyContent: "space-between" });
const justifyStart = css({ justifyContent: "flex-start" });
const justifyEnd = css({ justifyContent: "flex-end" });
const gapXS = css({ gap: 5 });
const gapSM = css({ gap: "8px" });
const gapMD = css({ gap: "16px" });
const gapLG = css({ gap: "24px" });

const flexbox = {
  flex,
  inlineFlex,
  flexCol,
  flexRow,
  gapXS,
  alignItemsCenter,
  alignItemsStart,
  alignItemsEnd,
  justifyCenter,
  justifyBetween,
  justifyStart,
  justifyEnd,
  gapSM,
  gapMD,
  gapLG,
};

export default flexbox;
