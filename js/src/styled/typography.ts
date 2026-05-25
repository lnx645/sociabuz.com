import { css } from "@emotion/css";

const textCenter = css({ textAlign: "center" });
const textLeft = css({ textAlign: "left" });
const textRight = css({ textAlign: "right" });

const fontBold = css({ fontWeight: "bold" });
const fontSemiBold = css({ fontWeight: 600 });
const fontNormal = css({ fontWeight: "normal" });

const textSm = css({ fontSize: "12px" });
const textMd = css({ fontSize: "14px" });
const textBase = css({ fontSize: "16px" });
const textLg = css({ fontSize: "20px" });
const textXl = css({ fontSize: "24px" });
const uppercase = css({
    textTransform:"uppercase"
})
const underline = css({
  textDecoration: "underline",
});

const noUnderline = css({
  textDecoration: "none",
});

const lineThrough = css({
  textDecoration: "line-through",
});
const typography = {
  textCenter,
  underline,
  noUnderline,
  uppercase,
  lineThrough,
  textLeft,
  textMd,
  textRight,
  fontBold,
  fontSemiBold,
  fontNormal,
  textSm,
  textBase,
  textLg,
  textXl,
};

export default typography;
